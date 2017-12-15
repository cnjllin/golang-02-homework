package main

import (
	"log"
	"net"
	"os"
	"fmt"
	"io"
	//"github.com/toolkits/file"
	"bufio"
	"strings"
)

var error_404=`HTTP/1.1 404 Not Found
Server: nginx/1.6.1
Date: Thu, 14 Dec 2017 12:49:18 GMT
Content-Type: text/html; charset=UTF-8
Content-Length: 69
Expires: 0
Cache-Control: no-cache
Proxy-Connection: Keep-alive

<html><title>404: Not Found</title><body>404: Not Found</body></html>
`

var error_500=`HTTP/1.1 500 Internal Server Error
Date: Thu, 14 Dec 2017 13:09:59 GMT
Content-Length: 93
Content-Type: text/html; charset=UTF-8

<html><title>500: Internal Server Error</title><body>500: Internal Server Error</body></html>
`

var mainContent = `HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:18:23 GMT
Content-Type: text/html
Content-Length: %d
Connection: Close
Server: reboot

`
var imgHeader = `HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:18:23 GMT
Content-Type: image/webp
Content-Length: %d
Connection: Close
Server: reboot

`
func handleConn(conn net.Conn) {
	bufReader := bufio.NewReader(conn)
	getLine, _, _ := bufReader.ReadLine()
	sliceLine := strings.Split(string(getLine), " ")
	fmt.Println("sliceLine:",sliceLine)
	fmt.Println("sliceLine[1]:",sliceLine[1])
	fmt.Println("len(sliceLine[1]):",len(sliceLine[1]))
	if len(sliceLine[1]) <= 2 {

		var htmlBody = `<h1 style="color:red">hello golang</h1>`
		imgDir, _ := os.Open("./img")
		defer imgDir.Close()
		imgs, _ := imgDir.Readdirnames(-1)
		for _, img := range imgs {
			if strings.Contains(img, ".webp") {
				htmlBody += fmt.Sprintf(`<img src="/img/%s"></br>`, img)
			}
		}
		conn.Write([]byte(fmt.Sprintf(mainContent, len(htmlBody))))
		conn.Write([]byte(htmlBody))
	} else {
		url := "." + sliceLine[1]
		fmt.Println("url:",url)
		img, err := os.Open(url)
		if err != nil {
			conn.Write([]byte(error_404))

			return
		}
		defer img.Close()
		imgInfo, err := img.Stat()
		if err != nil {
			conn.Write([]byte(error_500))

			return
		}

		conn.Write([]byte(fmt.Sprintf(imgHeader, imgInfo.Size())))
		io.Copy(conn, img)
	}
	conn.Close()

}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}