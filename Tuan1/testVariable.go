package main

import "fmt"

func TestVariable() {
	var a int = 10
	_ = a
	fmt.Println("Hello World")

	var number int = 20
	fmt.Println(number)

	var b, c int
	c = 1
	b = 2
	fmt.Printf("Giá trị của b: %d, c:%d \n", b, c)

	var d, e int = 10, 20
	fmt.Printf("Giá trị của d:%d, e:%d, tổng:%d\n", d, e, d+e)

	var (
		name    string
		age     int
		salarry float64
	)

	name = "Nguyen Anh Tung"
	age = 21
	salarry = 5000.5000

	fmt.Printf("Tôi tên là %s, tôi %d tuổi, tôi có mức lương là %f\n", name, age, salarry)
	fmt.Printf("Tôi tên là %s, tôi %d tuổi, tôi có mức lương là %.3f\n", name, age, salarry)
	fmt.Printf("Tôi tên là %s, tôi %d tuổi, tôi có mức lương là %e\n", name, age, salarry)
	fmt.Printf("Tôi tên là %s, tôi %d tuổi, tôi có mức lương là %g\n", name, age, salarry)
	fmt.Printf("Tôi tên là %s, tôi %d tuổi, tôi có mức lương là %v\n", name, age, salarry)

	// Khai báo ngắn ngọn các biến
	address := "Quy nhơn"
	fmt.Println(address)
}
