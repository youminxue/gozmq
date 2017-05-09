package main
import(
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main(){
	context,_ := zmq.NewContext()
	defer context.Term()

	socket,_ := context.NewSocket(zmq.SUB)
	defer socket.Close()
	socket.Connect("tcp://127.0.0.1:9014")
	socket.SetSubscribe("12")
	for{
		content,_ := socket.Recv(0) 
		fmt.Println(content)
		content,_ = socket.Recv(0) 
		fmt.Println(content)
	}
}
