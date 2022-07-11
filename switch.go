package main

import "fmt"

var x int = 2

func main() {
	switch x {
		case 1: fmt.Println(x)
		case 2: fmt.Println(x)
		default: fmt.Println("Default")
	}
}
