package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"strings"
	//"io/ioutil"
	"os"
	"io"
)

// ftp server
// client -> GET /home/jiangk/.bashrc\n
// server -> 返回文件内容，之后关闭
/*
扩展：
1. 文件大小对服务器对影响
	ioutil.ReadFile(path)  // 一次性读取到内存
	按块读取，一边读取文件，一边向socket中写
2. 如何设置ftp服务器的虚拟根目录
3. ftp客户端 **作业**
	// ioutil.ReadAll(conn) 一次性，大文件不适合
	io.Copy(f, conn)

4. 其它操作
	PUT 上传文件
	ls
*/

func HandleConn(conn net.Conn) {
	// 5. 断开连接
	defer conn.Close()
	log.Print("remote: ", conn.RemoteAddr())
	log.Print("locael: ", conn.LocalAddr())
	// telnet 127.0.0.1 8080
	// 4. 发送接收数据
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("#%s#", line)
	fs := strings.Fields(line)
	method := fs[0]
	path := fs[1]
	log.Printf("method: %s", method)
	log.Printf("path: #%s#", path)

	if strings.ToUpper(method) == "GET" {
		/*
		// 方法1. ioutil.ReadFile 一次性全部读入内存

		// 只适合小文件
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
			//fmt.Fprintf(conn, err)
		}
		result := string(content)
		fmt.Fprintf(conn, result)
		*/
		/*
		// 方法2.按块block读取，应对大文件
		buf := make([]byte, 4096)
		f, err := os.Open(path)
		if err != nil {
			log.Print(err)
		}
		defer f.Close()
		for {
			// read 会有游标 记住当前位置
			n, err := f.Read(buf)
			if err != nil {
				break
			}
			conn.Write(buf[:n])
		}
		*/

		/*
		// 当前位置 向后偏移10
		f.Seek(10, os.SEEK_CUR)
		// 回到开头
		f.Seek(0, os.SEEK_SET)
		*/
		// 方法3. bufio 内部开buf

		// 方法4. io.Copy 本质上根按块一样
		f, err := os.Open(path)
		if err != nil {
			log.Print(err)
			return
		}
		defer f.Close()
		// io.Copy 32K buffer
		io.Copy(conn, f)
		// io.CopyBuffer 自建buf
		//buf := make([]byte, 4096) // 从内存池里面拿出来一个buffer 复用buffer空间
		//io.CopyBuffer(conn, f, buf)
	} else {
		result := "Operation Error!!"
		fmt.Fprintf(conn, result)
	}


	//fmt.Fprintf(conn, result)
}

func main() {
	// 1.监听端口 0.0.0.0:8080 -> :8080
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		// 协议错误；port被占用
		log.Fatal(err)
	}

	for {
		// 拆分出协程的关键点：找到阻塞的原因，把阻塞部分拆到一个新协程中
		// 阻塞是因为 Accept连接 和 处理连接 在同一个协程中
		// 阻塞发生在 处理连接 过程，每次Accept后单独放一个新的协程中
		// Accept过程 不会发送阻塞
		// 2. 接收新连接
		conn, err := listener.Accept()  // net.Dial
		if err != nil {
			log.Fatal(err)
		}

		// 3. 启动协程
		go HandleConn(conn)
	}
}
