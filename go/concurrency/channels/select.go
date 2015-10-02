package main

import (
	"fmt"
	"time"
)

func worker(msg chan string) {
	t := time.NewTimer(time.Duration(1) * time.Second)
	stop := time.NewTimer(time.Duration(10) * time.Second)
	for {
		select {
		case word := <-msg:
			fmt.Printf("%s\n", word)
			time.Sleep(time.Duration(150) * time.Millisecond)
		case <-t.C:
			fmt.Println("PAUSE")
			time.Sleep(time.Duration(500) * time.Millisecond)
			t = time.NewTimer(time.Duration(1) * time.Second)
		case <-stop.C:
			return
		}
	}
}

func sender(words []string, chMsg chan string, done chan bool) {
	i := 0
	for {
		chMsg <- words[i]
		i++
		if i == 3 {
			i = 0
		}
	}

}
func main() {
	words := []string{"2pac", "big", "nas", "lil jon"}
	chMsg := make(chan string)
	done := make(chan bool)

	go worker(chMsg)
	go sender(words, chMsg, done)

	time.Sleep(12 * time.Second)
}
