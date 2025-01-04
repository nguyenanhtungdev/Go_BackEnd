package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"

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
// API cập nhật sản phẩm
func UpdateProductHandler(db *pgx.Conn) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPut {
            http.Error(w, "Chỉ hỗ trợ PUT", http.StatusMethodNotAllowed)
            return
        }

        var product Product
        err := json.NewDecoder(r.Body).Decode(&product)
        if err != nil {
			http.Error(w, "Dữ liệu không hợp lệ", http.StatusBadRequest)
            return
        }

        // Cập nhật sản phẩm trong DB
        query := `UPDATE Product SET "productName" = $1, price = $2, description = $3 WHERE "productId" = $4`
		_, err = db.Exec(context.Background(), query, product.ProductName, product.Price, product.Description, product.ProductId)
        if err != nil {
            http.Error(w, "Lỗi khi cập nhật sản phẩm", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

//API sap xep 
func GetSortedProductsHandler(db *pgx.Conn)  http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		query := "Select *from Product"
		rows, err := db.Query(context.Background(),query)

		if err != nil{
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var product Product
			if err := rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Description); err != nil {
				http.Error(w, "Error scanning user", http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}

		sortField := r.URL.Query().Get("sortField")
		sortOrder := r.URL.Query().Get("sortOrder")

		sort.Slice(products, func(i, j int) bool {
			if sortField == "productName" {
				if sortOrder == "asc" {
					return products[i].ProductName < products[j].ProductName
				}
				return products[i].ProductName > products[j].ProductName
			}
			// Mặc định sắp xếp theo Price
			if sortOrder == "asc" {
				return products[i].Price < products[j].Price
			}
			return products[i].Price > products[j].Price
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func AddNewProductHandler(db *pgx.Conn) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
            http.Error(w, "Chỉ hỗ trợ POST", http.StatusMethodNotAllowed)
            return
        }
		var product Product
		err := json.NewDecoder(r.Body).Decode((&product))
		if err != nil{
			http.Error(w,"Dữ liệu không hợp lệ",http.StatusBadRequest)
			return
		}

		query := `Insert into Product ("productId","productName",price, description) values ($1,$2,$3,$4)`
		_, err = db.Exec(context.Background(),query,product.ProductId,product.ProductName, product.Price, product.Description)
		if err != nil {
            http.Error(w, "Lỗi khi thêm sản phẩm", http.StatusInternalServerError)
            return
        }
		w.WriteHeader(200)
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
