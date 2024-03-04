package repository

//el import siempre es desde el directorio raiz, donde está el go.mod
import "bases/clases/05-structs/internal"

// manejarlo con una variable global complica el testing ya que cada test puede modificar la variable.
// Debería reiniciarse para cada test.
//var movies []internal.Movie

type MoviesSlice struct {
	//evita que la variable global sea una instancia que existe.
	movies []internal.Movie
}

func NewMovieSlice(movies []internal.Movie) MoviesSlice {
	return MoviesSlice{movies}
}

func (m MoviesSlice) Get() (mv []internal.Movie) {
	mv = make([]internal.Movie, len(m.movies))
	copy(mv, m.movies)

	return
}
