package main

import (
	"os"
	"os/exec"
)

/*
实现 ps aux | grep bash | grep -v grep
程序默认是有三个接口，一个输入接口。两个输出接口。
管道的意义就是类似于水管，让上一个水管的输出接口接到下一个水管的输入接口
StdoutPipe() 默认是把标准输出和标准错误输出都打印
*/

func main() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "bash")
	cmd3 := exec.Command("grep", "-v", "grep")
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd3.Stdin, _ = cmd2.StdoutPipe()
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	cmd3.Start()
	cmd2.Start()
	cmd1.Run()
	cmd3.Wait()

}
