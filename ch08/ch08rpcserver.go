package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//net/rpc and net/rpc/jsonrpc provide and easy way to expose methods so they can be invoked over a network
//create a struct

//Server is a struct
type Server struct{}

//write a method with Server as reciever

//Negate is a method
func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

//write the server
func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

//write the client
func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999)= ", result)
	}
}

//call the server and client as go subroutines
func main() {
	fmt.Println("RPC Server")
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}
