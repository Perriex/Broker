package broker

import (
	"fmt"
	"sync"
)

type Sync struct{
	wg sync.WaitGroup
}

func (m *Memory) Synchronous(message string, res *string) error {
	*res = "Sent"

	source := Sync{}
	source.wg.Add(1)

	data := Data{
		message: message,
		_type:   &source,
	}

	if len(broker.messages) == BUFF_COUNT {
		fmt.Println("Message overflow: ", message)

	} else {
		broker.messages <- data
	}
	source.wg.Wait()

	return nil
}


func (_type *Sync) Send(){
	_type.wg.Done()
}