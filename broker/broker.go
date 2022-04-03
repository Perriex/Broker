package broker

import (
	"log"
	"net/rpc"
)

type Memory int

var BUFF_COUNT int
var broker Broker

type Client struct {
	port string
}
type Source interface {
	Send()
}
type Data struct {
	message string
	_type   Source
}

type Reply struct {
	message string
}


func (m *Memory) Subscribe(client string, res *string) error {
	c := new(Client)
	c.port = client
	broker.clients = append(broker.clients, *c)
	*res = "client subscribed"
	return nil
}

type Broker struct {
	clients  []Client
	messages chan Data
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
		for _, client := range broker.clients {
			c, err := rpc.Dial("tcp", "localhost:"+client.port)

			if err != nil {
				log.Fatal(err)
			}

			var relpy string
			message := data.message
			err = c.Call("Client.Get", message, &relpy)

			if err != nil {
				log.Fatal(err)
			}
		}
		data._type.Send()
	}
}
