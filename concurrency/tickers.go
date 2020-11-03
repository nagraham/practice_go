package concurrency

import (
	"fmt"
	"time"
)

func Tickers() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
			case <-done:
				return
			}
		}
	}()

	time.Sleep(3100 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Done!")
}