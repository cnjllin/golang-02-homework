package main

import (
	"net/http"
	"encoding/json"
	log "github.com/auxten/logrus"
)

type MovieDataTable struct {
	// struct 中的变量要大写才能格式化为json
	Data []Movie 	`json:"data"`
}

type Movie struct {
	Title string 	`json:"title"`
	Year int 		`json:"released"`
	Color bool 		`json:"color"`
	Actors string 	`json:"actors"`
}

func handleMovie(w http.ResponseWriter, r *http.Request) {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: "Humphrey Bogart"},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: "Paul Newman"},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: "Steve McQueen"},
	}

	dt := MovieDataTable{
		Data: movies,
	}
	//fmt.Println(dt)
	dtData, err := json.Marshal(dt)
	//fmt.Println(string(dtData))
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	if r.Method == "POST" {
		w.Write(dtData)
	}
}


func main() {
	// movie 数据接口
	http.HandleFunc("/movie", handleMovie)

	// go 静态站点
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	http.ListenAndServe(":8000", nil)
}