package main

import (
	"fmt"
	"math"
)

func p1() {
	const pi = 3.14
	var a string = "string init"
	b := 4.22
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println("Kurmilla " + "kashyap")
	fmt.Println(90 + 7.1)
	fmt.Println(7.0 / 3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var rad float32
	fmt.Println("Enter radius :")
	fmt.Scanln(&rad)
	cir := 2 * pi * rad
	fmt.Println(cir)

	fmt.Println(math.Sin(float64(rad)))
}
