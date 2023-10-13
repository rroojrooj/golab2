package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	for {
		// Send "ping" to the channel
		fmt.Println("Foo is sending: ping")
		channel <- "ping"

		// Receiving message from the channel
		msg := <-channel
		fmt.Println("Foo has received:", msg)
	}
}

func bar(channel chan string) {
	for {
		// Receive message from the channel
		msg := <-channel
		fmt.Println("Bar has received:", msg)

		// Send "pong" to the channel
		fmt.Println("Bar is sending: pong")
		channel <- "pong"
	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	// Create a channel for string communication
	channel := make(chan string)

	// Start both foo() and bar() as goroutines
	go foo(channel)
	go bar(channel)

}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
	time.Sleep(500 * time.Millisecond)
}
