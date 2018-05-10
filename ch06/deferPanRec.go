package main

import "fmt"

func main() {
	fmt.Println("Defer panic recover")
	recoverFrmPan()
	fmt.Println("can this print")
}

func recoverFrmPan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in recoverFrmPan", r)
		}
	}()
	testPanic()
}

func testPanic() {
	panic("i panic")
}
