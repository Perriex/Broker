package broker

import(
	"sync"
)

type Sync struct{
	wg sync.WaitGroup
}

func CallSync() *Sync{
	_type := Sync{}
	_type.wg.Add(1)
	return &_type
}

func (_type *Sync) Patient(){
	_type.wg.Wait()
}

func (_type *Sync) Send(){
	_type.wg.Done()
}