package main
import (
	"time"
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main(){
	context,_ := zmq.NewContext()
	defer context.Term()

	socket,_ := context.NewSocket(zmq.DEALER)
	defer socket.Close()
	socket.SetIdentity("123")
	socket.Connect("tcp://127.0.0.1:9019")
	socket.Send("hello",0)
	poller := zmq.NewPoller()
	poller.Add(socket, zmq.POLLIN)

	for{
	        //fmt.Println("1213")		
		sockets, _ := poller.Poll(500 * time.Millisecond)
		for _,socketm := range sockets{
			switch s:=socketm.Socket;s{
			case socket:
				content,_ := socket.Recv(0)
				fmt.Println(content)
			}
		}
	}
}
