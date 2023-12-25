package basic

import "fmt"

func if_else() {
	userAge := 26

	if userAge < 20 {
		fmt.Println("Below 20")
	} else if userAge >= 20 && userAge <= 60 {
		fmt.Println("Between 20 and 60")
	}

	// If provide init value, then check logic
	if userName := "Jeremy"; userName == "John" {
		fmt.Println("You are John")
	} else { // else phải viết đúng theo format này, nếu k sẽ lỗi
		fmt.Println("You are not John")
	}
}
