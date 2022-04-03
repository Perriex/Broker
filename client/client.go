package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"sync"
)

var BROKER_PORT = "8080"

func main() {
	fmt.Println("Enter requested port: ")
	in := bufio.NewReader(os.Stdin)

	port, _, err := in.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)

	go start(wg, string(port))

	fmt.Println("Receiver from Broker ...")
	client, err := rpc.Dial("tcp", "0.0.0.0:"+BROKER_PORT)

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	err = client.Call("Memory.Subscribe", string(port), &relpy)

	println(relpy)

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}

type Receiver int

type Delivery struct {
	Port    string
	Message string
}

func (r *Receiver) Get(message string, reply *string) error {
	fmt.Println("Message: " + message)

	return nil
}


// start client itself
func start(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	fmt.Println("Starting Client on port: "+port+" ...")
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+string(port))

	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	service := new(Receiver)
	rpc.Register(service)
	rpc.Accept(inbound)
}
