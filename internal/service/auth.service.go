package service

import (
	"context"
	"errors"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/pkg"
)

type AuthService struct {
	authRepository *repository.AuthRepository
}

func NewAuthService(authRepository *repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (a AuthService) Register(ctx context.Context, newUser dto.NewUser) (dto.RegisterResponse, error) {
	hashConfig := &pkg.HashConfig{}
	hashConfig.UseRecomended()

	hashedPassword, err := hashConfig.GenHash(newUser.Password)
	if err != nil {
		log.Println("Error hashing password:", err.Error())
		return dto.RegisterResponse{}, err
	}

	newUser.Password = hashedPassword

	data, err := a.authRepository.CreateNewUser(ctx, newUser)
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

func (a AuthService) Login(ctx context.Context, loginReq dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := a.authRepository.FindUserByEmail(ctx, loginReq.Email)
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

	jwtClaim := pkg.NewJWTClaim(user.Id, user.Email, user.Role)
	token, err := jwtClaim.GetToken()
	if err != nil {
		log.Println("Error generating token:", err.Error())
		return dto.LoginResponse{}, errors.New("internal server error")
	}

	response := dto.LoginResponse{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}

	return response, nil
}
