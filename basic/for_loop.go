// For loop là loại loop duy nhất cuả Go
package basic

import "fmt"

func for_loop() {
	// Syntax giống C
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Cũng có thể viết syntax giống while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Có thể iterate qua các container với index, value giống enumerate
	myList := []int{1, 2, 3}
	for index, value := range myList {
		fmt.Printf("%d is index, %d is value", index, value)
	}
	// Nếu không cần index, có thể dùng _ để bỏ qua
	for _, value := range myList {
		fmt.Println(value)
	}
}
