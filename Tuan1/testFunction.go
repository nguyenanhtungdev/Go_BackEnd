package main

import "fmt"

func TestFunction() {
	SayHello1()
	SayHello2("Nguyen Tung")
	fmt.Println(SayHello3(3, 4))
	q, r := SayHello4(1, 2, 3)
	fmt.Println(q)
	fmt.Println(r)

	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))
}

// Hàm không có tham số
func SayHello1() {
	fmt.Println("Xin chào!")
}

// Hàm có tham số
func SayHello2(name string) {
	fmt.Println(name)
}

// Hàm có giá trị trả về
func SayHello3(a int, b int) int {
	return a + b
}

// Hàm có nhiều giá trị trả về
func SayHello4(a int, b int, c int) (int, int) {
	return a + b, b + c
}

// Hàm ẩn danh: là hàm không có tên, thường được sử dụng để tạo các hàm nhỏ gọn

// Trong Go, không thể khai báo hai hàm trùng tên trong cùng một package
// Trong Go, tất cả các file trong cùng một thư mục thì phải có cùng một package
