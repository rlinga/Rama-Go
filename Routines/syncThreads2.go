package main

import "fmt"
import "sync"

var balance int

func Deposit(wg *sync.WaitGroup, amount int) {
	balance += amount
	wg.Done()
}

func Balance(wg *sync.WaitGroup) int {
	fmt.Println("Balance=",balance)
	wg.Done()
	return balance
}

func main() {

	var wg sync.WaitGroup

	for i:=0; i<10; i++ {
		wg.Add(1)
		go Deposit(&wg,300)
		wg.Wait()
		wg.Add(1)
		go Balance(&wg)
		wg.Wait()
	}

	fmt.Println("Final Balance=",balance)
}
