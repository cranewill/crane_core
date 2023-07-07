package node

import (
	"crane_core/server"
	"fmt"
	"net"
)

type Node struct {
	Name string
	Port int

	Servers map[string]*server.GameServer
}

func NewNode(name string, port int) *Node {
	return &Node{
		Name: name,
		Port: port,
	}
}

func (n *Node) Init() error {
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", n.Port))
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go n.Serve(conn)
	}
}

func (n *Node) Serve(conn net.Conn) {

}
