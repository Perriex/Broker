package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Receiver int

var BROKER_PORT = "8080"

type Delivery struct {
	port    string
	message string
}

func (client *Receiver) Get(message string, reply *string) error {
	fmt.Println("Message: " + message)

	return nil
}

func main() {
   go start()

	client, err := rpc.Dial("tcp", "0.0.0.0:8081")

	if err != nil {
		log.Fatal(err)
	}

	del := Delivery{
		port:    "0.0.0.0:" + BROKER_PORT,
		message: "Hello, world!",
	}
	var relpy string

	err = client.Call("Memory.Asynchronous", del, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	for {

	}
}

func start() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	receiver := new(Receiver)
	rpc.Register(receiver)
	rpc.Accept(inbound)

}
