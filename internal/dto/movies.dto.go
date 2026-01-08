package dto

import "time"

type GetUpcomingMovie struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Synopsis        string    `json:"synopsis"`
	Duration        int       `json:"duration"`
	ReleaseDate     time.Time `json:"release_date"`
	Director        Director  `json:"director"`
	Cast            string    `json:"cast"`
	PosterUrl       string    `json:"poster_url"`
	BackDropUrl     string    `json:"backdrop_url"`
	PopularityScore float64   `json:"popularity_score"`
	GenresName      string    `json:"genre_name"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Director struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
