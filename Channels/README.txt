Channels
==========

1. Read from a channel

for i := range c {

}

Above for loop runs until channel c is closed by sender.

2. Select supports reading from one of the channels / which ever has data

select {

	case a = <- c1 :
		fmt.Println(a)
	case b = <- c2 :
		fmt.Println(b)

}

// Select either send or receive channels

select {

	case a = <- c1:
		fmt.Println("Received a")
	case c2 <- b:
		fmt.Println("Sent b")
}

// Abort channel implementation using for loop

for {
	select {
		case a <- c:
			fmt.Println(a)
		case <- abort:
			return
	}
}

// Interleaving

i = i+1 is set of 3 operations in machine code => Read i, increment, write I

Interleaving causes corruption of shared data if there is no proper mutex

To fix that, dont let 2 goroutines write to shared variable at same time

=> Binary semaphores
	=> If flag is down, shared variable is available for use. Else, its not available
	=> Lock() sets the flag up. Unlock() will set the flag down
	=> Example	
		var i int = 0
		var mut sync.Mutex
		func inc() {
			mut.Lock()
			i = i + 1
			mut.Unlock()
		}

=> Sync.Once gives the option to execute the function only once even if called from multiple goroutines
=> All calls to Once.Do(f) blocks until the f() is executed once.
=> Sync.Once is useful for initialization activities

var on sync.Once
var wg sync.WaitGroup

func setup() {
	fmt.Println("Init operations")
}

func doStuff() {
	on.Do(setup)
	fmt.Println("Good to proceed now that initialization is done")
	wg.Done() 
}

func main() {
	wg.Add(1)
	wg.Add(1)
	go doStuff()
	go doStuff()	
	wg.Wait()
}

// Deadlock implementation (circular dependencies)

fund doStuff(c1 chan int, c2 chan int) {
	<- c1
	c2 <- 100
	wg.Done()
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	wg.Add(2)
	doStuff(c1, c2)
	doStuff(c2, c1)
	wg.Wait()
}

// Go can find deadlocks if all goroutines are deadlocked. But Go CANNOT find deadlocks only some of the goroutines are causing deadlock.

