package concurrency

import "fmt"

func simpleDeadlock() {
	// ch := make(chan int) 		// <-- deadlock
	ch := make(chan int, 1)			// <-- fix with buffered channel
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	fmt.Println(<-ch)
}
