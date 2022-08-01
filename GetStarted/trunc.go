package main
import "fmt"

func main() {

	var x float64
	
	fmt.Println("Enter a floating number of your choice")
	fmt.Scan(&x)

	y := int(x)

	fmt.Println("Integer value is ", y)
}
