package service

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/redis/go-redis/v9"
)

type MovieService struct {
	movieRepository *repository.MovieRepository
	redis           *redis.Client
}

func NewMovieService(movieRepository *repository.MovieRepository, rdb *redis.Client) *MovieService {
	return &MovieService{
		movieRepository: movieRepository,
		redis:           rdb,
	}
}

func (s MovieService) GetUpcomingMovies(ctx context.Context) ([]dto.GetUpcomingMovie, error) {
	rkey := "bian:tickitz:upcommingMovie"
	rsc := s.redis.Get(ctx, rkey)

	if rsc.Err() == nil {
		var result []dto.GetUpcomingMovie
		cache, err := rsc.Bytes()
		if err != nil {
			log.Println(err.Error())
		} else {
			err := json.Unmarshal(cache, &result)
			if err != nil {
				log.Println(err.Error())
			} else {
				return result, nil
			}
		}
	}

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

	cachestr, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to marshal", err.Error())
	} else {
		status := s.redis.Set(ctx, rkey, string(cachestr), 24*time.Hour)
		if status.Err() != nil {
			log.Println("caching failed:", status.Err().Error())
		}
	}
	return response, nil
}

func (s MovieService) GetPopularMovie(ctx context.Context) ([]dto.GetPopularMovie, error) {
	rkey := "bian:tickitz:popularMovie"
	rsc := s.redis.Get(ctx, rkey)

	if rsc.Err() == nil {
		var result []dto.GetPopularMovie
		cache, err := rsc.Bytes()
		if err != nil {
			log.Println(err.Error())
		} else {
			err := json.Unmarshal(cache, &result)
			if err != nil {
				log.Println(err.Error())
			} else {
				return result, nil
			}
		}
	}
	if rsc.Err() == redis.Nil {
		log.Println("movie cache miss")
	}

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

	cachestr, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to marshal", err.Error())
	} else {
		status := s.redis.Set(ctx, rkey, string(cachestr), 24*time.Hour)
		if status.Err() != nil {
			log.Println("caching failed:", status.Err().Error())
		}
	}
	return response, nil
}

func (s MovieService) GetMovieWithFilter(ctx context.Context, search *string, genreIds []int, page int, limit int) ([]dto.GetMovieWitFilter, dto.PaginationMeta, error) {
	offset := (page - 1) * limit
	movies, err := s.movieRepository.GetMovieWithFilter(ctx, search, genreIds, limit, offset)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, dto.PaginationMeta{}, err
	}

	totalData, err := s.movieRepository.CountMovieWithFilter(ctx, search, genreIds)
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
