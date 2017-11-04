package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

/*
作业:

要求:

编写一个mini ps，输出两列: pid和命令行

类似 1234 /bin/bash

思路:

1.读取/proc目录下的文件
2.获取所有pid子目录
3.读取pid子目录下的cmdline文件
*/

var titleFlag bool

func matchFileName(fileName string) bool {
	match, _ := regexp.MatchString("^[0-9]*$", fileName)
	return match
}

func printTitle() {
	if !titleFlag {
		fmt.Println("PID\t CMD")
		titleFlag = true
	}
}

func main() {
	// 读取/proc目录下的文件
	procFiles, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal("读取/proc文件夹相关内容失败")
	}

	for _, procFile := range procFiles {
		// 匹配名称为数字的PID文件夹名称
		isTrue := matchFileName(procFile.Name())
		if isTrue {
			if procFile.IsDir() {
				// 组装cmdfile文件路径
				cmdFile := "/proc/" + procFile.Name() + "/cmdline"

				file, err := ioutil.ReadFile(cmdFile)
				if err != nil {
					fmt.Printf("读取%s文件失败!\n", cmdFile)
					continue
				}
				// 打印命令首行Title
				printTitle()
				fmt.Printf("%s\t %s\n", procFile.Name(), file)
			}
		}
	}

}
