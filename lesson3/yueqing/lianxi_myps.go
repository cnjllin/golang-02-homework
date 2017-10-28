package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"strconv"
)

// 打印相应pid 对应的cmdline 信息
func printCmdline(s string) string {
	// 字符串拼接
	str := "/proc/" + s + "/cmdline"
	body, err := ioutil.ReadFile(str)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

//判断 是否是pid 进程对应的文件
func isPidBool(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
func main() {
	// 获取 /proc 目录路径下的所有目录的信息
	infos, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		// 判断是否是路径目录
		if info.IsDir() {
			// 判断是否是pid 进程的文件目录
			if isPidBool(info.Name()) {
				pidinfo := printCmdline(info.Name())
				fmt.Println(info.Name(), pidinfo)
			}
		}
	}
}
