package main

import "fmt"
import "time"

// Two routines without synchronization can run in any order
// Following program demoes 2 go routines that run for 10 times each and the order is indeterministic
func t1(favMovies []string) {
	favMovies = append(favMovies, "Movie of T1 added")
	fmt.Println("In T1", favMovies)
}

func t2(favMovies []string) {
	favMovies = append(favMovies, "Movie of T2 added")
	fmt.Println("In T2", favMovies)
}

func change(movies []string) {
	movies[0] = "India";
}

func main() {
	movies :=[]string{"USA"}
	for i := 0; i < 10; i++ {
		go t1(movies)
		go t2(movies)
	}
	fmt.Println("Movies => ", movies)
	change(movies);
	fmt.Println("Movies => ", movies)
 	time.Sleep(time.Second)
}
