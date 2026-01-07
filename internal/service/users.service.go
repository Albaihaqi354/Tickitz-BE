package service

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) AddUser(ctx context.Context, newUser dto.NewUser) (dto.User, error) {
	data, err := u.userRepository.CreateNewUser(ctx, newUser)
	if err != nil {
		log.Println(err.Error())
		return dto.User{}, err
	}

	response := dto.User{
		Id:       data.Id,
		Email:    data.Email,
		Password: data.Password,
	}

	return response, nil
}
