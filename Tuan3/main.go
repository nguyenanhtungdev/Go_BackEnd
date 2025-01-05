package main

import (
	"Tuan3/entity"
	"fmt"
	"sort"
)

type Salary float64
type Price float64

func main() {
	var user entity.User
	{
		user.UserId = "NV0001"
		user.FullName = "Nguyen Anh Tung"
		user.Age = 20
	}

	student := entity.Student{
		ClassName: "DHKTPM18ATT",
		Point: 20.5,
		Person: entity.Person{
			PersonID: "ID93812",
			FullName: "Nguyen Anh Tung",
			Age: 20,
		},
	}

	fmt.Println(user)
	fmt.Println(student)	
	fmt.Println(student.Person.Age)

	// jsonData, _ := json.Marshal(user)
	// fmt.Println(string(jsonData))
	// nhapLuong()
	// sortSalary()
	// testMap()
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
func nhapLuong() {
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

func sortSalary() {
	listEmployee := []float64{6.500, 1.500, 2.500, 2.000}
	listEmployee1 := []float64{1.1, 1.2, 1.3, 1.4}
	sort.Float64s(listEmployee)                               // Tăng dần
	sort.Sort(sort.Reverse(sort.Float64Slice(listEmployee1))) // Giảm dần
	fmt.Println("Danh sách nhân viên 1")
	for _, value := range listEmployee {
		fmt.Println(value)
	}
	fmt.Println("Danh sách nhân viên 2")
	for _, value := range listEmployee1 {
		fmt.Println(value)
	}

	people := []struct {
		name string
		age  int
	}{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}

	// Sắp xếp theo tuổi
	sort.Slice(people, func(i, j int) bool { return people[i].age < people[j].age })
	fmt.Println(people)

	// Kiểm tra một slice có đc sắp xếp tăng dần hay không (xem phần tử đã đúng vị trí chưa)
	fmt.Println(sort.Float64sAreSorted(listEmployee))
	fmt.Println(sort.Float64sAreSorted(listEmployee1))
}

func testMap() bool {
	listColor := map[string]string{"Red": "#da1337", "Orange": "#e95a22"} //Key nó giống như index trong array
	listColor["Blue"] = "#e95a22"
	delete(listColor, "Blue")
	for key, value := range listColor {
		fmt.Println(key + " " + value)
	}
	fmt.Println(len(listColor))
	// Kiểm tra một key đã tồn tại trong map chưa
	// value, exists := listColor["Blue"]
	// if exists {
	// 	fmt.Println(value)
	// }
	return true
}
