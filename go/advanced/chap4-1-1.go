package advanced

import (
	"log"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	// register the service to the HelloService
	rpc.RegisterName("HelloService", new(HelloService))

	// create TCP transport link
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	// serve the service via the TCP conn
	rpc.ServeConn(conn)
}

// ClientSide
func client() {
	// create the link via tcp transport
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	// call a RPC service with the parameters
	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal("call error:", err)
	}
	fmt.Println(reply)

}
