package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var s Speaker
	d:= Dog{}
	s = d
	s.Speak()

	c:= Cat{}
	s=c
	s.Speak()

	//Khởi tạo một interface rỗng, nghĩa là nó có thể chứa được bất kỳ giá trị nào
	var i interface{}
	i = 5
	fmt.Println(i)
	i = "Hello"
	fmt.Println(i)

	check, ok := i.(string)
	if ok{
		fmt.Println("True " + check)
	}else{
		fmt.Println("False " + check)
	}

	//Sử dụng switch để kiểm tra dữ liệu
	checkType(42)
	checkType(true)
	checkType("Go")
	checkType(nil)

	var temp Speaker = Dog{}
	temp.Speak()
	temp.Move()
}

func checkType(i interface{}){
	switch v:= i.(type){
	case int:
		fmt.Println("La so nguyen ",v)
	case string:
		fmt.Println("La chuoi ", v)
	case bool:
		fmt.Println("La boolean", v)
	default:
		fmt.Println("Kieu khong xac dinh",v)
	}
}

type Speaker interface{
	Speak()
	Move()
}

type Dog struct{}

func(d Dog) Speak(){
	fmt.Println("Dog")
}

func(d Dog) Move(){
	fmt.Println("Dog move!")
}

type Cat struct{}

func(c Cat) Speak(){
	fmt.Println("Cat")
}

func(c Cat) Move(){
	fmt.Println("Cat move!")
}
