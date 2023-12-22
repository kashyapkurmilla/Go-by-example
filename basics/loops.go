package main

import "fmt"

func loops(count int) int { //single return value
	var sum int
	for i := 0; i < count; i++ {
		sum = sum + i
	}
	i := 1
	for {
		fmt.Println("in loop")
		i = i + 1
		if i == 2 {
			break
		}
	}
	return sum
}

func loops2(a float32, b int) (float32, int) {
	var sum float32
	var mul int

	//typecasting
	sum = a * float32(b)
	mul = int(a) + b

	fmt.Println(sum)
	fmt.Println(mul)
	return sum, mul
}
