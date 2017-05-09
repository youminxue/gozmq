package main
import(
	"time"
	"fmt"
	zmq "github.com/pebbe/zmq4"
)
func main(){
	fmt.Println("123")
	context,_ := zmq.NewContext()
	defer context.Term()
	context.SetIoThreads(8)
	xpub,_ := context.NewSocket(zmq.XPUB)
	defer xpub.Close()
	xsub,_ := context.NewSocket(zmq.XSUB)
	defer xsub.Close()

	xpub.Bind("tcp://127.0.0.1:9014")
	xsub.Bind("tcp://127.0.0.1:9015")
	//subMsg := []byte{1}
	//xsub.SendBytes([]byte("12"), 0)
	poller := zmq.NewPoller()
	poller.Add(xsub, zmq.POLLIN)
	poller.Add(xpub, zmq.POLLIN)
	for {
		sockets, _ := poller.Poll(500 * time.Millisecond)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
				case xpub:
					content,_ := xpub.RecvBytes(0)
					//fmt.Println(content)
					xsub.SendBytes(content, 0)
				case xsub:
					content,_ := xsub.Recv(0)
					fmt.Println(content)
					xpub.Send(content, zmq.SNDMORE)
					content,_ = xsub.Recv(0)
					fmt.Println(content)
					xpub.Send(content, 0)
				}
			}
	}
}
