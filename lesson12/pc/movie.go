package main

import (
	"net/http"
)

type Movie struct {
	Name, Positon, Office, Ext, StartDate, Salary string
}

/*
{
  "data": [
    [
      "Tiger Nixon",
      "System Architect",
      "Edinburgh",
      "5421",
      "2011/04/25",
      "$320,800"
    ],
    [
      "Garrett Winters",
      "Accountant",
      "Tokyo",
      "8422",
      "2011/07/25",
      "$170,750"
    ]
}
*/
func handlerx(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var movies []Movie
	strangelove := Movie{
		Name:      "Dr. Strangelove",
		Positon:   "xxx",
		Office:    "dfa",
		Ext:       "dfsa",
		StartDate: "2018",
		Salary:    "100000",
	}
	movies = append(movies, strangelove)

	http.HandleFunc("/data", handlerx)
	http.ListenAndServe(":8000", nil)
}
