package main

import "fmt"

func Swap(nums []int, j int) {
	if nums[j] > nums[j+1] {
		nums[j], nums[j+1] = nums[j+1], nums[j]
	}	
}

func BubbleSort(nums []int) {
	for i:=0; i<len(nums);i++ {
		for j:=0; j<len(nums)-i-1;j++ {
			Swap(nums, j)
		}
	} 	

}

func main() {
	nums := make([]int, 10)
	fmt.Println("Enter the numbers")
	for i:=0; i<10; i++ {
		fmt.Scanln(&nums[i])
	}

	BubbleSort(nums)
	
	fmt.Println("Sorted numbers are => ", nums)
}
