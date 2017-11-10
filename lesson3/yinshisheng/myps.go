package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func title() {
	fmt.Printf("%s\t%s\n", "PID", "cmdline")
	fmt.Printf("%s\n", "--------------------------")
}

func getpid(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		pid, _ := strconv.Atoi(file.Name())

		if pid != 0 {
			getcmdline(pid)
		}
	}
}

func getcmdline(pid int) {
	p := strconv.Itoa(pid)
	cmdline := "/proc/" + p + "/cmdline"
	cont, err := ioutil.ReadFile(cmdline)

	if err != nil {
		log.Fatal(err)
	}

	if len(cont) == 0 {
		fmt.Printf("%d\t%s\n", pid, "--")
	} else {
		fmt.Printf("%d\t%s\n", pid, cont)
	}
}

func main() {
	path := "/proc"
	title()
	getpid(path)
}
