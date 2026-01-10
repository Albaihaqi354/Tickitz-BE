package service

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
)

type MovieService struct {
	movieRepository *repository.MovieRepository
}

func NewMovieService(movieRepository *repository.MovieRepository) *MovieService {
	return &MovieService{
		movieRepository: movieRepository,
	}
}

func (s MovieService) GetUpcomingMovies(ctx context.Context) ([]dto.GetUpcomingMovie, error) {
	movies, err := s.movieRepository.GetUpcomingMovie(ctx)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, err
	}

	var response []dto.GetUpcomingMovie
	for _, m := range movies {
		response = append(response, dto.GetUpcomingMovie{
			Id:          m.Id,
			Title:       m.Title,
			PosterUrl:   m.PosterUrl,
			ReleaseDate: m.ReleaseDate,
			GenresName:  m.GenresName,
		})
	}

	return response, nil
}

func (s MovieService) GetPopularMovie(ctx context.Context) ([]dto.GetPopularMovie, error) {
	movies, err := s.movieRepository.GetPopularMovie(ctx)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, err
	}

	var response []dto.GetPopularMovie
	for _, m := range movies {
		response = append(response, dto.GetPopularMovie{
			Id:         m.Id,
			Title:      m.Title,
			PosterUrl:  m.PosterUrl,
			GenresName: m.GenresName,
		})
	}
	return response, nil
}
