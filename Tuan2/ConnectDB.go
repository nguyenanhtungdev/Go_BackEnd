package main

// cai đặt thư viện để kết nối go get github.com/jackc/pgx/v5
import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Customer struct {
	nameID   int
	fullName string
	age      int
}

func main() {
	// Chuỗi kết nối đến PostgreSQL
	connStr := "postgres://postgres:12345@localhost:5433/QLSV"

	// Kết nối đến PostgreSQL
	conn, err := pgx.Connect(context.Background(), connStr) //context.Background() nó được sử dụng để kiểm soát ngữ cảnh kết nối với PostgreSQL.
	if err != nil {                                         //nil là một giá trị đặc biệt được sử dụng để biểu thị rằng một biến chưa được khởi tạo hoặc không trỏ tới bất kỳ giá trị nào
		log.Fatalf("Không thể kết nối đến PostgreSQL: %v\n", err)
	}
	defer conn.Close(context.Background()) //Giải phóng tài nguyên được sử dụng bởi kết nối.

	fmt.Println("Kết nối thành công đến PostgreSQL!")

	// // Thêm dữ liệu vào table
	// _, err = conn.Exec(context.Background(), "INSERT INTO cars (brand, model, year) VALUES ($1, $2, $3)", "BMW", "p1800", 2018)
	// if err != nil {
	// 	log.Fatalf("Lỗi khi thêm dữ liệu qua mã Go: %v\n", err)
	// } else {
	// 	fmt.Println("Thêm dữ liệu thành công!")
	// }

	// // Kiểm tra truy vấn
	// rows, err := conn.Query(context.Background(), "Select *From cars")
	// if err != nil {
	// 	log.Fatalf("Lỗi khi truy vấn dữ liệu: %v\n", err) //In ra thông báo lỗi với định dạng cụ thể. Thêm dấu thời gian vào thông báo lỗi. Kết thúc chương trình ngay lập tức với mã thoát (exit code) là 1.
	// }
	// defer rows.Close() //  cần gọi rows.Close() để giải phóng tài nguyên. Sử dụng defer giúp bạn không quên việc đóng rows. Kết quả truy vấn rows sẽ được đóng ngay trước khi hàm main kết thúc.

	// fmt.Println("Danh sách người dùng:")
	// for rows.Next() {

	// 	var brand string
	// 	var model string
	// 	var year int
	// 	err = rows.Scan(&brand, &model, &year)
	// 	if err != nil {
	// 		log.Fatalf("Lỗi khi đọc dữ liệu: %v\n", err)
	// 	}
	// 	fmt.Printf("Brand: %s, model:%s, year:%d\n", brand, model, year)
	// }
	rows, err := conn.Query(context.Background(), "Select *From Customer")
	if err != nil {
		log.Fatalf("Lỗi khi truy vấn dữ liệu:%v\n", err)
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {

		var customer Customer
		//Kết quả từ truy vấn SQL thường được trả về dưới dạng con trỏ *sql.Rows.
		err = rows.Scan(&customer.nameID, &customer.fullName, &customer.age) // được sử dụng để quét (scan) dữ liệu từ một hàng kết quả (row) trong cơ sở dữ liệu vào các biến tương ứng.
		if err != nil {
			log.Fatalf("Lỗi khi đọc dữ liệu: %v\n", err)
		}
		customers = append(customers, customer)
	}

	for _, cucustomer := range customers {
		fmt.Printf("ID: %d, Họ và tên:%s, Tuổi:%d\n", cucustomer.nameID, cucustomer.fullName, cucustomer.age)
	}
	fmt.Printf("Address: %p", &customers)
}
