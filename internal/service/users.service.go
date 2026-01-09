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

func (u UserService) AddUser(ctx context.Context, newUser dto.NewUser) (dto.RegisterResponse, error) {
	hashConfig := &pkg.HashConfig{}
	hashConfig.UseRecomended()

	hashedPassword, err := hashConfig.GenHash(newUser.Password)
	if err != nil {
		log.Println("Error hashing password:", err.Error())
		return dto.RegisterResponse{}, err
	}

	newUser.Password = hashedPassword

	data, err := u.userRepository.CreateNewUser(ctx, newUser)
	if err != nil {
		log.Println(err.Error())
		return dto.RegisterResponse{}, err
	}

	response := dto.RegisterResponse{
		Id:    data.Id,
		Email: data.Email,
		Role:  data.Role,
	}

	return response, nil
}

func (u UserService) Login(ctx context.Context, loginReq dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := u.userRepository.FindUserByEmail(ctx, loginReq.Email)
	if err != nil {
		log.Println("User not found:", err.Error())
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	hashConfig := &pkg.HashConfig{}
	hashConfig.UseRecomended()

	isValid, err := hashConfig.ComparePwdAndHash(loginReq.Password, user.Password)
	if err != nil {
		log.Println("Error comparing password:", err.Error())
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	if !isValid {
		log.Println("Invalid password for user:", loginReq.Email)
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	response := dto.LoginResponse{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
	}

	return response, nil
}
