package service

import "bases/clases/05-structs/internal"

const (
	ErrMsgNoMovies = "there are no movies"
)

func NewMovieDefault(rp internal.MovieRepository) MovieDefault {
	return MovieDefault{rp}
}

type MovieDefault struct {
	//rp is the movie repository
	rp internal.MovieRepository
}

func (m MovieDefault) AverageRating() (avg float64, err string) {
	movies := m.rp.Get()

	if len(movies) == 0 {
		err = ErrMsgNoMovies
		return
	}

	var sum float64
	for _, movie := range movies {
		sum += movie.Rating
	}

	avg = sum / float64(len(movies))

	return
}
