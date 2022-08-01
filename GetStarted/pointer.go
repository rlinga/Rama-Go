package main
import "fmt"

func main() {
	var x int = 1000
	var y int

	var xp *int
	xp = &x
	y = *xp

	fmt.Println(y)
}
