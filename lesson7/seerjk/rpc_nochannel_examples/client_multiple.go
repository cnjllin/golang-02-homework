package main

import (
	"fmt"
	"log"
	"net/rpc"
	"bufio"
	"os"
)

type Client struct {
	ip string
	port string
}

func clientProcess(c Client, cmd string) {
	client, err := rpc.Dial("tcp", c.ip+c.port)
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	// cmd reply 类型要和server段相同
	err = client.Call("Listener.GetLine1", cmd, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s REPLY: %s\n", c, reply)
	client.Close()
}

func main() {
	serverList := []Client {
		Client{
			ip: "localhost",
			port: ":10001",
		},
		Client{
			ip: "127.0.0.1",
			port: ":10002",
		},
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range serverList {
			go clientProcess(c, string(line))
		}
	}

}
