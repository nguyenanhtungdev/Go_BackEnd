package main

import "fmt"

func SayHello() {
	fmt.Println("Xin chào từ helper.go")
	var isGioiTinh bool = true
	// vòng lặp thông thường
	for i := 0; i < 3; i++ {
		if isGioiTinh {
			fmt.Println("Là giới tính nam")
			break
		} else {
			fmt.Println("Là giới tính nu")
			continue
		}
	}
	// vòng lặp không có điều kiện, hoặc cũng có thể coi nó như vòng lặp do while
	for {
		if isGioiTinh {
			fmt.Println("hi")
			break
		}
	}
	// Vòng lặp tương tự như vòng lặp while trong các ngôn ngữ khác, vì trong Go không hỗ trợ while
	i := 1
	for i < 5 {
		i++
		fmt.Println("Yeeuuu em")
	}

	// Vòng lặp với mảng
	arr := []int{1, 2, 3, 4}
	for index, value := range arr {
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}
	// Nếu chỉ cần giá trị và không cần chỉ mục
	for _, value := range arr {
		fmt.Printf("Value: %d\n", value)
	}
	//Lặp qua map
	listSV := map[string]string{"name": "Nguyen Anh Tung", "name1": "Lam minh thai"}
	for key, value := range listSV {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	//Lặp qua chuỗi
	str := "nguyen anh tung"
	for _, value := range str {
		fmt.Printf("Value: %c\n", value)
	}
}
