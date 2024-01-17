package main

import (
	"fmt"
	"log"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	start, end := 1, x
	var result int
	for start <= end {
		mid := start + (end-start)/2
		if mid <= x/mid {
			start = mid + 1
			result = mid
		} else {
			end = mid - 1
		}

	}
	return result
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	result := mySqrt(number)
	fmt.Println("Square root of", number, "is", result)
}
