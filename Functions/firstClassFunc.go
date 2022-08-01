package main

import "fmt"
import "math"

func main() {

	var a, v, s float64

	fmt.Println("Enter acceleration:")
	fmt.Scanf("%f", &a)
	fmt.Println("Enter velocity:")
	fmt.Scanf("%f", &v)
	fmt.Println("Enter displacement:")
	fmt.Scanf("%f", &s)

	fn := GenDisplaceFn(a, v, s)
	fmt.Println( "Result of fn(3) =>", fn(3))
	fmt.Println( "Result of fn(5) =>", fn(5))
}

func GenDisplaceFn(a, v, s float64) func (t float64) float64 {
	fn := func (t float64) float64 {
		displacement := 0.5 * a * math.Pow(t, 2) + v * t + s
		return displacement
	}
	return fn	
}
