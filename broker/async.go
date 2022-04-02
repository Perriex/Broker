package broker

import (
	"log"
	"net/rpc"
)

type ASync struct {
	source string
}

func CallAsync(src string) *ASync {
	_type := ASync{source: src}
	return &_type
}

func (_type ASync) Send(){
	client, err := rpc.Dial("tcp", _type.source)

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	err = client.Call("Client.Get", "Message received", &relpy)

	if err != nil {
		log.Fatal(err)
	}
}