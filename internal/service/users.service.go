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
