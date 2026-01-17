package service

import (
	"context"
	"log"
	"math"

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

func (s MovieService) GetMovieWithFilter(ctx context.Context, search *string, genreId *int, page int, limit int) ([]dto.GetMovieWitFilter, dto.PaginationMeta, error) {
	offset := (page - 1) * limit
	movies, err := s.movieRepository.GetMovieWithFilter(ctx, search, genreId, limit, offset)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, dto.PaginationMeta{}, err
	}

	totalData, err := s.movieRepository.CountMovieWithFilter(ctx, search, genreId)
	if err != nil {
		log.Println("Service Error (Count):", err.Error())
		return nil, dto.PaginationMeta{}, err
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	var response []dto.GetMovieWitFilter
	for _, m := range movies {
		response = append(response, dto.GetMovieWitFilter{
			Id:         m.Id,
			Title:      m.Title,
			PosterUrl:  m.PosterUrl,
			GenresName: m.GenresName,
		})
	}

	meta := dto.PaginationMeta{
		Page:      page,
		TotalPage: totalPage,
	}

	return response, meta, nil
}

func (s MovieService) GetMovieDetail(ctx context.Context, movieId int) ([]dto.GetMovieDetail, error) {
	movies, err := s.movieRepository.GetMovieDetail(ctx, movieId)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, err
	}

	var response []dto.GetMovieDetail
	for _, m := range movies {
		response = append(response, dto.GetMovieDetail{
			Id:          m.Id,
			Title:       m.Title,
			Synopsis:    m.Synopsis,
			Duration:    m.Duration,
			ReleaseDate: m.ReleaseDate,
			Director:    m.Director,
			Cast:        m.Cast,
			PosterUrl:   m.PosterUrl,
			BackDropUrl: m.BackdropUrl,
			GenresName:  m.GenresName,
		})
	}
	return response, nil
}
