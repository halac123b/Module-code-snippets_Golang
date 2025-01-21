package m_1

import "fmt"

func add(a int, b int) int {
	return a + b
}

// Go có thể return nhiều giá trị
func sumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func add2(a int, b int) (sum int) { // here we are defining the variable name of what we are returning
	sum = a + b
	return // so no need for return a value, just call return to end function
}

// Đặt tên cho return, modify rồi return cùng lúc
func swap(a, b int) (x int, y int) {
	x = a
	y = b
	return
}

func call_func() {
	fmt.Println(add(1, 2))

	// Có thể ignore return value bằng cách dùng _
	sum, _ := sumAndDiff(5, 3)
	fmt.Println(sum)
}
