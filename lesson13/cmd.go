package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	out, _ := cmd.Output()
	fmt.Printf("%s\n", out)
}
