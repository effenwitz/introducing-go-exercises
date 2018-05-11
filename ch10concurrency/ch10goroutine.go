package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func main is also a go routine

func main() {
	//goroutine is a function that is capable of running concurrently with other functions

	fmt.Println("Go Concurrency and GoRoutine")
	//to create a goroutine, we use the keyword go
	//with goroutine we immidiately return to the next line and dont wait for the function to complete.
	//go gortn(0)
	//creating multiple go routines
	for i := 0; i < 10; i++ {
		go gortn(i)
	}
	var input string
	fmt.Scanln(&input)
}

func gortn(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
