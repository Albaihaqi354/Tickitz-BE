package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/pkg"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	authRepository *repository.AuthRepository
	redis          *redis.Client
}

func NewAuthService(authRepository *repository.AuthRepository, rdb *redis.Client) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		redis:          rdb,
	}
}

func (a AuthService) Register(ctx context.Context, newUser dto.NewUser) (dto.RegisterResponse, error) {
	hc := pkg.HashConfig{}
	hc.UseRecomended()

	hp, err := hc.GenHash(newUser.Password)
	if err != nil {
		log.Println(err.Error())
		return dto.RegisterResponse{}, err
	}

	data, err := a.authRepository.CreateNewUser(ctx, newUser, hp)
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
		log.Println(err.Error())
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	hc := pkg.HashConfig{}
	hc.UseRecomended()

	isValid, err := hc.ComparePwdAndHash(loginReq.Password, user.Password)
	if err != nil {
		log.Println(err.Error())
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	if !isValid {
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	jwtClaim := pkg.NewJWTClaim(user.Id, user.Email, user.Role)
	token, err := jwtClaim.GetToken()
	if err != nil {
		log.Println(err.Error())
		return dto.LoginResponse{}, errors.New("internal server error")
	}

	err = a.authRepository.SaveToken(ctx, token, time.Hour)
	if err != nil {
		log.Println("Error save token: ", err.Error())
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

func (a AuthService) Logout(ctx context.Context, token string) error {
	return a.authRepository.DeleteToken(ctx, token)
}
