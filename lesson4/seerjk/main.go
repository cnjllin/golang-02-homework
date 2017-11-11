package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"github.com/51reboot/golang-02-homework/lesson4/seerjk/store"
)

type restfulDB struct {
	// 同步锁，防止同时写入
	lock sync.Mutex
	dbMap map[string]interface{}
}

var DBGlobal = restfulDB{dbMap: make(map[string]interface{})}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		// chrome 每次都会请求 /favicon.ico
		return
	}

	urlPathlList := strings.Split(r.URL.Path, "/")
	//fmt.Println(urlPathlList, len(urlPathlList))

	if urlPathlList[1] == "get" && len(urlPathlList) == 3 {
		// get key
		// 读取是否有必要lock？
		DBGlobal.lock.Lock()
		defer DBGlobal.lock.Unlock()
		key := urlPathlList[2]
		value, ok := DBGlobal.dbMap[key]
		if ok {
			// string类型断言，强制转换为string
			v, _ := value.(string)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%s", v)))
		} else {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(fmt.Sprintf("key %s not exist", key)))
		}
	} else if urlPathlList[1] == "set" && len(urlPathlList) == 4{
		// set key value
		DBGlobal.lock.Lock()
		defer DBGlobal.lock.Unlock()
		key := urlPathlList[2]
		value := urlPathlList[3]
		// write to memory
		DBGlobal.dbMap[key] = value
		// write to disk
		store.Dump(DBGlobal.dbMap)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("setted %s = %s\n", key, value)))

	} else {
		//参数错误
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintf("error request")))
	}
	fmt.Println(DBGlobal.dbMap)
}

func main() {
	DBGlobal.dbMap = store.Load()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
