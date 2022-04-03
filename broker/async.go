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
	port    string
	message string
}

func (m *Memory) Asynchronous(del Delivery, res *string) error {
	*res = "Sent"

	source := ASync{source: del.port}
	data := Data{
		message: del.message,
		_type:   &source,
	}
	if len(broker.messages) == BUFF_COUNT {
		fmt.Println("Message overflow: ", del.message)

	} else {
		broker.messages <- data
	}

	return nil
}

func (_type ASync) Send() {
	client, err := rpc.Dial("tcp", "localhost:"+_type.source)

	if err != nil {
		log.Fatal(err)
	}

	var relpy string
	err = client.Call("Client.Receiver", "Message received", &relpy)

	if err != nil {
		log.Fatal(err)
	}
}
