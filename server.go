package main

import(
	"fmt"
	"strconv"
	"strings"
	"time"
)

func handle(value string){
	sum := 0
	values := strings.Split(value, ",")
	for i := range values {
		v, _ := strconv.Atoi(values[i])
		sum += v
	}
	fmt.Println("Sum : ",value, " is :", strconv.Itoa(sum))
}

func main(){
	Key := "task"
	broker := New()
	for{
		time.Sleep(time.Second)
		taskValue := broker.recieve(Key)
		if len(taskValue) > 0 {
			go handle(taskValue)
		}
	}
}