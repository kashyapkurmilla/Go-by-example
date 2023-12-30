package main

import "fmt"

func main() {
	var a float32
	var b int
	fmt.Printf("Enter floating point number : ")
	fmt.Scanf("%f", &a)
	fmt.Println(a)
	b = int(a)
	fmt.Println(b)
}
