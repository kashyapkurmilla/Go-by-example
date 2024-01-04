// Write a program to sort an array of integers. The program should partition the array into 4 parts,
// each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
// When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"fmt"
	"sort"
	"sync"
)

func SortPartitions(arr []int, wg *sync.WaitGroup, index int) {
	defer wg.Done()

	start := index * (len(arr) / 4)
	end := (index + 1) * (len(arr) / 4)

	subarray := arr[start:end]
	sort.Ints(subarray)
	copy(arr[start:end], subarray)

	fmt.Printf("Sub-Array %d : %v\n", index+1, subarray)

}

func mergeArrays(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()

	sort.Ints(arr)
}
func main() {
	var n int
	fmt.Print("Enter number of elements in the array: ")
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf(">")
		fmt.Scanln(&arr[i])
	}

	var wg sync.WaitGroup
	wg.Add(4)

	//creates four go routines
	for i := 0; i < 4; i++ {
		go SortPartitions(arr, &wg, i)
	}
	wg.Wait()

	wg.Add(1)
	go mergeArrays(arr, &wg)
	wg.Wait()

	fmt.Printf("Sorted Array after merge : %v\n", arr)
}
