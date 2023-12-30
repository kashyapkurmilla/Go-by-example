// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
//  The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length)3.
//  During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
//  The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
//  The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli1 := make([]int, 0, 3)
	fmt.Println("Enter integers (type kill to exit): ")

	for {
		var temp string
		fmt.Printf(">")
		fmt.Scanln(&temp)

		if temp == "x" || temp == "X" {
			break
		}

		num, err := strconv.Atoi(temp)
		if err != nil {
			fmt.Println("Enter valid integer")
			continue
		}
		sli1 = append(sli1, num)
		sort.Ints(sli1)
	}

	fmt.Println(sli1)
}
