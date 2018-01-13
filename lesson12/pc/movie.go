package main

import (
	"net/http"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
}

func handlerx(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var movies []Movie
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
	}
	movies = append(movies, strangelove)

	http.HandleFunc("/data", handlerx)
	http.ListenAndServe(":8000", nil)
}
