package dto

import (
	"mime/multipart"
	"time"
)

type GetAllMovieAdmin struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Synopsis        string    `json:"synopsis"`
	Duration        int       `json:"duration"`
	ReleaseDate     time.Time `json:"release_date"`
	Director        string    `json:"director"`
	Cast            string    `json:"cast"`
	PosterUrl       string    `json:"poster_url"`
	BackdropUrl     string    `json:"backdrop_url"`
	PopularityScore float64   `json:"popularity_score"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	GenresName      string    `json:"genres_name"`
	ScheduleCount   int       `json:"schedule_count"`
}

type UpdateMovieRequest struct {
	Title           *string               `form:"title"`
	Synopsis        *string               `form:"synopsis"`
	Duration        *int                  `form:"duration"`
	ReleaseDate     *string               `form:"release_date"`
	DirectorId      *int                  `form:"director_id"`
	Poster          *multipart.FileHeader `form:"poster"`
	Backdrop        *multipart.FileHeader `form:"backdrop"`
	PosterUrl       *string               `form:"-"`
	BackdropUrl     *string               `form:"-"`
	PopularityScore *float64              `form:"popularity_score"`
}

type UpdateMovieResponse struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Synopsis        string    `json:"synopsis"`
	Duration        int       `json:"duration"`
	ReleaseDate     time.Time `json:"release_date"`
	DirectorId      int       `json:"director_id"`
	PosterUrl       string    `json:"poster_url"`
	BackdropUrl     string    `json:"backdrop_url"`
	PopularityScore float64   `json:"popularity_score"`
}

type CreateMovieRequest struct {
	Title           *string               `form:"title"`
	Synopsis        *string               `form:"synopsis"`
	Duration        *int                  `form:"duration"`
	ReleaseDate     *string               `form:"release_date"`
	DirectorId      *int                  `form:"director_id"`
	Poster          *multipart.FileHeader `form:"poster"`
	Backdrop        *multipart.FileHeader `form:"backdrop"`
	PosterUrl       *string               `form:"-"`
	BackdropUrl     *string               `form:"-"`
	GenreIds        string                `form:"genre_ids"`
	Genres          []int                 `form:"-"`
	PopularityScore *float64              `form:"popularity_score"`
}

type CreateMovieResponse struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Synopsis        string    `json:"synopsis"`
	Duration        int       `json:"duration"`
	ReleaseDate     time.Time `json:"release_date"`
	DirectorId      int       `json:"director_id"`
	PosterUrl       string    `json:"poster_url"`
	BackdropUrl     string    `json:"backdrop_url"`
	PopularityScore float64   `json:"popularity_score"`
}
