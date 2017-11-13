package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[golang@%s]$ ", host)
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
		//log.Println("line -->", line)
		buffer, err := exec.Command("bash", "-c", line).Output()
		if err != nil {
			log.Println(err)
		}
		fmt.Print(string(buffer))
	}
}
