package main
import (
   "log"
   "net"
   "net/rpc"
   "fmt"
)

type Reply struct {
   Data string
}

type Receiver int

func (r *Receiver) Send(data []byte, reply *Reply) error {

   rv := string(data)
   fmt.Printf("%v\n", rv)
   *reply = Reply{rv}
   return nil

}

func main() {

   // start server
   addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
   if err != nil {
     log.Fatal(err)
   }

   // listen 
   inbound, err := net.ListenTCP("tcp", addy)
   if err != nil {
     log.Fatal(err)
   }

   // publish receiver
   receiver := new(Receiver)
   rpc.Register(receiver)
   rpc.Accept(inbound)

   memory, err = rpc.Dial("tcp","0.0.0.0:8080")

   if err != nil {
      log.Fatal(err)
   }

   // send data to broker
}