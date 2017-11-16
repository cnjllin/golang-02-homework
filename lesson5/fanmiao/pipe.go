package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[fanmiao@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	//r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		if !r.Scan() {
			break
		}
		line := r.Text()
		// line, _ := r.ReadString('\n')
		// line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cmds := strings.Split(line, "|")
		s1 := strings.Fields(cmds[0])
		s2 := strings.Fields(cmds[1])
		s3 := strings.Fields(cmds[2])

		c1 := exec.Command(s1[0], s1[1:]...)
		c2 := exec.Command(s2[0], s2[1:]...)
		c3 := exec.Command(s3[0], s3[1:]...)

		c2.Stdin, _ = c1.StdoutPipe()
		c3.Stdin, _ = c2.StdoutPipe()

		c3.Stdout = os.Stdout
		c3.Start()
		c2.Start()
		c1.Run()
		c3.Wait()

	}
}
