// Implement the dining philosopher’s problem with the following constraints/modifications.

// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

// The host allows no more than 2 philosophers to eat concurrently.

// Each philosopher is numbered, 1 through 5.

// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself,
// where <number> is the number of the philosopher.

// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself,
// where <number> is the number of the philosopher.

// Submission: Upload your source code for the program.

package main

import (
	"fmt"
	"sync"
	"time"
)

const NumPhilosophers = 5
const NumMeals = 3

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	id                            int
	leftChopstick, rightChopstick *chopstick
	eatCount                      int
}

func (p philosopher) eat(host chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for p.eatCount < NumMeals {
		select {
		case <-host: //permission granted by the host
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()

			fmt.Printf("%d is Eating...\n", p.id)
			time.Sleep(time.Millisecond * 1000) //let him eat
			fmt.Printf("%d Finished eating\n", p.id)

			p.leftChopstick.Unlock()
			p.rightChopstick.Unlock()

			p.eatCount++
			host <- true

		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func main() {
	host := make(chan bool, 2) //two per session
	for i := 0; i < 2; i++ {
		host <- true
	}

	//creating chopsticks for each philosopher
	chopsticks := make([]*chopstick, NumPhilosophers)
	for i := 0; i < NumPhilosophers; i++ {
		chopsticks[i] = new(chopstick)
	}

	philosophers := make([]*philosopher, NumPhilosophers)

	var wg sync.WaitGroup

	for i := 0; i < NumPhilosophers; i++ {
		philosophers[i] = &philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%NumPhilosophers],
		}
	}

	wg.Add(NumPhilosophers)

	for _, p := range philosophers {
		go p.eat(host, &wg)
	}

	wg.Wait()
}
