package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	connStr := "postgres://postgres:12345@localhost:5433/QLSV"
	// Kết nối đến PostgreSQL
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Không thể kết nối đến PostgreSQL: %v\n", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Kết nối thành công đến PostgreSQL!")
	return conn
}
