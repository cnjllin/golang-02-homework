package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

// client -> server: ls | grep go\n
// server -> client: content of command EOF
func handle(conn net.Conn) {
	log.Printf("connection %s", conn.RemoteAddr())
	defer conn.Close()
	//r := bufio.NewReader(conn)
	//cmdstr, _ := r.ReadString('\n')
	cmd := exec.Command("bash")
	fd, err := pty.Start(cmd)
	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.Create("help.sys")
	defer f.Close()
	go io.Copy(fd, conn)
	io.Copy(conn, io.TeeReader(fd, f))
}

func main() {
	l, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := l.Accept()
		handle(conn)
	}
}
