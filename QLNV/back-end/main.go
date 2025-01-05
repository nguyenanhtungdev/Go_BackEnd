package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/cors"
)

func main() {
	db := ConnectDB()
	defer db.Close(context.Background())

	// Đăng ký các route và handler
	mux := http.NewServeMux()
	mux.HandleFunc("/users", GetUsersHandler(db))
	mux.HandleFunc("/users/create", CreateUserHandler(db))
	mux.HandleFunc("/products", GetProductHandler(db))
	mux.HandleFunc("/products/update", UpdateProductHandler(db))
	mux.HandleFunc("/products/add",AddNewProductHandler(db))
	mux.HandleFunc("/products/sorted", GetSortedProductsHandler(db))


	// Cấu hình CORS chỉ định các domain được phép: domain, method, header
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS","PUT"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(mux)

	log.Println("Server is running on port 8020")
	log.Fatal(http.ListenAndServe(":8020", handler))
}
