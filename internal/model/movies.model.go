package model

import "time"

type Movie struct {
	Id              int       `db:"id"`
	Title           string    `db:"title"`
	Synopsis        string    `db:"synopsis"`
	Duration        int       `db:"duration"`
	ReleaseDate     time.Time `db:"release_date"`
	DirectorId      int       `db:"director_id"`
	Director        string    `db:"director"`
	PosterUrl       string    `db:"poster_url"`
	BackdropUrl     string    `db:"backdrop_url"`
	PopularityScore float64   `db:"popularity_score"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
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
	MovieId int `db:"movie_id"`
	GenreId int `db:"genre_id"`
}

type MovieCast struct {
	MovieId int `db:"movie_id"`
	ActorId int `db:"actor_id"`
}

type MovieDetail struct {
	Id              int       `db:"id"`
	Title           string    `db:"title"`
	Synopsis        string    `db:"synopsis"`
	Duration        int       `db:"duration"`
	ReleaseDate     time.Time `db:"release_date"`
	Director        string    `db:"director"`
	Cast            string    `db:"cast"`
	PosterUrl       string    `db:"poster_url"`
	BackdropUrl     string    `db:"backdrop_url"`
	PopularityScore float64   `db:"popularity_score"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	GenresName      string    `db:"genre_name"`
	ScheduleCount   int       `db:"schedule_count"`
}
