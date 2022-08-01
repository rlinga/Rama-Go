package main

import "fmt"
import "sync"
//import "time"

var currentEatersCount int = 0

type ChopStick struct { sync.Mutex }

type Philosopher struct {
	leftCS, rightCS *ChopStick
	eatenSofarCount int
	done	bool
} 

func (p *Philosopher) eat(wg *sync.WaitGroup, n int) {
	p.leftCS.Lock()
	p.rightCS.Lock()
	currentEatersCount += 1
	fmt.Println("Philisopher", n, "Started Eating")

	p.eatenSofarCount += 1
//	fmt.Println("Philosopher", n, "ate", p.eatenSofarCount, "times")
	p.rightCS.Unlock()
	p.leftCS.Unlock()
	currentEatersCount -= 1
	fmt.Println("Philisopher", n, "Finished Eating")
	wg.Done()
}

func main() {
	CSticks := make([]*ChopStick, 5)
	for i:=0; i<5; i++ {
		CSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i:=0; i<5; i++ {
		philosophers[i] = &Philosopher{CSticks[i], CSticks[(i+1)%5], 0, false}
	}

	// Start the dining in main

	var eatenPhilosopherCount int = 0
	var wg sync.WaitGroup
	for (eatenPhilosopherCount < 5) {
		if (currentEatersCount == 2) {
			fmt.Println("CurrentEatersCount = ", currentEatersCount)
			continue
		}		
		for i:=0; i<5; i++ {
			if (philosophers[i].eatenSofarCount == 3) {
				if ( ! philosophers[i].done) {
					fmt.Println("Marking Done with philosopher", i+1)
					eatenPhilosopherCount += 1
					philosophers[i].done = true
				}
				continue
			}
			wg.Add(1)
			go philosophers[i].eat(&wg, i+1)	
			wg.Wait()
		}			
	}

	fmt.Println("Final eaten philosophers count", eatenPhilosopherCount)
}

