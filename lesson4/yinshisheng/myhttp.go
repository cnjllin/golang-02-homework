package main

import (
	"fmt"
	"log"
	"encoding/json"
	"os"
	"bufio"
	"io"
	"net/http"
	"strings"
)

type UrlData struct {
	Key   string
	Value string
}

//初始化data.json
func InitJson() {
	f, err := os.OpenFile("data.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 100)
	count,err := f.Read(data)
	if err != nil {
		fmt.Println("文件内容长度为:", count, "已初始化文件")
		f.WriteString("{\"Key\":\"\",\"Value\":\"\"}")
	}
}

//set函数
func set(w http.ResponseWriter, request *http.Request) {
	k := strings.Split(request.URL.Path,"/")[2]
	v := strings.Split(request.URL.Path,"/")[3]
	SaveJson(k,v)
	fmt.Fprintf(w,"<p>添加成功 key:%s value:%s</p>", k, v)
}

//get函数
func get(w http.ResponseWriter, request *http.Request) {
	k := strings.Split(request.URL.Path,"/")[2]
	s1 := ReadJson("data.txt")
	var u UrlData
	var n int
	for _, v1 := range s1 {
		if v1 == "" {
			continue
		}
		err := json.Unmarshal([]byte(v1), &u)
		if err != nil {
			log.Fatal(err)
		}
		if k == u.Key {
			n = 1
			fmt.Fprintf(w,"<p>key:%s value:%s</p>", u.Key, u.Value)
		}
	}
	if n == 0 {
		fmt.Fprintf(w,"<p>404</p>")
	}
}

//读取文件数据
func ReadJson(json_file string) []string {
	var s []string
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		s = append(s, line)
		if err != nil || io.EOF == err {
			break
		}
	}
	return s
}

//数据存盘
func SaveJson(k, v string) {
	f, err := os.OpenFile("data.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := "{\"Key\":\"" + k + "\",\"Value\":\"" + v + "\"}\n"
	f.WriteString(data)
}


func main() {
	http.HandleFunc("/set/", set)
	http.HandleFunc("/get/", get)
	http.ListenAndServe(":8081",nil)
}
