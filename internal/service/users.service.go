package service

import (
	"context"
	"errors"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/pkg"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) GetProfile(ctx context.Context, userId int) (dto.GetProfile, error) {
	user, err := u.userRepository.GetProfile(ctx, userId)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return dto.GetProfile{}, err
	}
	response := dto.GetProfile{
		Id:            user.Id,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		PhoneNumber:   user.PhoneNumber,
		ProfileImage:  user.ProfileImage,
		LoyaltyPoints: user.LoyaltyPoints,
		Role:          user.Role,
		CreatedAt:     user.CreatedAt,
	}

	return response, nil
}

func (u UserService) GetHistory(ctx context.Context, userId int) ([]dto.GetHistory, error) {
	histories, err := u.userRepository.GetHistory(ctx, userId)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, err
	}

	var response []dto.GetHistory
	for _, history := range histories {
		h := dto.GetHistory{
			Id:            history.Id,
			BookingCode:   history.BookingCode,
			TotalPrice:    history.TotalPrice,
			PaymentStatus: history.PaymentStatus,
			CreatedAt:     history.CreatedAt,
			MovieId:       history.MovieId,
			Title:         history.Title,
			PosterUrl:     history.PosterUrl,
			CinemaName:    history.CinemaName,
			CinemaLogo:    history.CinemaLogo,
			ShowDate:      history.ShowDate,
			ShowTime:      history.ShowTime,
			TicketCount:   history.TicketCount,
		}
		response = append(response, h)
	}

	return response, nil
}

func (u UserService) UpdatePassword(ctx context.Context, userId int, req dto.UpdatePasswordRequest) error {
	currentHashedPassword, err := u.userRepository.GetPasswordById(ctx, userId)
	if err != nil {
		log.Println("Error fetching password from DB:", err.Error())
		return errors.New("internal server error")
	}

	hashConfig := &pkg.HashConfig{}
	hashConfig.UseRecomended()

	isValid, err := hashConfig.ComparePwdAndHash(req.OldPassword, currentHashedPassword)
	if err != nil {
		log.Println("Error comparing passwords:", err.Error())
		return errors.New("internal server error")
	}

	if !isValid {
		return errors.New("invalid old password")
	}

	newHashedPassword, err := hashConfig.GenHash(req.NewPassword)
	if err != nil {
		log.Println("Error hashing new password:", err.Error())
		return errors.New("internal server error")
	}

	err = u.userRepository.UpdatePassword(ctx, userId, newHashedPassword)
	if err != nil {
		log.Println("Error updating password in DB:", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (u UserService) UpdateProfile(ctx context.Context, userId int, req dto.UpdateProfileRequest) (dto.UpdateProfileResponse, error) {
	updateProfile, err := u.userRepository.UpdateProfile(ctx, userId, req)
	if err != nil {
		log.Println("Service Error", err.Error())
		return dto.UpdateProfileResponse{}, err
	}

	response := dto.UpdateProfileResponse{
		Id:          updateProfile.Id,
		FirstName:   updateProfile.FirstName,
		LastName:    updateProfile.LastName,
		PhoneNumber: updateProfile.PhoneNumber,
	}

	return response, nil
}
