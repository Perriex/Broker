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

func start(b *Broker) {
	BUFF_COUNT = 5
	broker = Broker{
		clients:  []Client{},
		messages: make(chan Data, BUFF_COUNT),
	}
	
	if len(broker.messages) > 0 {
		go push()
	}

}

func push() {
	for data := range broker.messages {
		for _, client := range broker.clients {
			c, err := rpc.Dial("tcp", "localhost:"+client.Port)

			if err != nil {
				log.Fatal(err)
			}

			var relpy string
			message := "Message recieved"
			err = c.Call("Client.Get", message, &relpy)

			if err != nil {
				log.Fatal(err)
			}
		}
		data.Type.Send()
	}
}
