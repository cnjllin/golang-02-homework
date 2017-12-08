package main

import (
	"net/rpc"
	"log"
	"bufio"
	"os"
	"fmt"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:10001")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		var reply string
		err = client.Call("Listener.GetLine1", string(line), &reply)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("REPLY: %s\n", reply)
	}

}