package main

import "fmt"

func main() {
	for _, v := range "kirty" {
		switch v {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Unknown Number")
		}
	}
}
