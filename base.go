package main

import (
	"log"
	"net"
	"net/rpc"
	"broker/broker"
)

func main() {

	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8080")

	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	service := new(broker.Memory)
	rpc.Register(service)
	rpc.Accept(inbound)
}
