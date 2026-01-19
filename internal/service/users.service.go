package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
	"github.com/Albaihaqi354/Tickitz-BE/pkg"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	userRepository *repository.UserRepository
	redis          *redis.Client
}

func NewUserService(userRepository *repository.UserRepository, rdb *redis.Client) *UserService {
	return &UserService{
		userRepository: userRepository,
		redis:          rdb,
	}
}

func (u UserService) GetProfile(ctx context.Context, userId int) (dto.GetProfile, error) {
	rkey := fmt.Sprintf("bian:tickitz:users:%d", userId)
	rsc := u.redis.Get(ctx, rkey)

	if rsc.Err() == nil {
		var result dto.GetProfile
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
		log.Println("users cache miss")
	}

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

	cachestr, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to marshal:", err.Error())
	} else {
		status := u.redis.Set(ctx, rkey, string(cachestr), 1*time.Hour)
		if status.Err() != nil {
			log.Println("caching failed:", status.Err().Error())
		}
	}
	return response, nil
}

func (u UserService) GetHistory(ctx context.Context, userId int) ([]dto.GetHistory, error) {
	rkey := fmt.Sprintf("bian:tickitz:history:%d", userId)
	rsc := u.redis.Get(ctx, rkey)

	if rsc.Err() == nil {
		var result []dto.GetHistory
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
		log.Println("history cache miss")
	}

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

	cachestr, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to marshal:", err.Error())
	} else {
		status := u.redis.Set(ctx, rkey, string(cachestr), 30*time.Minute)
		if status.Err() != nil {
			log.Println("caching failed:", status.Err().Error())
		}
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

	rkey := fmt.Sprintf("bian:tickitz:users:%d", userId)
	u.redis.Del(ctx, rkey)

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

	rkey := fmt.Sprintf("bian:tickitz:users:%d", userId)
	u.redis.Del(ctx, rkey)

	return response, nil
}
