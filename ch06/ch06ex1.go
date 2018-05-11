package main

import (
	"fmt"
)

func main() {
	// i, j := retnMulti()
	// fmt.Printf("%d,%d\n", i, j)
	// fmt.Println(f2())
	// //fmt.Println(add(1, 2, 3, 4))
	// xs := make([]int, 0)
	// xs = append(xs, 1, 2, 3, 4, 5)
	// fmt.Println(add(xs...))
	// ys := []int{1, 2, 2, 2, 2} //slice
	// fmt.Println(add(ys...))
	//closureEx()

	// nextEven := makeEvenNumberGen()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	fmt.Println(factorialRec(10))
}

//multiple values can be retruned
//required to return err
//x, err:=f()
//or boolean to indicate success
//x, ok:=f()
func retnMulti() (int, int) {
	return 2, 3
}

//return values can have names
func f2() (r int) {
	r = 1
	return
}

//variadic functions
func add(args ...int) (total int) {
	total = 0
	for _, v := range args {
		total += v
	}
	return
}

//closure - functions inside function
func closureEx() {
	total := 0
	avg := func(args ...int) int {
		for _, v := range args {
			total += v
		}
		return total / len(args)
	}
	zs := []int{1, 2, 3, 4, 5}
	fmt.Println(avg(zs...))
}

//function returns another function
func makeEvenNumberGen() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

//recursion
//uint-unsigned means only positive intgers
func factorialRec(x uint) uint {

	if x == 0 {
		return 1
	}
	return x * factorialRec(x-1)
}
