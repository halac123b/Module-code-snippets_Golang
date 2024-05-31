// Mỗi file code cần đc define thuộc 1 package (thường là folder chứa file đó)
package m_2

// Package fmt cung cấp các hàm để đọc ghi dữ liệu
import "fmt"

func main() {
	// Println in string bình thường
	fmt.Println("Hello World")
	// Printf in format string
	fmt.Printf("%d is a number", 10)

	//--- Variable area ---
	// Khai báo biến trong Go thứ tự ngược
	var a int = 4
	var b int
	b = 5
	fmt.Printf("Sum a + b = %d", a+b)

	// Declare rút gọn, biến phải là biến mới
	x := 9
	fmt.Printf("x = %d", x)
	// --- End Variable area ---

	// --- String area ---
	var str string = "Hello World"
	var multiLineStr string = `Hello
		World`
	fmt.Printf("str = %s\nMultiline String = %s", str, multiLineStr)
	// --- End String area ---
}
