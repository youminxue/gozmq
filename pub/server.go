package main
import(
	"time"
	"fmt"
	zmq "github.com/pebbe/zmq4"
)
func main(){
	context,_ := zmq.NewContext()
	defer context.Term()
	socket,_ := context.NewSocket(zmq.XPUB)
	defer socket.Close()
	socket.Bind("tcp://127.0.0.1:9014")
	for{
		var itype,info string
		fmt.Println("Please input your full name: ")
		fmt.Scanln(&itype, &info)
		socket.Send(itype, zmq.SNDMORE)
		socket.Send(info, 0)							   }

	poller := zmq.NewPoller()
	poller.Add(socket, zmq.POLLIN)
	for{
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
