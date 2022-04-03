package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"
)

type Receiver int

var BROKER_PORT = "8080"
var SERVER_PORT = "8081"

type Delivery struct {
	Port    string
	Message string
}

func (r *Receiver) Get(message string, reply *string) error {
	fmt.Println("Message: " + message)

	return nil
}

func main() {
	go start()

	client, err := rpc.Dial("tcp", "0.0.0.0:8080")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to Broker ...")

	for {
		fmt.Println("Enter message: ")
		in := bufio.NewReader(os.Stdin)

		message, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		del := Delivery{
			Port:    "0.0.0.0:8081",
			Message: string(message),
		}
		var relpy string

		err = client.Call("Memory.Asynchronous", del, &relpy)

		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1000)
	}

	// for testing sync process
	// err = client.Call("Memory.Synchronous", "hey", &relpy)
}

// start server itself
func start() {
	fmt.Println("Starting Server on port 8081 ...")
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+SERVER_PORT)
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
