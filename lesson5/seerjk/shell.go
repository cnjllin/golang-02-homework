package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"os/exec"
	"io"
	"log"
)

func onePipeCommand(cmdLine string) {
	// 切分管道两端 支持单个管道
	cmds := strings.Split(cmdLine, "|")
	s1 := strings.Fields(cmds[0])
	s2 := strings.Fields(cmds[1])

	cmd1 := exec.Command(s1[0], s1[1:]...)
	cmd2 := exec.Command(s2[0], s2[1:]...)

	r, w := io.Pipe()
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = w
	cmd2.Stdin = r
	cmd2.Stdout = os.Stdout

	cmd1.Start()
	cmd2.Start()

	cmd1.Wait()
	w.Close()
	cmd2.Wait()
	r.Close()
}

func multiplePipesCommand2(cmdLine string) {
	// 有问题 单个Pipe() 无法实现多重管道机制？
	// [理解golang io.Pipe](http://www.jianshu.com/p/aa207155ca7d)
	pipesNum := strings.Count(cmdLine, "|")
	cmds := strings.Split(cmdLine, "|")

	s1 := strings.Fields(cmds[0])
	cmd1 := exec.Command(s1[0], s1[1:]...)
	r, w := io.Pipe()
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = w
	cmd1.Start()
	cmd1.Wait()

	for i := 1; i <= pipesNum; i++ {
		s2 := strings.Fields(cmds[i])
		cmd2 := exec.Command(s2[0], s2[1:]...)
		if i == pipesNum {
			cmd2.Stdin = r
			cmd2.Stdout = os.Stdout
			cmd2.Start()
			cmd2.Wait()
			w.Close()
			r.Close()
		} else {
			cmd2.Stdin = r
			cmd2.Stdout = w
			cmd2.Start()
			cmd2.Wait()
		}
	}
}

func multiplePipesCommand(cmdLine string) {
	// 多重管道
	// 参考 https://stackoverflow.com/questions/41361929/go-pipe-3-or-more-commands-with-os-exec
	cmds := strings.Split(cmdLine, "|")
	cmdsList := make([]*exec.Cmd, len(cmds))
	for i := 0; i < len(cmds); i++ {
		sCmd := strings.Fields(cmds[i])
		cmdsList[i] = exec.Command(sCmd[0], sCmd[1:]...)
		if i == 0 {
			// 什么都不做
			//cmdsList[i].Stdin = os.Stdin
			//cmdsList[i+1].Stdin, _ = cmdsList[i].StdoutPipe()
		} else if i == len(cmds) - 1 {
			// 最后一次 cmd 输出到 os 屏幕 Stdout
			cmdsList[i].Stdin, _ = cmdsList[i-1].StdoutPipe()
			cmdsList[i].Stdout = os.Stdout
		} else {
			// 中间：当前cmd Stdin 来自 上一个cmd的 StdoutPipe
			cmdsList[i].Stdin, _ = cmdsList[i-1].StdoutPipe()
		}
	}
	// Start 和 Wait 需要分开
	for i := 0; i < len(cmds); i++ {
		cmdsList[i].Start()
	}
	for i := 0; i < len(cmds); i++ {
		cmdsList[i].Wait()
	}
}

func noPipeCommand(cmdLine string) {
	cmdLineList := strings.Fields(cmdLine)
	cmd := exec.Command(cmdLineList[0], cmdLineList[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[seerjk@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		if !r.Scan() {
			break
		}
		line := r.Text()
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, "|") {
			multiplePipesCommand(line)
		} else {
			noPipeCommand(line)
		}

	}
}
