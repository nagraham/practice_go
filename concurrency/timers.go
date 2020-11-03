package concurrency

import (
	"fmt"
	"time"
)

// This will only wait once! the timer does not reset itself
func sayHello(wait <-chan time.Time) {
	for {
		_ = <-wait
		fmt.Println("Hello world!")
	}
}

func sayHelloAsync(d time.Duration) *time.Timer {
	timer := time.NewTimer(d)
	go func() {
		for {
			_ = <-timer.C
			fmt.Println("Hello world!")
			d = d * 2
			timer.Reset(d)
		}
	}()
	return timer
}

func Timers() {
	timer := sayHelloAsync(500 * time.Millisecond)
	time.Sleep(10 * time.Second)
	timer.Stop()
}
