#homework-9

继续完成自己的http图片浏览服务器，访问http://127.0.0.1:8021/ 能返回一个img文件夹里的webp格式的图片墙

参考代码如下：
```golang
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
	"github.com/auxten/logrus"
)

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
	logrus.Debug(sliceLine)
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
		img, err := os.Open(url)
		if err != nil {
			return
		}
		defer img.Close()
		imgInfo, err := img.Stat()
		if err != nil {
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
```
