package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	// Kết nối tới PostgreSQL
	db := ConnectDB()
	defer db.Close(context.Background()) // Đảm bảo đóng kết nối lại khi hàm main kết thúc

	// Đăng ký các route và handler
	http.HandleFunc("/users", GetUsersHandler(db)) //Dùng để liên kết (map) một route (đường dẫn URL) với một hàm xử lý
	http.HandleFunc("/users/create", CreateUserHandler(db))
	http.HandleFunc("/products", GetProductHandler(db))

	// Chạy server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
