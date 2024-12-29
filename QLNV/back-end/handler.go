package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

// Lấy danh sách sản phẩm
func GetProductHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(context.Background(), "Select * from Product")
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var product Product
			if err := rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Description); err != nil {
				http.Error(w, "Error scanning product", http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

// Lấy danh sách người dùng
func GetUsersHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(context.Background(), "SELECT id, name, email FROM users")
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
				http.Error(w, "Error scanning user", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")
		// tạo một encoder JSON và gắn nó với http.ResponseWriter. Điều này có nghĩa là bất kỳ dữ liệu nào được mã
		//hóa bằng encoder này sẽ được ghi trực tiếp vào HTTP response và gửi về cho client.
		json.NewEncoder(w).Encode(users)
	}
}

// Tạo người dùng mới
func CreateUserHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
		if err != nil {
			http.Error(w, "Error inserting user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}
