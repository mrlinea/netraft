package main

import (
	"fmt"
	"github.com/netraft"
)

func main() {
	n := netraft.CreateNetwork("toto")
	fmt.Println(n.Name())
}
