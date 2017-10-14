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
		// 打印错误日志并退出
		log.Fatal(err)
	}
	for {
		// 无限循环
		conn, err := l.Accept()
		// golang异常通过返回值处理 err
		if err != nil {
			// 打印错误日志并退出
			log.Fatal(err)
		}
		// 协程 运行 handle函数，分身术，类似线程方式的逻辑实现有限状态机
		go handle(conn)
	}
}

// nc 127.0.0.1 8080
// 进程：重（重量） 早（历史出现次序）
// 线程：中（重量） 晚（历史出现次序）
// 协程：轻（重量） 中（历史出现次序）
// 进程、线程、协程 （由重到轻）
// 进程、协程、线程 （历史出现次序）
// kill -9 xxx 释放cpu和mem
// 进程：OS管理资源的最小单位，2个进程的CPU和MEM是独立的
// 线程：进程内部，每个线程可以占用1个CPU core，
//      产生原因 － 为了充分利用CPU的多核心
//      进程不共享MEM，同一个进程中的线程共享MEM
//      Linux上 线程ID空间占用进程PID空间 同样的PID空间

//[root@devops opsweb]# pstree -ap 1813
//mysqld,1813 --basedir=/usr --datadir=/var/lib/mysql --plugin-dir=/usr/lib64/mysql/plugin--log-err
//├─{mysqld},1894
//├─{mysqld},1895
//├─{mysqld},1896
//├─{mysqld},1897
//├─{mysqld},1898
//...
//├─{mysqld},2816
//└─{mysqld},2817

// 线程共享：内存空间、PID空间、文件描述符 FD
//    进程或线程 都是 clone()

// squid  进程
// varnish 多线程，非常多线程抢占，切换CPU
// nginx 默认单线程，多进程，性能无冕之王
// redis 单进程，单线程
//       I/O多路复用， 一个人对应多个妹子同时聊天：收一下，读完信息，编辑发送；切换下一个妹子的信息，读，编辑与发送。
//       设计困难：防止妹子信息搞混，记住每一个妹子的信息。有限状态机编程

// golang，用开线程方式写逻辑，然后 go 调用，实现了 有限状态机，分身术
// 参考： https://zhuanlan.zhihu.com/auxten?topic=%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B
//       网络编程（九）：Python裸写异步非阻塞网络框架  https://zhuanlan.zhihu.com/p/23614423

// CSP模式  golang CSP
//    golang; CSP式的并发模型 https://studygolang.com/articles/763