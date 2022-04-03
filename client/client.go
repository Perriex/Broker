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
	client, err := rpc.Dial("tcp", "localhost:"+BROKER_PORT)

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	err = client.Call("Memory.Subscribe", port, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}

func start(wg *sync.WaitGroup,port string) {
	defer wg.Done()

	fmt.Println("Starting Client...")
	addy, err := net.ResolveTCPAddr("tcp", "localhost:"+string(port))

	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	rpc.Accept(inbound)
}
