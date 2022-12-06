package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go snd(c)

	// receive
	rcv(c)

	fmt.Println("about to exit")
}

// send
func snd(c chan<- int) {
	c <- 42
}

// receive
func rcv(c <-chan int) {
	fmt.Println(<-c)
}
