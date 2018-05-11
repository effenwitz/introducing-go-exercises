package main

import "fmt"

//pointers - represented by asterisk *
//pointers are useful when paired with structs.
func main() {
	x := 5
	//zero(x)
	fmt.Println(x)
	//& used to find address of variable
	zeroPtr(&x)
	fmt.Println(x)

	//new - another way to get a pointer
	xPtr := new(int)
	zeroPtr(xPtr)
	fmt.Println(*xPtr)
}

func zero(x int) {
	x = 0
}
func zeroPtr(xptr *int) {
	*xptr = 1 //dereferencing
}
