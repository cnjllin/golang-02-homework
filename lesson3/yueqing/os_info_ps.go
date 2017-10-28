package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := "ps -auxw | awk '{print $2, $NF}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
