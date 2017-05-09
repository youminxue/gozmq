package main
import(
	"fmt"
	zmq "github.com/pebbe/zmq4"
)
func main(){
	fmt.Println("123")
	context,_ := zmq.NewContext()
	defer context.Term()
	xpub,_ := context.NewSocket(zmq.XPUB)
	defer xpub.Close()
	xsub,_ := context.NewSocket(zmq.XSUB)
	defer xsub.Close()

	xpub.Bind("tcp://127.0.0.1:9014")
	xsub.Bind("tcp://127.0.0.1:9015")
	zmq.Proxy(xsub,xpub,nil)
}

