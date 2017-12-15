package main

import (
	"net"
	"log"
	"bufio"
	"os"
	"strings"
	"io"
	"fmt"
)
/*
3. ftp客户端 **作业**
// ioutil.ReadAll(conn) 一次性，大文件不适合
io.Copy(f, conn)
*/

func getFile(conn net.Conn, method string, path string) {
	// 2.发送数据
	// conn.Write 传输的是 []byte
	//conn.Write(cmdLine)
	// fmt.Fprintf 传输的是string
	fmt.Fprintf(conn, method+" "+path+"\n")

	// 3.接收数据
	// 打印标准输出
	//io.Copy(os.Stdout, conn)

	// 存入文件
	pathList := strings.Split(path, "/")
	fileName := pathList[len(pathList)-1]
	f, err := os.Create(fileName)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()
	log.Printf("Receving file: %s", fileName)
	io.Copy(f, conn)
	log.Printf("File '%s' recevied completly", fileName)
}

func main() {
	// 1. 创建连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	// 4. 关闭连接
	defer conn.Close()

	in := bufio.NewReader(os.Stdin)
	for {
		quit := false
		var method, path string

		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		//log.Print(string(line))
		// 拆分method和path
		cmdLine := string(line)
		fs := strings.Fields(cmdLine)

		switch len(fs) {
		case 0:
			continue
		case 1:
			method = fs[0]
		case 2:
			method = fs[0]
			path = fs[1]
		default:
			log.Print("args number ERROR!!")
			continue
		}
		log.Printf("method: %s", method)

		// 根据method选择对应操作的函数
		switch strings.ToUpper(method) {
		case "QUIT":
			log.Print("Quit Now!!!!")
			quit = true
		case "GET":
			getFile(conn, method, path)
		default:
			log.Printf("ERROR: %s", string(line))
		}

		if quit {
			break
		}
	}
}
