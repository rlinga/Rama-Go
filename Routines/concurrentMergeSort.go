package main

import (
	"fmt"
	"sync"
	"sort"
)

func main() {
	fmt.Printf("How many numbers would you like to sort =>")
	n := 0
	fmt.Scanf("%d", &n)
	
	slice := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scanf("%d", &slice[i])
	}


	qs := n/4 // Quarter size
	slice1 := slice[0:qs]
	slice2 := slice[qs:2*qs]
	slice3 := slice[2*qs:3*qs]
	slice4 := slice[3*qs:n]

	var wg sync.WaitGroup
	wg.Add(4)
	go subSort(&wg, slice1)
	go subSort(&wg, slice2)
	go subSort(&wg, slice3)
	go subSort(&wg, slice4)

	wg.Wait()
	wg.Wait()
	wg.Wait()
	wg.Wait()

	resultSet := merge(slice1, slice2)
	resultSet = merge(resultSet, slice3)
	resultSet = merge(resultSet, slice4)

	fmt.Println("Final Sorted Array", resultSet)
}

func subSort(wg * sync.WaitGroup, nums []int) {
	fmt.Println("Sorting slice ", nums)
	sort.Ints(nums)	
	wg.Done()
}

func merge(firstSet []int, secondSet []int) []int {
	resultSet := make([]int, len(firstSet) + len(secondSet))
	i, j, k := 0, 0, 0	
	for i < len(firstSet) && j < len(secondSet) {
			if firstSet[i] < secondSet[j] {
				resultSet[k] = firstSet[i]
				k, i = k+1, i+1
			} else {
				resultSet[k] = secondSet[j]
				k, j = k+1, j+1
			}
	}

	for (i < len(firstSet)) {
		resultSet[k] = firstSet[i]
		k, i = k+1, i+1
	}

	for (j < len(secondSet)) {
		resultSet[k] = secondSet[j]
		k, j = k+1, j+1
	}

	return resultSet
}
