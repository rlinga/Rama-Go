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

		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})

		fmt.Println("New Sorted Contents:", a)
	}
}
