package broker

import (
	"log"
	"net/rpc"
	"sync"
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
	wg       sync.WaitGroup
}

func (m *Memory) Subscribe(client string, res *string) error {
	c := new(Client)
	c.Port = client
	broker.clients = append(broker.clients, *c)
	*res = "client subscribed"
	broker.wg.Done()
	return nil
}

func init() {
	broker.wg.Wait()
	start(&broker)
}

func start(b *Broker) {
	BUFF_COUNT = 5
	broker = Broker{
		clients:  []Client{},
		messages: make(chan Data, BUFF_COUNT),
	}

	broker.wg.Add(1)
	go push(&broker.wg)
}

func push(wg *sync.WaitGroup) {
	wg.Wait()
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
