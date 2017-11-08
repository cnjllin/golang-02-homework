package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//支持set、get、list、delete操作

var redismap map[string]string

func main() {
	redismap = make(map[string]string)
	buf, err := ioutil.ReadFile("redis.data")
	if err != nil {
		f, _ := os.Create("redis.data")
		f.Close()
	}
	json.Unmarshal(buf, &redismap)
	defer FormatJson(redismap)
	http.HandleFunc("/", HandleRedis)
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func FormatJson(redismap map[string]string) {
	js, err := json.Marshal(redismap)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("redis.data", js, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleRedis(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formstag := strings.Split(r.URL.Path, "/")
	switch formstag[1] {
	case "set":
		if len(formstag) != 4 {
			fmt.Fprintf(w, "http输入格式错误,支持set、get、list、delete操作，set实例：http://xxxx/set/{key}/{value}")
		}
		redismap[formstag[2]] = formstag[3]
		FormatJson(redismap)
	case "get":
		if len(formstag) != 3 {
			fmt.Fprintf(w, "http输入格式错误,支持set、get、list、delete操作，get实例：http://xxxx/get/{key}")
		}
		value, ok := redismap[formstag[2]]
		if !ok {
			w.WriteHeader(404)
		}
		fmt.Fprintf(w, value)
	case "list":
		fmt.Fprintln(w, redismap)
	case "delete":
		if len(formstag) > 3 {
			fmt.Fprintf(w, "http输入格式错误,支持set、get、list、delete操作，delete实例：http://xxxx/delete/{key}")
		}
		delete(redismap, formstag[2])
		FormatJson(redismap)
	default:
		fmt.Fprintf(w, "http输入格式错误,支持set、get、list、delete操作;set格式：http://xxxx/set/{key}/{value};get格式：http://xxxx/get/{key};delete格式：http://xxxx/delete/{key}")
	}
}
