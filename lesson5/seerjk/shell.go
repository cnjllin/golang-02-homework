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
			onePipeCommand(line)
		} else {
			noPipeCommand(line)
		}

	}
}
