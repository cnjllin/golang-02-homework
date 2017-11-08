package main

import (
	"fmt"
	"net/http"
	"strings"
	"html"
	"github.com/51reboot/golang-02-homework/lesson4/alex-marmot/store"
)

type db struct {
	db_map map[string]interface{}
}

var DB_g = db{db_map:make(map[string]interface{})}

func init()  {
	DB_g.db_map = store.Load()
}

func processURL(url string) (key, val string) {
	fragments := strings.Split(html.EscapeString(url), "/")
	key = fragments[2]
	val = fragments[3]
	return key, val
}

func getKey(url string) (key string) {
	fragments := strings.Split(html.EscapeString(url), "/")
	key = fragments[2]
	return key
}

func setKV(w http.ResponseWriter, r *http.Request) {
	key, val := processURL(r.URL.Path)
	DB_g.db_map[key] = val
	store.Dump(DB_g.db_map)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln("Done!")))
}

func getKV(w http.ResponseWriter, r *http.Request) {
	key := getKey(r.URL.Path)
	val, exist := DB_g.db_map[key]
	if exist {
		v, _ := val.(string)
		w.Write([]byte(v))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("key %s not found", key)))
	}
}

func list(w http.ResponseWriter, r *http.Request)  {
	resp := store.Load()
	w.Write([]byte(fmt.Sprintln("List Page")))
	w.Write([]byte(fmt.Sprintf("Data: ", resp)))
}

func main() {
	http.HandleFunc("/set/", setKV)
	http.HandleFunc("/get/", getKV)
	http.HandleFunc("/", list)
	http.ListenAndServe(":3000", nil)
}

