package main

import (
	"fmt"
	"log"
	"net"
	//"time"
)

//为什么通过浏览器访问不到？使用telnet或curl可以请求到内容。不是http协议？

func handle(conn net.Conn) { //数据类型不是一般只有int,string,float等，为什么net.Conn可以作为数据类型？
	fmt.Fprintf(conn, "%s", "<p>test c10k</p>")
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

