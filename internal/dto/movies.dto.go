package dto

import "time"

type GetUpcomingMovie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	PosterUrl   string    `json:"poster_url"`
	ReleaseDate time.Time `json:"release_date"`
	GenresName  string    `json:"genres"`
}
