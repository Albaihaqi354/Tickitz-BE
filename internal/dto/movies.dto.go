package dto

import "time"

type GetUpcomingMovie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	PosterUrl   string    `json:"poster_url"`
	ReleaseDate time.Time `json:"release_date"`
	GenresName  string    `json:"genres"`
}

type GetPopularMovie struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	PosterUrl  string `json:"poster_url"`
	GenresName string `json:"genres"`
}

type GetMovieWitFilter struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	PosterUrl  string `json:"poster_url"`
	GenresName string `json:"genres"`
}

type GetMovieDetail struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Synopsis        string    `json:"synopsis"`
	Duration        int       `json:"duration"`
	ReleaseDate     time.Time `json:"release_date"`
	Director        string    `json:"director"`
	Cast            string    `json:"cast"`
	PosterUrl       string    `json:"poster_url"`
	BackDropUrl     string    `json:"backdrop_url"`
	PopularityScore int       `json:"popularity_score"`
	GenresName      string    `json:"genres"`
}
