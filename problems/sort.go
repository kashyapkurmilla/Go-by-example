// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

// Submit your Go program source code.

package main

import "fmt"

func Swap(sli1 []int, i int) {
	var temp int
	temp = sli1[i]
	sli1[i] = sli1[i+1]
	sli1[i+1] = temp
}

func BubbleSort(sli1 []int, n int) {
	if n == 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		if sli1[i] > sli1[i+1] {
			Swap(sli1, i)
		}
	}

	BubbleSort(sli1, n-1)
}
func main() {
	var i int
	var sli1 = make([]int, 0)
	fmt.Println("Enter Elements :")
	for i = 0; i < 10; i++ {
		var temp int
		fmt.Printf(">")
		fmt.Scanln(&temp)
		sli1 = append(sli1, temp)
	}
	// fmt.Println(sli1)
	BubbleSort(sli1, len(sli1))
	fmt.Println("Sorted Array")
	fmt.Println(sli1)
}
