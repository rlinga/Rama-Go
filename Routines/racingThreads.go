package main

import "fmt"
import "time"

var balance int

// Two routines without synchronization can run in any order
// Following program demoes 2 go routines that run for 10 times each and the order is indeterministic
// There is no guarantee that depositing 200 for 10 times will make balance 2000 at the end until there is a synchronization control in between
func Deposit(amount int) {
	balance += amount
	fmt.Println("Balance in deposit", balance)
}

func Balance() int {
	return balance
}

func main() {
	for i := 0; i < 10; i++ {
		go Deposit(200)
		go Balance()
	}
	fmt.Println("Final Balance => ", balance)
 	time.Sleep(time.Second)
}
