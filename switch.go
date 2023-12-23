package main

import (
	"fmt"
	"time"
)

func switchs() {

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weeeekeenndd!!1")
	case time.Friday:
		fmt.Println("Thank god its friday !!")
	default:
		fmt.Println("Normal days :(")
	}
}
