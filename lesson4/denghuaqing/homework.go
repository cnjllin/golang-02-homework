package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"net/http"
	"html"
	"strings"
	"log"
)

func getdb(keys string)  []byte{
	db,_ := leveldb.OpenFile("./leveldb",nil)
	defer db.Close()
	date,_ := db.Get([]byte(keys),nil)
	return date
}

func putdb(keys string ,vaule string) {
	db,_ := leveldb.OpenFile("./leveldb",nil)
	defer db.Close()
	db.Put([]byte(keys),[]byte(vaule),nil)
}

type Htt struct {}

func (htt Htt)  ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	url_list := strings.Split(html.EscapeString(r.URL.Path),"/")
	if url_list[1] == "get" && len(url_list) == 3 {
		db_get := getdb(url_list[2])
		w.Write(db_get)
	}else if url_list[1] == "set" && len(url_list) ==4 {
		db_file := getdb(url_list[2])
		if db_file != nil {
			w.Write([]byte(fmt.Sprintf("keys %s is exists",url_list[2])))
		} else {
		putdb(url_list[2], url_list[3])
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("set %s to %s", url_list[3], url_list[2])))
		}
		} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprint("error url")))
	}
	return
}

func main()  {
	var htt Htt
	err := http.ListenAndServe("0.0.0.0:8000",htt)
	if err != nil{
		log.Fatal(err)
	}
}

