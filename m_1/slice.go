package m_1

// Slice giống list, là array có thể thay đổi độ dài
func slice() {
	var slice []int
	// Add element to slice
	slice = append(slice, 1)

	// Quick way to create slice
	exampleSlice := []int{1, 2, 3}
	exampleSlice[0] = 5

	// Dùng make() để tạo slice với độ dài ban đầu
	// 0: initial length, 5: capacity
	slice2 := make([]int, 0, 5)
	slice2[0] = 1
}
