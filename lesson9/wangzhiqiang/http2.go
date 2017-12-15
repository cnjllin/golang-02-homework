package main

import (
	"net"
	"log"
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
)

var content = `HTTP/1.1 200 OK
Date: Sat, 9 Dec 2017 17:02:03 GMT
Content-Type: text/html
Content-Length: %d
Connection: Keep-Alive
Server: reboot


`
var imgHeader = `HTTP/1.1 200 OK
Date: Sat, 9 Dec 2017 17:02:03 GMT
Content-Type: img/jpg
Content-Length: %d
Connection: Close
Server: reboot


`

func handleConn(conn net.Conn){
	bufReader := bufio.NewReader(conn)
	getLine,_,_ := bufReader.ReadLine()
	fmt.Println(getLine)
	sliceLine := strings.Split(string(getLine)," ")
	fmt.Println(sliceLine)
	fmt.Println(len(sliceLine[1]))
	if len(sliceLine[1]) <= 2 {
		var htmlBody = `<h1 style="color:red">hello golang</h1>`
		imgDir, _ := os.Open("/Users/knowbox/Downloads/img")
		defer imgDir.Close()
		imgs,_ := imgDir.Readdirnames(-1)
		for _,img := range imgs {
			if strings.Contains(img,".jpg"){
				fmt.Println(img)
				htmlBody += fmt.Sprintf(`<img src="/Users/knowbox/Downloads/img/%s"></br>`, img)
			}
		}
		fmt.Println(htmlBody)
		conn.Write([]byte(fmt.Sprintf(content,len(htmlBody))))
		fmt.Println(len(htmlBody))
		conn.Write([]byte(htmlBody))
		fmt.Println(htmlBody)
	} else {
		url := "." + sliceLine[1]
		img, err := os.Open(url)
		if err != nil {
			return
		}
		defer img.Close()
		imgInfo, err := img.Stat()
		if err != nil {
			return
		}
		conn.Write([]byte(fmt.Sprintf(imgHeader,imgInfo.Size())))
		io.Copy(conn,img)
	}
	conn.Close()
}
func main(){
	listener, err := net.Listen("tcp",":8021")
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn,err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
