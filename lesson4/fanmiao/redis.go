package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

//定义结构
type db struct {
	dbLock sync.Mutex
	dbMap  map[string]interface{}
}

var DB db

//读取数据
func getUrl(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	DB.dbMap = Load()
	if v, ok := DB.dbMap[url_list[2]]; ok {
		fmt.Fprintf(w, "读取成功，Key：%s,Value:%s", url_list[2], v)
	} else {
		fmt.Fprintf(w, "您读取的数据不存在，请您重新输入")
	}
}

//写入数据
func setUrl(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	DB.dbMap = make(map[string]interface{})
	if len(url_list) == 4 {
		DB.dbMap = Load()
		DB.dbLock.Lock()
		defer DB.dbLock.Unlock()
		DB.dbMap[url_list[2]] = url_list[3]
		Dump(DB.dbMap)
		fmt.Fprintf(w, "写入成功，Key：%s,Value:%s", url_list[2], url_list[3])
	}
}

//写入数据的主函数
func Dump(data map[string]interface{}) {
	buf, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Marshall error")
	}
	ioutil.WriteFile("dump.db", buf, 0644)
}

//读取数据的主函数
func Load() map[string]interface{} {
	f, err := ioutil.ReadFile("dump.db")
	if err != nil {
		log.Fatal("文件加载失败")
	}
	DB.dbMap = make(map[string]interface{})
	err1 := json.Unmarshal(f, &DB.dbMap)
	if err1 != nil {
		fmt.Println("Can't decode json message", err1)
	}
	return DB.dbMap
}

//初始化，判断文件是否存在
func init() {
	if _, err := os.Stat("dump.db"); err != nil {
		_, err1 := os.Create("dump.db")
		if err1 != nil {
			log.Fatal("文件创建失败")
		}
	}
}

func main() {
	http.HandleFunc("/get/", getUrl)
	http.HandleFunc("/set/", setUrl)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
