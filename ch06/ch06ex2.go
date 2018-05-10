package main

import (
	"fmt"
	"os"
)

//defer, panic, recover
//defer - schedules a function call to be run after the function completes
//defer use it when resources needs to be freed. such as closing of IO operations - open a file, stream etc
//defer is very much similar to Java finally block.
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {

	defer second() //defers moves the call to second to the end of function main
	first()
	panRec()

}
func fileOpen() int {
	f, _ := os.Open("test.txt")
	defer f.Close()
	//keeps close call near open, so easier to understand
	//defer will run even if runtime exception occurs.
	if 1 > 2 {
		fmt.Println("yes")
	}
	//close will happen before return statment
	return 1
}

//panic and recover
//panic - generate runtime error
//recover - handle a runtime panic
func panRec() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("this is panic")
}
