package main

import "fmt"

func main() {
	fmt.Println("Enter a value to change to Celsius")
	var farVal float64
	fmt.Scanf("%f", &farVal)
	output := (farVal - 32) * 5 / 9
	fmt.Printf("Temp in Celcius %v\n", output)

}
