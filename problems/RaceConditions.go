package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

func increment() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		counter++
	}
}
func main() {
	wg.Add(2)
	go increment()
	go increment()

	wg.Wait()
	fmt.Printf("Final Value : %d\n", counter)
}

/*
Race conditions occur due to the unpredictability of
the interleaving of these operations. Since there's
no guarantee about the order or timing of how these
goroutines are executed.

In this code there are two go routines created that councurrently incrementing counter value without any coordination
As a result, the final value of counter becomes unpredictable because the updates might overwrite each other,
leading to unexpected results might also result in deadloack.

*/
