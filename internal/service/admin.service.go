package service

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
)

type AdminService struct {
	adminRepository *repository.AdminRepository
}

func NewAdminService(adminRepository *repository.AdminRepository) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
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
	return nil
}
