package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

//writing distributed, network application
//tcp server, http server, rcp are the three common approaches

//creatting a TCP Server in Go
func server() {
	//listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	//accept incoming on port 9999
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConn(c)
	}
}

//handle what is received on the port
func handleServerConn(c net.Conn) {
	//receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

//client
func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hello world"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	fmt.Println("Servers - TCP, HTTP, RCP")
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
