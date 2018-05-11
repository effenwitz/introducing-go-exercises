package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	//y := [...]string{"a", "b", "c"}
	//fmt.Println(y)
	//sliceEx()
	//mapEx()
	//mapComplex()
	findSmallestNumber()
}

func sliceEx() {
	x := make([]float64, 0)
	x = append(x, 1.5, 2.5)
	fmt.Println(x)
}

func mapEx() {
	x := make(map[string]int)
	x["a"] = 1
	x["b"] = 2
	x["c"] = 3
	x["d"] = 4
	fmt.Println(x)
	// for k, v := range x {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }
	// delete(x, "b")
	// fmt.Println(len(x))
	// if val, ok := x["e"]; ok {
	// 	fmt.Println(val, ok)
	// }

}

func mapComplex() {
	x := make(map[string]map[int]map[string]int)
	y := make(map[string]int)
	y["a"] = 1
	z := make(map[int]map[string]int)
	z[1] = y
	x["b"] = z
	fmt.Println(x)
}
func findSmallestNumber() {
	x := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17}
	sm := x[0]
	for _, v := range x {

		if sm > v {
			sm = v
		}
	}
	fmt.Println(sm)
}
