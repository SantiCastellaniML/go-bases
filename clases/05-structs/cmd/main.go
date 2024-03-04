package main

import (
	"bases/clases/05-structs/internal"
	"bases/clases/05-structs/internal/repository"
	"bases/clases/05-structs/internal/service"
	"fmt"
)

func main() {
	// dependencies
	var rp internal.MovieRepository
	db := []internal.Movie{
		{ID: 1, MovieAttributes: internal.MovieAttributes{Title: "Avengers", Rating: 8.0}},
		{ID: 2, MovieAttributes: internal.MovieAttributes{Title: "Batman", Rating: 7.0}},
	}

	rp = repository.NewMovieSlice(db)
	sv := service.NewMovieDefault(rp)

	avg, err := sv.AverageRating()

	if err != "" {
		fmt.Println(err)
		return
	}

	println(avg)
}
