package broker

import (
	"fmt"
	"log"
	"net/rpc"
)

type ASync struct {
	source string
}

type Delivery struct {
	Port    string
	Message string
}

func (m *Memory) Asynchronous(del Delivery, res *string) error {
	*res = "Sent"
	println("message recieved")
	source := ASync{source: del.Port}
	data := Data{
		Message: del.Message,
		Type:    &source,
	}
	if len(broker.messages) == BUFF_COUNT {
		fmt.Println("Message overflow: ", del.Message)

	} else {
		broker.messages <- data
	}
	println("len recieved")

	return nil
}

func (_type ASync) Send() {
	client, err := rpc.Dial("tcp", _type.source)

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	err = client.Call("Receiver.Get", "Message received", &relpy)

	if err != nil {
		log.Fatal(err)
	}
}
