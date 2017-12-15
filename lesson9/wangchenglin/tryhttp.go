package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var content = `HTTP/1.1 200 OK
Date: Sat, 9 Dec 2017 17:02:03 GMT
Content-Type: text/html
Content-Length: %d
Connection: Keep-Alive
Server: reboot

`
var imageHeader = `HTTP/1.1 200 OK
Date: Sat, 9 Dec 2017 17:02:03 GMT
Content-Type: image/jpg
Content-Length: %d
Connection: Close
Server: reboot

`
var i = 0

func httpHandler(conn net.Conn) {
	i++
	fmt.Println("i:", i)

	bufReader := bufio.NewReader(conn)
	getLine, _, _ := bufReader.ReadLine()
	fmt.Println("getLine:", getLine)
	sliceLine := strings.Split(string(getLine), " ")
	fmt.Println("sliceLine:", sliceLine)
	if len(sliceLine[1]) <= 2 {

		var htmlBody = `<h1 style="color:red">hello golang</h1>`
		imageDir, _ := os.Open("./image")
		defer imageDir.Close()

		images, _ := imageDir.Readdirnames(-1)
		for _, image := range images {
			fmt.Println("image:", image)
			if strings.Contains(image, ".jpg") {
				htmlBody += fmt.Sprintf(`<image src="/image/%s"></br>`, image)
				fmt.Println("htmlBody:", htmlBody)
			}
		}
		conn.Write([]byte(fmt.Sprintf(content, len(htmlBody))))
		conn.Write([]byte(htmlBody))
	} else {
		fmt.Println("else case...")
		url := "." + sliceLine[1]
		image, err := os.Open(url)
		if err != nil {
			return
		}
		defer image.Close()
		imageInfo, err := image.Stat()
		if err != nil {
			return
		}

		conn.Write([]byte(fmt.Sprintf(imageHeader, imageInfo.Size())))
		io.Copy(conn, image)
	}
	conn.Close()
}

func main() {
	tcpListener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("err")
	}
	defer tcpListener.Close()

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			log.Fatal("err")
		}

		go httpHandler(conn)
	}
}
