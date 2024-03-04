package internal

type Movie struct {
	ID    int
	Title string
	MovieAttributes
}

type MovieAttributes struct {
	Title  string
	Rating float64
}
