package database

import (
	"context"
	"fmt"
	"log"

	// "os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	connStr := "postgres://postgres:12345@localhost:5432/QLSV"
	fmt.Println("Chuỗi kết nối:", connStr)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Không thể kết nối đến PostgreSQL: %v\n", err)
	}

	log.Println("Kết nối thành công đến PostgreSQL!")
	return conn
}
