package main

import (
	"fmt"
	"strings"
	"time"
)

func printMessage(ch, quit chan string) {
	for {
		select {
			case msg := <-ch:
				fmt.Println(msg)
			case <- quit:
				fmt.Println("All done")
				return
		}
	}
}

func createMessages(messages chan string) {
	strSlice := strings.Split("The sky above the port was the color of " +
		"television tuned to a dead channel", " ")
	for _, str := range strSlice {
		messages <- str
	}
}

func main() {
	messages := make(chan string)
	quit := make(chan string)
	go printMessage(messages, quit)
	go createMessages(messages)

	messages <- "Hello!"
	messages <- "Can you hear me?"
	messages <- "I'm talking on a different thread!"
	messages <- "Hey!"
	messages <- "Stop talking when I'm trying to talk!"
	messages <- "You're being very rude!"

	time.Sleep(1 * time.Second)
	quit <- ""
}

