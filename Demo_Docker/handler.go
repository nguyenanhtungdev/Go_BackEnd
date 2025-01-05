package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func GetProductHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(context.TODO(), "Select * from product;")
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			fmt.Println(err)
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {	
			var product Product
			if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Description); err != nil {
				http.Error(w, "Error scanning product", http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}