package main

import (
	"fmt"

	"github.com/effenwitz/introducing-go-exercises/ch08package/math"
)

func main() {
	fmt.Println("Creating packages")
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
