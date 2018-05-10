package main

import (
	"fmt"
	"time"
)

//channels is a way for two go routines to communicate with each other and synchronize their execution
//bidirectional channel
func pinger(c chan string) {

	c <- "ping"
}

//bidirectional channel
func ponger(c chan string) {

	c <- "pong"
}

func printer(c chan string) {

	msg := <-c
	fmt.Println(msg)
}

// 'chan<- val-type' means only allowed to send to c
func pingerOnlySend(c chan<- string) {

}

// '<-chan val-type' means only allowed to receive
func printOnlyReceive(c <-chan string) {

}

func main() {
	fmt.Println("Go channels")
	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "messagePing" }()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	//var c chan string = make(chan string)
	// c := make(chan string)
	// go pinger(c)
	// go printer(c)

	//when a pinger attempts to send a message on the channel, it will wait unit the printer is ready to receive the message (known as blocking)
	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from channel c1"
			time.Sleep(time.Second + 1)
		}
	}()

	go func() {
		for {
			c2 <- "from channel c2"
			time.Sleep(time.Second + 2)
		}
	}()
	//'select' keyword example which acts like a switch for channels
	func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	//buffered channels - asynchronous channel. sending or receiving a message will not wait unless the channel is already full
	//creates a buffered channel with a capacity of 1.

	//c := make(chan int, 1)

	var input string
	fmt.Println(&input)
}
