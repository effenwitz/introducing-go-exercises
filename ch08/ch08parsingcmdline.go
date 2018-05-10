package main

import (
	"flag"
	"fmt"
	"math/rand"
)

//flag package allows us to parse cmd line args
//additional non-flag arguments can be returned by flag.Args[]
func main() {
	fmt.Println("Parsing command line args")

	maxp := flag.Int("max", 6, "the max value")
	flag.Parse()
	fmt.Println(rand.Intn(*maxp))
}
