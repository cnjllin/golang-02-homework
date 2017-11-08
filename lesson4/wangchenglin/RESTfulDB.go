package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

/*
实现一个RESTful的KV存储，能持久化保存数据，并且在启动的时候载入之前保存的数据

Set接口

http://192.168.1.8:4000/set/Key/Value
写入Key:Value

Get接口

http://192.168.1.8:4000/get/xxxxxx
获取xxxxx对应的Value，如果不存在就返回HTTP Code 404

持久化

每次Set操作都进行存盘操作
*/

type db struct {
	dbLock sync.Mutex
	dbMap  map[string]interface{}
}

var g_DB = db{dbMap: make(map[string]interface{})}

const fileName string = "dump.db"

// 创建一个RESTfulDBHandler类型以实现ServeHTTP接口方法
type RESTfulDBHandler struct{}

type Json struct {
	data interface{}
}

func init() {
	// 判断文件是否存在, 若不存在则创建文件
	fs, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		log.Printf("不存在数据文件，即将创建%s文件...\n", fileName)
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("创建%s文件失败, 程序退出!!!\n", fileName)
		}
		f.Close()
		log.Printf("创建%s文件成功.\n", fileName)
	}

	if fs.Size() == 0 {
		log.Println("空文件")
	}

}

func main() {
	var dbh RESTfulDBHandler

	// 读取文件
	g_DB.dbMap = Load()

	// 启动HTTP服务监听
	err := http.ListenAndServe(":8080", dbh)
	if err != nil {
		log.Fatal(err)
	}
}

// 写入数据库文件
func Dump(data map[string]interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Printf("Marshal json error: %s", err)
		return
	}
	ioutil.WriteFile(fileName, j, 0644)
}

// 读取数据库文件
func Load() map[string]interface{} {

	j, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Load db error: %s", err)
		return nil
	}
	log.Printf("%s数据文件加载成功~\n", fileName)
	if len(j) == 0 {
		return make(map[string]interface{})
	}

	json_j, err := NewJson(j)
	if err != nil {
		log.Printf("Unmarshal error: %s", err)
		return nil
	}

	m := json_j.Map()
	return m
}

// 实现支持RESTful协议的ServeHTTP接口方法
func (dbh RESTfulDBHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(html.EscapeString(r.URL.Path), "/")
	if url_list[1] == "get" && len(url_list) == 3 {
		g_DB.dbLock.Lock()
		defer g_DB.dbLock.Unlock()
		v, exist := g_DB.dbMap[url_list[2]]
		if exist {
			v_str, _ := v.(string)
			w.Write([]byte(v_str))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("key %s not found", url_list[2])))
		}
	} else if url_list[1] == "set" && len(url_list) == 4 {
		g_DB.dbLock.Lock()
		defer g_DB.dbLock.Unlock()
		g_DB.dbMap[url_list[2]] = url_list[3]
		Dump(g_DB.dbMap)
		w.Write([]byte(fmt.Sprintf("key:%s, value:%s: SET成功", url_list[2], url_list[3])))
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprint("error request length")))
	}
	return
}

func NewJson(body []byte) (*Json, error) {
	j := new(Json)
	err := j.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// 实现json.Marshal接口方法
func (j *Json) MarshalJSON() ([]byte, error) {
	return json.Marshal(&j.data)
}

// 实现json.Unmarshal接口方法
func (j *Json) UnmarshalJSON(body []byte) error {
	return json.Unmarshal(body, &j.data)
}

// Map方法将Json类型的data转换为一个map类型
func (j *Json) Map() map[string]interface{} {
	m, ok := (j.data).(map[string]interface{})
	if ok {
		return m
	}

	log.Println("type assertion to map[string]interface{} failed")
	return nil
}
