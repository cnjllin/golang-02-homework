package main

import (
	"fmt"
	"strings"
)

func main() {
	pipeCommand("ps aux | grep ssh | grep -v grep")
}

func pipeCommand(line string) {
	// fmt.Println("pipe command.")
	lineList := strings.Split(line, "|")
	var cmdlines [][]string
	//var cmdline []string
	for i := 1; i < len(lineList)-1; i++ {
		cmdlines[i] = append(cmdlines[i], strings.Fields(lineList[i])...)
	}
	fmt.Println(cmdlines)
	/*
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
	*/
}
