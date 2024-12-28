package main

import (
	"Tuan3/entity"
	"encoding/json"
	"fmt"
)

type Salary float64
type Price float64

func main() {
	fmt.Println("Hi")
	var user entity.User
	{
		user.UserId = "NV0001"
		user.FullName = "Nguyen Anh Tung"
		user.Age = 20
	}

	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData))

	var salaryEmployee Salary
	fmt.Println("Nhập lương: ")
	fmt.Scanf("%f\n", &salaryEmployee)
	if !salaryEmployee.isCheck() {
		fmt.Println("Error!")
	}
	var price Price
	fmt.Println("Nhập giá sản phẩm")
	fmt.Scanf("%f\n", &price)
	if !price.isCheck() {
		fmt.Println("The price is error!")
	}
}

//Go không hỗ trợ toán tử ba ngôi (ternary operator)
// func (salary Salary) isCheck() bool {
// 	return (salary > 0.0) ? true : false
// }

func (salary Salary) isCheck() bool {
	return salary > 0.0
}

func (price Price) isCheck() bool {
	return price > 0.0
}
