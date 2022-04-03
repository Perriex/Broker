package broker

import (
	"log"
	"net/rpc"
	"sync"
)

type Memory int

var BUFF_COUNT = 5
var CLIENT_NUM = 1
var broker Broker

type Source interface {
	Send()
}

type Data struct {
	Message string
	Type    Source
}

type Broker struct {
	clients  []string
	messages chan Data
	wg       sync.WaitGroup
}

func (m *Memory) Subscribe(client string, res *string) error {
	defer broker.wg.Done()
	broker.clients = append(broker.clients, client)
	*res = "client subscribed"
	return nil
}

func init() {
	broker.wg.Wait()
	broker = Broker{
		clients:  []string{},
		messages: make(chan Data, BUFF_COUNT),
	}

	broker.wg.Add(CLIENT_NUM)
	go push(&broker.wg)
}


func push(wg *sync.WaitGroup) {
	wg.Wait()
	for data := range broker.messages {
		for _, client := range broker.clients {
			c, err := rpc.Dial("tcp", "0.0.0.0:"+client)
			if err != nil {
				log.Fatal(err)
			}

			var relpy string
			err = c.Call("Receiver.Get", data.Message, &relpy)

			if err != nil {
				log.Fatal(err)
			}
		}
		data.Type.Send()
	}
}
