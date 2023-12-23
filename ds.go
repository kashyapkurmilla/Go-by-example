package main

import (
	"fmt"
)

func main() {
	fmt.Println("Data Structures")
	arrays()
	slices()
	maps()
	ranges()
}

func arrays() {
	var arr [5]int
	fmt.Println(arr)

	arr[4] = 68
	fmt.Println(arr)
	var ele int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&ele)
		arr[i] = ele
	}
	fmt.Println(arr)
	fmt.Println(len(arr))
}

func slices() {
	var sli []string
	var sli2 []string
	s := make([]int, 0) //intial length
	sli = append(sli, "Slices ")
	sli = append(sli, "are the best")
	sli2 = append(sli2, ",are they ?")
	sli2 = append(sli2, sli...)
	fmt.Println(sli2)
	s = append(s, 2)
	fmt.Println(s)
}

func maps() {
	var n int
	m1 := make(map[string]int)
	fmt.Printf("Enter number of elements : ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var key string
		var value int
		fmt.Printf("Enter key : ")
		fmt.Scanln(&key)
		fmt.Printf("Enter corresponding value : ")
		fmt.Scanln(&value)

		m1[key] = value
	}
	fmt.Println(m1)
}
func ranges() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 9}
	sum := 0
	for _, num := range nums {
		sum = sum + num
	} // _ - blank identifier , we are not interested in the index for this iteration
	fmt.Println(sum)

	kvs := map[string]string{"hello": "world", "b": "a"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
