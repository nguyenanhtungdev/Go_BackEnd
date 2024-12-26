package main

import (
	"fmt"
)

func TestArray() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	// for _, value := range a {
	// 	fmt.Println(value)
	// }

	for index, value := range a {
		fmt.Printf("Index: %d, value:%d\n", index, value)
	}

	// Khai báo nhanh
	// a := [3]int{1, 2, 3}

	//Khai báo mà không cần ước lượng số phần tử ban đầu có
	// a := [...]int{1, 2, 3}

	// Mảng là một kiểu giá trị: Mảng trong Go là các loại giá trị không phải tham chiếu như các ngôn ngữ khác.
	//  Do đó ta thể tạo ra 1 copy của bản gốc bằng cách gán một mảng mới bằng mảng gốc ban đầu. Mọi thay đổi của mảng mới sẽ không ảnh hưởng đến mảng gốc
	// b := a
	// b[0] = 5
	// for index, value := range b {
	// 	fmt.Printf("Index: %d, value:%d\n", index, value)
	// }

	// Kích thước của mảng là một phần của kiểu giá trị. Trong go, mảng được định nghĩa bởi: kích thước + loại dữ liệu của phần tử
	// var b [5]int
	// b = a => chương trình sẽ bị báo lỗi
	// [3]int và [5]int là hai kiểu dữ liệu hoàn toàn khác nhau, dù cả hai đều chứa các phần tử kiểu int.
	// Người ta thiết kế kiểu này nhằm: tránh các lỗi liên quan đến việc truy cập ngoài phạm vi mảng (out-of-bound).

	// Slice không có kích thước cố định và có thể thay đổi linh hoạt. Slice này dùng chung vùng nhớ với mảng ban đầu.
	// arr := [...]int{1, 2, 3, 4, 5}
	// slice := arr[1:4]
	// for i := range slice {
	// 	slice[i]++
	// }
	// fmt.Println(slice)
	// fmt.Println(arr)

	// Tạo một slice độc lập:  là một cấu trúc dữ liệu rất linh hoạt trong Go, được sử dụng để làm việc với một danh sách các phần tử có kích thước động
	// arr := [...]int{1, 2, 3, 4, 5}
	// slice := append([]int{}, arr[1:4]...) // Dấu ... trong trường hợp này sẽ "mở rộng" slice arr[1:4] thành các phần tử riêng lẻ.
	// for i := range slice {
	// 	slice[i]++
	// }
	// fmt.Println("Slice:", slice)
	// fmt.Println("Array:", arr)

	// arr1 := [...]int{1, 2, 3, 4, 5}
	// slice1 := append([]int{}, arr1[:4]...) // Dấu ... trong trường hợp này sẽ "mở rộng" slice arr1[1:4] thành các phần tử riêng lẻ. [:]: tất cả, [:n], từ đầu đến n, [n:] từ n đến cuối
	// for i := range slice1 {
	// 	slice1[i]++
	// }
	// fmt.Println("Slice1:", slice1)
	// fmt.Println("array1:", arr1)

	// Tạo một slice trực tiếp
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 4, 5) // Thêm phần tử vào
	fmt.Println(slice)
	//Sao chép một slice
	desk := make([]int, len(slice)) // được sử dụng để khởi tạo và cấp phát bộ nhớ cho slice.      make([]T, length, capacity)
	copy(desk, slice)               // Sao chép slice
	fmt.Println(desk)
}
