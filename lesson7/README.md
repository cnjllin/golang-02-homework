## homework 7

    如下示例是一个RPC调用，server能执行client发来的命令并返回执行结果。
    请把客户端改造成可以并发向多个server发送命令并返回结果的模式

RPC客户端
```golang
package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
	"fmt"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:42586")
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
		err = client.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}
}
```

RPC服务端
```golang
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
```
