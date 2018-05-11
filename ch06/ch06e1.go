package main

import (
	"fmt"
)

func main() {
	// x := []int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("Sum %d\n", sumSlice(x))
	// fmt.Println(halfInt(5))
	// fmt.Printf("Greatest %d\n", greatestNum(x...))
	// makeOdd := makeOddGen()
	// fmt.Println(makeOdd())
	// fmt.Println(makeOdd())
	// fmt.Println(makeOdd())
	// deferEx()
	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	x := 1.5
	fmt.Println(&x)
	fmt.Println(swappy(1, 2))
}

//closure example
func makeOddGen() func() int {
	x := 1
	tryFunc := func() (ret int) {
		ret = x
		x += 2
		return
	}
	return tryFunc
}

//takes a slice as an input param
func sumSlice(x []int) (total int) {
	total = 0
	for _, v := range x {
		total += v
	}
	return
}

//takes var args as inout param
func greatestNum(args ...int) int {
	greatest := 0
	for _, v := range args {
		if greatest < v {
			greatest = v
		}
	}
	return greatest
}

func halfInt(x uint) (uint, bool) {

	if x%2 == 0 {
		return x / 2, true
	}
	return 1, false
}

func deferEx() {
	defer fmt.Println("last to print")
	fmt.Println("1st to print")
}

//fnext i = fcurrent i + fcurrent i - 1
//and
//fnext i - 1 = fcurrent i - 1
// fibonacci returns a function that returns
// successive fibonacci numbers from each
// successive call
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func square(x *float64) {
	*x = *x * *x
}

func swappy(x, y int) (int, int) {
	x, y = y, x
	return x, y
}
