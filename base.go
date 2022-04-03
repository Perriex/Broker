package main

import (
	"broker/broker"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

var BROKER_PORT = "8080"

func main() {
	fmt.Println("Welcome abroad! starting in port " + BROKER_PORT)

	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+BROKER_PORT)

	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)

	if err != nil {
		log.Fatal(err)
	}

	b := new(broker.Memory)
	rpc.Register(b)
	rpc.Accept(inbound)
}
