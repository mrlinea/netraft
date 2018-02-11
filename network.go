package netraft

import ()

type Network interface {
	Name() string
}

type network struct {
	_name string
}

func (n *network) Name() string {
	return n._name
}

func CreateNetwork(name string) Network {
	n := network{_name: name}
	return &n
}
