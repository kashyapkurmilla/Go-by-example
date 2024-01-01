// Delayed exit
package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("New GOOOOOO")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Main Routine")
}

/*
The main routine starts.
It initiates a new Goroutine using go fmt.Println("New GOOOOOO").
The main routine then encounters time.Sleep(1000 * time.Millisecond) and pauses for 1 second.- NOTE :  weare assuming that this routine takes 1sec (Not recomm	ended)
Meanwhile, the Goroutine executes fmt.Println("New GOOOOOO") concurrently and independently from the main routine.
*/
