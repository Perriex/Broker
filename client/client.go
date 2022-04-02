package main

import (
	"broker/broker"
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"sync"
)

type Reply struct {
	message string
}

type Client int

func (client *Client) Get(message string, reply *Reply) error {
	fmt.Println("Message: " + message)

	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)

	port, _, err := in.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	//client, err := rpc.Dial("tcp", "localhost:"+string(port))

	client, err := rpc.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)

	go get(wg, string(port))

	var relpy string

	err = client.Call("Memory.Subscribe", port, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	del := broker.Delivery{
		port:    string(port),
		message: "Hello, world!",
	}
	err = client.Call("Memory.Asynchronous", del, &relpy)

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}

func get(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	addy, err := net.ResolveTCPAddr("tcp", "localhost:"+port)

	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	client := new(Client)
	rpc.Register(client)
	rpc.Accept(inbound)
}
