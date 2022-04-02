package broker

import (
	"log"
	"net/rpc"
)

var CLIENT_COUNT int
var broker Broker

type Client struct {
	port string
}

type Broker struct {
	clients  []Client
	messages chan Data
}

type Data struct {
	message string
	_type  Source
}

type Source interface {
	Send()
}

func start(b *Broker) {
	broker = Broker{
		[]Client{},
		make(chan Data, 5),
	}
	CLIENT_COUNT = 3

	for client := 0; client < CLIENT_COUNT; client++ {
		go push()
	}
}

func push() {
	for data := range broker.messages {
		for _, client := range broker.clients {
			hook, err := rpc.Dial("tcp", client.port)
			if err != nil {
				log.Fatal(err)
			}

			var relpy string
			message := data.message
			err = hook.Call("Client.Get", message, &relpy)

			if err != nil {
				log.Fatal(err)
			}
		}
		data._type.Send()
	}
}

type Memory int

func (m *Memory) Subscribe(client string, res *string) error {
	hook := Client{client}
	broker.clients = append(broker.clients, hook)
	*res = "added"
	return nil
}

func (m *Memory) Synchronous(message string, res*string) error{
	source := CallSync()
	data := Data{
		message,
		source,
	}
	broker.messages <- data
	source.Patient()

	*res = "sent"

	return nil
}

type Delivery struct{
	port string
	message string
}

func (m *Memory) Asynchronous(del Delivery, res*string) error{
	source := CallAsync(del.port)
	data := Data{
		del.message,
		source,
	}
	broker.messages <- data

	*res = "sent"

	return nil
}
