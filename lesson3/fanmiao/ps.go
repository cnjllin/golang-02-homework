package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

//判断是否有错误
func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//打印pid和正在执行的命令
func printcmd(name string) {
	path := "/proc/" + name + "/cmdline"
	buf, err := ioutil.ReadFile(path)
	checkerr(err)
	//判断cmdline是否为空，为空则不打印
	if string(buf) != "" {
		fmt.Printf("%s\t%s\n", name, string(buf))
	}
}

func main() {
	piddirs, err := ioutil.ReadDir("/proc")
	checkerr(err)
	for _, piddir := range piddirs {
		if m, _ := regexp.MatchString("^[0-9]{1,5}$", piddir.Name()); m {
			printcmd(piddir.Name())
		}
	}

}
