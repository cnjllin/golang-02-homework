package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

// 进程、线程、协程（由重到轻）
// 进程、协程、线程 （历史出现次序）