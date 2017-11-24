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
	lineList := strings.Split(line, "|")

	type pips struct {
		r *io.PipeReader
		w *io.PipeWriter
	}
	mypips := make([]pips, len(lineList)-1)
	for i := 0; i < len(lineList)-1; i++ {
		mypips[i].r, mypips[i].w = io.Pipe()
	}
	// or mypips := make([]pips, len(lineList)-1)
	cmdlists := make([][]string, len(lineList))
	cmds := make([]*exec.Cmd, len(cmdlists))

	for i := 0; i < len(lineList); i++ {
		cmdlists[i] = strings.Fields(lineList[i])
		cmds[i] = exec.Command(cmdlists[i][0], cmdlists[i][1:]...)
		if i == 0 {
			cmds[i].Stdin = os.Stdin
		} else {
			cmds[i].Stdin = mypips[i-1].r
		}
		if i == len(lineList)-1 {
			cmds[i].Stdout = os.Stdout
		} else {
			cmds[i].Stdout = mypips[i].w
		}
		cmds[i].Stderr = os.Stderr
	}

	for i := 0; i < len(lineList); i++ {
		cmds[i].Start()
	}
	for i := 0; i < len(lineList); i++ {
		if i == 0 {
			//cmds[i].Start()
			cmds[i].Wait()
		} else {
			//cmds[i].Start()
			mypips[i-1].r.Close()
			mypips[i-1].w.Close()
			cmds[i].Wait()
		}
	}
>>>>>>> cc43a0065e7aca9e26ab8aaca1037ea98b6de632
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
