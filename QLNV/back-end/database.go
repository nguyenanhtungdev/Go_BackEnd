package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn { // Giống như instance
	connStr := "postgres://postgres:12345@host.docker.internal:5432/QLSV"


	// Kết nối đến PostgreSQL
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Không thể kết nối đến PostgreSQL: %v\n", err)
	}

	log.Println("Kết nối thành công đến PostgreSQL!")
	return conn
}
