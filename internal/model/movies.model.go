package model

import "time"

type Movie struct {
	Id              int       `db:"id"`
	Title           string    `db:"title"`
	Synopsis        string    `db:"synopsis"`
	Duration        int       `db:"duration"`
	ReleaseDate     time.Time `db:"release_date"`
	Director        Director  `db:"director"`
	PosterUrl       string    `db:"poster_url"`
	BackDropUrl     string    `db:"backdrop_url"`
	PopularityScore float64   `db:"popularity_score"`
}

type Director struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type Actor struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type Genre struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type MovieGenre struct {
	MovieID int `db:"movie_id"`
	GenreID int `db:"genre_id"`
}

type MovieCast struct {
	MovieID int `db:"movie_id"`
	ActorID int `db:"actor_id"`
}

type MovieDetail struct {
	Movie
	Cast       string `db:"cast"`
	GenresName string `db:"genre_name"`
}
