package broker

import (
	"log"
	"net/rpc"
)

type Memory int

var BUFF_COUNT int
var broker Broker

type Client struct {
	Port string
}

type Source interface {
	Send()
}

type Data struct {
	Message string
	Type    Source
}

type Broker struct {
	clients  []Client
	messages chan Data
}

func (m *Memory) Subscribe(client string, res *string) error {
	c := new(Client)
	c.Port = client
	broker.clients = append(broker.clients, *c)
	*res = "client subscribed"
	return nil
}

func init() {
	start(&broker)
}

func start(b *Broker) {
	BUFF_COUNT = 5
	broker = Broker{
		clients:  []Client{},
		messages: make(chan Data, BUFF_COUNT),
	}

	go push()
}

func push() {
	for data := range broker.messages {
		println("message " + data.Message)
		for _, client := range broker.clients {
			c, err := rpc.Dial("tcp", "0.0.0.0:"+client.Port)
			println("client is sent " + client.Port)
			if err != nil {
				log.Fatal(err)
			}

			var relpy string
			err = c.Call("Receiver.Get", "Message recieved", &relpy)

			if err != nil {
				log.Fatal(err)
			}
		}
		data.Type.Send()
	}
}
