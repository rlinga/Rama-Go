I. There is always a main thread. "go" creates a new thread

// Main thread without new threads
a = 1
foo()
a = 2 // Will only be exexuted after foo

// Main thread without new threads
a = 1
go foo()
a = 2 // Will be exexuted before or after foo

II. Main thread exit will force other threads to exit

III. Synchronization between go routines can be controlled by using global events

Routine 1:

x = 1
x+=1
GLOBAL EVENT

Routine 2:

if GLOBAL EVENT 
	fmt.Println(x)

IV. sync package contains functions to sync between goroutines. sync.WaitGroup forces a goroutine to wait for other goroutines.

waitGroup contains an internal counter. It increments counter for each goroutine that it needs to wait on.

var wg sync.WaitGroup
wg.Add(1)

go foo(&wg)
wg.Wait()

=> foo() calls wg.Done()

wg.Add(1) => Increments the internal counter of wg
wg.Done() => Decrements the internal counter of wg

wg.Wait() => blocks until counter = 0

It decrements the counter for each gorouting that it is done with waiting.

Waiting gorouting cannot continue until counter = 0

V. Channels

Transfer data between goroutines using channels. Channels are typed.

a. channel := make(chan int) // Creates a channel to pass int
b. Send data on a channel using <- operator
	channel <- 3
c. Receive data from a channel
	x := <- channel

=> Channels are unbufferred by default. Sending blocks until data is received and receiving blocks until data is sent.

Task sending onto channel => c <- 300
Task receiving from channel => x := <- c

=> Declaring bufferred channels

c := make(chat int, 3) // Defines a channel that can hold max of 3 items
	=> Sending only blocks once it already has 3 items
	=> Receiving only blocks if channel is empty

