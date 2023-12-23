package main

import "fmt"

func main() {
	fmt.Println("Hello World !")
	p1()
	sum := loops(6)
	fmt.Println("Sum = ", sum)
	a, b := loops2(4.656, 89)
	fmt.Println(a)
	fmt.Println(b)
	switchs()
}
