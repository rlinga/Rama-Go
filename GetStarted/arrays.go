package main

import "fmt"

func main() {
	x := [5]int{2,3,4,5,6}

	for i:=0; i<len(x); i++ {
		fmt.Printf("int %d, value %d", i, x[i])
	}	
}
