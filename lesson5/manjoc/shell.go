package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// 支持重定向
func redirectCommand(line string) {
	// fmt.Println("redirect command.")
	lineList := strings.Split(line, ">")
	cmdLine := lineList[0]
	outFileName := strings.Fields(lineList[1])[0]
	outPutFile, _ := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	args := strings.Fields(cmdLine)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = outPutFile
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// 支持管道的命令
func pipeCommand(line string) {
	// fmt.Println("pipe command.")
	lineList := strings.Split(line, "|")
	//cmdLine := lineList[0]
	// pipeLeft, pipeRight
	pipeLeft := strings.Fields(lineList[0])
	pipeRight := strings.Fields(lineList[1])
	r, w := io.Pipe()
	cmdLeft := exec.Command(pipeLeft[0], pipeLeft[1:]...)
	cmdRight := exec.Command(pipeRight[0], pipeRight[1:]...)
	cmdLeft.Stdin = os.Stdin
	cmdLeft.Stdout = w
	cmdLeft.Stderr = os.Stderr

	cmdRight.Stdin = r
	cmdRight.Stdout = os.Stdout

	cmdLeft.Start()
	cmdRight.Start()

	cmdLeft.Wait()
	r.Close()
	w.Close()
	cmdRight.Wait()
}

// 一般命令，无 管道， 无 重定向
func generalCommand(line string) {
	// fmt.Println("general command.")
	args := strings.Fields(line)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("manjo@%s]$", host)
	stdRead := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s", prompt)
		if !stdRead.Scan() {
			break
		}
		line := stdRead.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, ">") {
			// >>
			redirectCommand(line)
		} else if strings.Contains(line, "|") {
			// pipe
			pipeCommand(line)
		} else {
			// general
			generalCommand(line)
		}
	}
}
