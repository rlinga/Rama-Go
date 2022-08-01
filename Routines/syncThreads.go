package main

import "fmt"
//import "time"
import "sync"

var balance int

func Deposit(wg *sync.WaitGroup, amount int) {
	balance += amount
	fmt.Println("Balance in deposit", balance)
	wg.Done()
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Deposit(&wg, 200)
		wg.Wait()
		go Balance()
	}

	fmt.Println("Final Balance => ", balance)
}
