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

/*
	通过这样的方式，会在执行程序的过程中新增由程序本身出现的进程信息

*/
