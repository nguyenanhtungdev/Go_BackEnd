package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	db := ConnectDB()
	defer db.Close(context.Background())

	// Đăng ký handler
	http.HandleFunc("/users", GetUsersHandler(db))
	http.HandleFunc("/users/create", CreateUserHandler(db))

	// Chạy server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
