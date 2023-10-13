package main

import (
	"fmt"
	"time"
)

// slowSender sends a string every 2 seconds.
func slowSender(c chan<- string) {
	for {
		time.Sleep(2 * time.Second)
		c <- "I am the slowSender"
	}
}

// fastSender sends consecutive ints every 500 ms.
func fastSender(c chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		c <- i
	}
}

func fasterSend(c chan<- []int) {
	for {
		c <- []int{1, 2, 3}
		time.Sleep(200 * time.Millisecond)

	}
}

// main starts the two senders and then goes into an infinite loop of receiving their messages.
func main() {
	ints := make(chan int)
	go fastSender(ints)
	strings := make(chan string)
	go slowSender(strings)
	sliceChan := make(chan []int)
	go fasterSend(sliceChan)

	for { // = while(true)
		select {
		case s := <-strings:
			fmt.Println("Received a string", s)
		case i := <-ints:
			fmt.Println("Received an int", i)
		case sl := <-sliceChan:
			fmt.Println("Received an slice:", sl)

		}
	}
}
