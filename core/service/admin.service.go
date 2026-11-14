package service

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/core/dto"
	"github.com/Albaihaqi354/Tickitz-BE/core/repository"
	"github.com/redis/go-redis/v9"
)

type AdminService struct {
	adminRepository *repository.AdminRepository
	redis           *redis.Client
}

func NewAdminService(adminRepository *repository.AdminRepository, rdb *redis.Client) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
		redis:           rdb,
	}
}

// invalidateMovieCache menghapus cache movie di Redis setelah operasi CRUD admin.
func (a AdminService) invalidateMovieCache(ctx context.Context) {
	if a.redis == nil {
		return
	}
	keys := []string{
		"bian:tickitz:upcommingMovie",
		"bian:tickitz:popularMovie",
	}
	for _, key := range keys {
		if err := a.redis.Del(ctx, key).Err(); err != nil {
			log.Println("Warning: failed to invalidate cache key:", key, err.Error())
		}
	}
}

func (a AdminService) GetAllMovieAdmin(ctx context.Context) ([]dto.GetAllMovieAdmin, error) {
	movies, err := a.adminRepository.GetAllMovieAdmin(ctx)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, err
	}

	response := make([]dto.GetAllMovieAdmin, 0, len(movies))
	for _, m := range movies {
		response = append(response, dto.GetAllMovieAdmin{
			Id:              m.Id,
			Title:           m.Title,
			Synopsis:        m.Synopsis,
			Duration:        m.Duration,
			ReleaseDate:     m.ReleaseDate,
			Director:        m.Director,
			Cast:            m.Cast,
			PosterUrl:       m.PosterUrl,
			BackdropUrl:     m.BackdropUrl,
			PopularityScore: m.PopularityScore,
			CreatedAt:       m.CreatedAt,
			UpdatedAt:       m.UpdatedAt,
			GenresName:      m.GenresName,
			ScheduleCount:   m.ScheduleCount,
		})
	}
	return response, nil
}

func (a AdminService) DeleteMovieAdmin(ctx context.Context, movieId int) error {
	err := a.adminRepository.DeleteMovieAdmin(ctx, movieId)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return err
	}
	a.invalidateMovieCache(ctx)
	return nil
}

func (a AdminService) UpdateMovieAdmin(ctx context.Context, id int, req dto.UpdateMovieRequest) (dto.UpdateMovieResponse, error) {
	updatedMovie, err := a.adminRepository.UpdateMovieAdmin(ctx, id, req)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return dto.UpdateMovieResponse{}, err
	}

	a.invalidateMovieCache(ctx)

	response := dto.UpdateMovieResponse{
		Id:              updatedMovie.Id,
		Title:           updatedMovie.Title,
		Synopsis:        updatedMovie.Synopsis,
		Duration:        updatedMovie.Duration,
		ReleaseDate:     updatedMovie.ReleaseDate,
		DirectorId:      updatedMovie.DirectorId,
		PosterUrl:       updatedMovie.PosterUrl,
		BackdropUrl:     updatedMovie.BackdropUrl,
		PopularityScore: updatedMovie.PopularityScore,
	}

	return response, nil
}

func (a AdminService) CreateMovieAdmin(ctx context.Context, req dto.CreateMovieRequest) (dto.CreateMovieResponse, error) {
	newMovie, err := a.adminRepository.CreateMovieAdmin(ctx, req)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return dto.CreateMovieResponse{}, err
	}

	a.invalidateMovieCache(ctx)

	response := dto.CreateMovieResponse{
		Id:              newMovie.Id,
		Title:           newMovie.Title,
		Synopsis:        newMovie.Synopsis,
		Duration:        newMovie.Duration,
		ReleaseDate:     newMovie.ReleaseDate,
		DirectorId:      newMovie.DirectorId,
		PosterUrl:       newMovie.PosterUrl,
		BackdropUrl:     newMovie.BackdropUrl,
		PopularityScore: newMovie.PopularityScore,
	}

	return response, nil
}
