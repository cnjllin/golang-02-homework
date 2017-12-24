package main

import (
	"fmt"
	"log"
	"net"
	"io"
	"os"
	"bufio"
	"strings"
	"github.com/auxten/logrus"
	"net/url"
	"time"
)

var rootHead = `HTTP/1.1 200 OK
Date: %s
Content-Type: text/html; charset=utf-8
Content-Length: %d
Connection: Close
Server: RebootServer

`
// head 和  body界线 /r/n/r/n 多一个空行
// Date: Sat, 29 Jul 2017 06:11:23 GMT
var imgHead = `HTTP/1.1 200 OK
Date: %s
Content-Type: image/%s
Content-Length: %d
Connection: Close
Server: RebootServer

`

func getHttpNowDateStr() string {
	tm := time.Now()
	return tm.Format("Mon, 02 Jan 2006 03:04:05 GMT")
}


func handleRouteConn(conn net.Conn) {
	// 路由处理层
	bufReader := bufio.NewReader(conn)
	firstLine, _, err := bufReader.ReadLine()
	if err != nil {
		logrus.Error(err)
		return
	}

	// GET /img/abc.webp HTTP/1.1
	sliceLine := strings.Split(string(firstLine), " ")
	logrus.Debug(sliceLine)
	slicePath := strings.Split(sliceLine[1], "/")
	// "/"               0: ""  1:""
	// "/img/abc.webp"   0: ""  1:"img"  2:"abc.webp"
	var path string
	var imgPath string

	if slicePath[1] == "" {
		path = "/"
	} else if slicePath[1] == "img" {
		path = "/img/"
		imgPath = sliceLine[1]
	}
	logrus.Debugf("len: %d, #%s# #%s#\n", len(slicePath), slicePath[0], slicePath[1])
	logrus.Debugf("path: %s\n", path)

	switch path {
	case "/":
		handleRootConn(conn)
	case "/img/":
		handleImgConn(conn, imgPath)
	default:
		// 404 Page
		logrus.Errorf("path %s not exist", sliceLine[1])
		conn.Close()
	}

	return
}

func handleRootConn(conn net.Conn) {
	// 处理 / 页面请求
	// 打开文件夹
	imgDir, err := os.Open("./img/")
	if err != nil {
		logrus.Error(err)
		return
	}

	bodyHtml := `<center><h1 style="color:red">Hi, Golang Image Server</h1></center></br>`
	// 读取文件夹下的所有jpg webp图片文件
	imgs, err := imgDir.Readdirnames(-1)
	for _, img := range imgs {
		if strings.Contains(img, "jpg") || strings.Contains(img, "webp") {
			bodyHtml += fmt.Sprintf(`<img src="/img/%s" alt="%s"></br></br>`, img, img)
		}
	}

	// 向conn发送http的头部和body
	conn.Write([]byte(fmt.Sprintf(rootHead, getHttpNowDateStr(), len(bodyHtml))))
	conn.Write([]byte(bodyHtml))
	// 关闭conn
	conn.Close()
}

func handleImgConn(conn net.Conn, imgPath string) {
	// 转为本地相对路径
	// url编码 --> string
	imgLocalPath, err := url.QueryUnescape(imgPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	imgLocalPath = "."+imgLocalPath

	imgFile, err := os.Open(imgLocalPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer imgFile.Close()

	imgInfo, err := imgFile.Stat()
	if err != nil {
		logrus.Error(err)
		return
	}

	// image type
	sliceImgPath := strings.Split(imgPath, ".")
	imgType := sliceImgPath[len(sliceImgPath)-1]

	conn.Write([]byte(fmt.Sprintf(imgHead, getHttpNowDateStr(), imgType,imgInfo.Size())))
	io.Copy(conn, imgFile)

	conn.Close()
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	addr := ":8021"
	// 监听地址
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// 接受请求
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// 处理请求，可能阻塞，放入协程
		go handleRouteConn(conn)
		//go handleConn(conn)
	}
}
