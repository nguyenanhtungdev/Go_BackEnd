package entity

//Tag JSON: Là metadata được định nghĩa trong dấu (backticks). Nó giúp kiểm soát cách
// các trường trong struct được ánh xạ (mapping) khi chuyển đổi giữastruct` và JSON
type User struct {
	UserId   string `json:"userId"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
}

type Person struct{
	PersonID string `json:"personID"`
	FullName string `json:"fullName"`
	Age int `json:"age"`
}

type Student struct{
	Person Person `json:"person"`
	ClassName string `json:"className"`
	Point float64
}