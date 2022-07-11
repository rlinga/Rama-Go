package main
import "fmt"

func main() {
	var x int 
	fmt.Println("How many iterations do you want?")
	fmt.Scan(&x)
	for i:=1; i<x; i++ {
		fmt.Println(i)
	}
}
