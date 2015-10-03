package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func printer(done chan bool, block chan bool) {
	i := 0
	for {
		select {
		case <-done:
			fmt.Printf("Shutdown goroutine....\n")
			time.Sleep(time.Duration(1) * time.Second)
			block <- true
			return

		default:
			fmt.Printf("Printing... %d\n", i)
			time.Sleep(time.Duration(400) * time.Millisecond)
			i++
		}
	}
}

func signaler(c chan os.Signal, done chan bool) {
	<-c
	fmt.Println("I GOT signal")
	done <- true
}
func main() {

	ch := make(chan os.Signal, 1)
	done := make(chan bool)
	block := make(chan bool)
	signal.Notify(ch, os.Interrupt)

	go printer(done, block)
	go signaler(ch, done)

	<-block
}
