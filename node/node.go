package node

import "crane_core/server"

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
