package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn { // Giống như instance
	// Định dạng lại chuỗi
	connStr := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Không thể kết nối đến PostgreSQL: %v\n", err)
	}

	log.Println("Kết nối thành công đến PostgreSQL!")
	return conn
}
