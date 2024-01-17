package main

import (
	"fmt"
)

func func1(nums []int, target int) []int {
	n := len(nums)
	var flag int
	var sli1 = make([]int, 0)
	for i := 0; i < n; i++ {
		if flag == 1 {
			break
		}
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				flag = 1
				sli1 = append(sli1, i)
				sli1 = append(sli1, j)
			}
		}
	}
	return sli1
}

func main() {
	// Example usage of func1
	numbers := []int{2, 7, 11, 15}
	goal := 9

	result := func1(numbers, goal)
	fmt.Println("Result:", result)
}
