package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

var help bool

// 定义需要等待的组
var wg sync.WaitGroup

// 命令用法说明
func usage() {
	fmt.Println(`Usage: gocat <filesource> or gocat <https://www.baidu.com> `)
	//flag.PrintDefaults()
}

func init() {
	// 注册usage()函数
	flag.Usage = usage
	flag.BoolVar(&help, "help", false, "gocat <filesource> or gocat <https://www.baidu.com>")
}

func main() {

	flag.Parse()

	// 如果输入--help参数则提示命令用法
	if help {
		flag.Usage()
		return
	}

	// 参数个数为0则提示命令的用法
	argsNum := flag.NArg()
	if argsNum == 0 {
		flag.Usage()
		return
	}

	for i := 0; i < argsNum; i++ {
		// 组中增加等待
		wg.Add(1)
		argv := flag.Arg(i)
		go parseArgs(argv)
	}

	// 等待组协程执行完成
	wg.Wait()
}

func parseArgs(argv string) {
	// 判断命令是文件还是网址
	if strings.HasPrefix(argv, "http:") || strings.HasPrefix(argv, "https:") {
		parseUrl(argv)
	} else {
		printFile(argv)
	}
	// 执行完毕
	wg.Done()
}

// 解析网站内容
func parseUrl(url string) {
	resp, _ := http.Get(url)
	fmt.Println(resp)
}

// 读取文件并显示
func printFile(filename string) {
	f, err := os.Open(filename)

	defer f.Close()

	if err != nil {
		fmt.Printf("打开文件错误! 文件名[%s]\n", filename)
		return
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(string(content))
}
