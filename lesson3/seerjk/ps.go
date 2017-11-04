package main

//要求:
//编写一个mini ps，输出两列: pid和命令行
//类似 1234 /bin/bash
//
//思路:
//读取 /proc 目录下的文件
//获取所有pid子目录
//读取pid子目录下的cmdline文件

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	infos, err := ioutil.ReadDir("/proc/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%6s %-20s\n", "PID", "CMD")

	for _, info := range infos {
		if info.IsDir() {
			pid, err := strconv.Atoi(info.Name())

			if err == nil {
				//string(pid)无法将int转为string
				//strconv.Atoi(pid_int) --> pid_list []byte --> string(pid_list)
				// 读取cmdline文件 中的命令
				body, err := ioutil.ReadFile("/proc/" + string(info.Name()) + "/cmdline")
				if err != nil {
					log.Fatal(err)
				}

				if len(body) > 0 {
					// 用户进程 cmdline中可以读取运行命令
					fmt.Printf("%6d %-20s\n", pid, string(body))
				} else {
					// 内核中的命令 body(cmdline)为空， 读取 status文件
					statusBody, err := ioutil.ReadFile("/proc/" + string(info.Name()) + "/status")
					if err != nil {
						log.Fatal(err)
					}
					lines := strings.Split(string(statusBody), "\n")
					kernelCmd := strings.Split(lines[0], ":")[1]
					kernelCmd = "[" + strings.TrimSpace(kernelCmd) + "]"

					fmt.Printf("%6d %-20s\n", pid, kernelCmd)
				}

			}
		}
	}
}
