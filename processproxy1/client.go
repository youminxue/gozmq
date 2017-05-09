package main
import(
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main(){
	context,_ := zmq.NewContext()
	defer context.Term()

	socket,_ := context.NewSocket(zmq.PUB)
	defer socket.Close()

	socket.Connect("tcp://127.0.0.1:9015")
	for{
		var itype,info string
		fmt.Println("Please input your full name: ")
		fmt.Scanln(&itype, &info)
		socket.Send(itype, zmq.SNDMORE)
		socket.Send(info, 0)
	}
}
