package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
实现 ps aux | grep bash | grep -v grep
程序默认是有三个接口，一个输入接口。两个输出接口。
管道的意义就是类似于水管，让上一个水管的输出接口接到下一个水管的输入接口
StdoutPipe() 默认是把标准输出和标准错误输出都打印
*/

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[Golang@%s] ^ ", host)
	r := bufio.NewScanner(os.Stdin) // 默认是按行读取
	for {
		fmt.Print(prompt)
		if !r.Scan() { // 判断是否有输入
			break
		}
		line := r.Text() // 获取内容
		if len(line) == 0 {
			continue
		}

		temp := strings.Split(line, "|")
		length := len(temp)
		switch length {
		case 1:
			args := strings.Fields(temp[0]) // 默认是以空格区分
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Print(err)
			}
		case 2:
			arg1 := strings.Fields(temp[0])
			arg2 := strings.Fields(temp[1])
			cmd1 := exec.Command(arg1[0], arg1[1:]...)
			cmd2 := exec.Command(arg2[0], arg2[1:]...)
			cmd2.Stdin, _ = cmd1.StdoutPipe()
			cmd2.Stdout = os.Stdout
			cmd2.Stderr = os.Stderr
			cmd2.Start()
			cmd1.Run()
			cmd2.Wait()
		case 3:
			arg1 := strings.Fields(temp[0])
			arg2 := strings.Fields(temp[1])
			arg3 := strings.Fields(temp[2])
			cmd1 := exec.Command(arg1[0], arg1[1:]...)
			cmd2 := exec.Command(arg2[0], arg2[1:]...)
			cmd3 := exec.Command(arg3[0], arg3[1:]...)

			cmd2.Stdin, _ = cmd1.StdoutPipe()
			cmd3.Stdin, _ = cmd2.StdoutPipe()
			cmd3.Start()
			cmd2.Start()
			cmd1.Run()
			cmd3.Wait()

		default:
			log.Println("真的是写不动了，这个方式太苦力了")
		}
	}
}
