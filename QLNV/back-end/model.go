package main

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ProductId   string  `json:"productID"`
	ProductName string  `json:"productName"`
	Price       int     `json:"price"`
	Description *string `json:"description"` // Sử dụng *string để xử lý null
}
