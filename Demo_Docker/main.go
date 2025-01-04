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
	mux.HandleFunc("/product",GetProductHandler(db))

	// Cấu hình CORS chỉ định các domain được phép: domain, method, header
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(mux)

	log.Println("Server is running on port 8070")
	http.ListenAndServe("0.0.0.0:8070", handler)

}
