package concurrency

import (
	"fmt"
	"time"
)

func asyncFetchMessage(messageChan chan<- string) {
	time.Sleep(2 * time.Second)
	messageChan <- "Hello world!"
}

func Timeouts() {
	messageChan := make(chan string, 5)
	go asyncFetchMessage(messageChan)

	select {
		case msg := <-messageChan:
			fmt.Println(msg)
		case <-time.After(1 * time.Second):
			fmt.Println("Request timed out!")
	}
}

