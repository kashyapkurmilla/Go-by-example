package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	hour := currentTime.Hour()

	greeting := getGreeting(hour)
	quote := getWelcomeQuote()

	fmt.Println(greeting)
	fmt.Println(quote)
}

func getGreeting(hour int) string {
	switch {
	case hour >= 5 && hour < 12:
		return "Good morning!"
	case hour >= 12 && hour < 18:
		return "Good afternoon!"
	default:
		return "Good evening!"
	}
}

func getWelcomeQuote() string {
	return "Welcome to the world of programming with Go!"
}
