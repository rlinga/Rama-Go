package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	a := make([]int, 3)
	
	var s string

	for ; ; {
		fmt.Println("Enter the number")
		fmt.Scan(&s)

		if s == "X" {
			break
		}
		
		x, _ := strconv.Atoi(s)
		a = append(a, x)

		sort.Ints(a)

		fmt.Println("New Sorted Contents:", a)
	}
}
