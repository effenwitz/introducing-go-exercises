package main

import (
	"fmt"
)

func main() {
	if err := test(); err != nil {
		fmt.Println("Main Error:", err)
	}
}

func catchPanic(err error, functionName string) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

		if err != nil {
			err = fmt.Errorf("%v", r)
		}
	} else if err != nil {
		fmt.Printf("%s : ERROR : %v\n", functionName, err)
	}
}

func mimicError(key string) error {
	return fmt.Errorf("Mimic Error : %s", key)
}

func test() (err error) {
	defer catchPanic(err, "Test")
	fmt.Println("Start Test")

	err = mimicError("1")

	fmt.Println("End Test")
	return err
}
