package service

import (
	"context"
	"log"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
}

func NewOrderService(orderRepository *repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (o OrderService) GetSchedules(ctx context.Context, movieId int, showDate *string, city *string) ([]dto.GetSchedules, error) {
	schedules, err := o.orderRepository.GetSchedules(ctx, movieId, showDate, city)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, err
	}

	var response []dto.GetSchedules
	for _, s := range schedules {
		response = append(response, dto.GetSchedules{
			Id:             s.Id,
			ShowDate:       s.ShowDate,
			ShowTime:       s.ShowTime,
			Price:          s.Price,
			CinemaId:       s.CinemaId,
			CinemaName:     s.CinemaName,
			CinemaLogo:     s.CinemaLogo,
			CinemaLocation: s.CinemaLocation,
			CinemaCity:     s.CinemaCity,
		})
	}
	return response, nil
}

func (o OrderService) GetSeatsByScheduleID(ctx context.Context, scheduleId int) ([]dto.SeatResponse, error) {
	seats, err := o.orderRepository.GetSeatsByScheduleID(ctx, scheduleId)
	if err != nil {
		log.Println("Service Error:", err.Error())
		return nil, err
	}

	var response []dto.SeatResponse
	for _, s := range seats {
		response = append(response, dto.SeatResponse{
			SeatId:     s.SeatId,
			RowLetter:  s.RowLetter,
			SeatNumber: s.SeatNumber,
			SeatType:   s.SeatType,
			Status:     s.Status,
		})
	}
	return response, nil
}

func (o OrderService) CreateOrder(ctx context.Context, userId int, req dto.CreateOrderRequest) (int, string, time.Time, error) {
	price, err := o.orderRepository.GetPriceFromSchedule(ctx, req.ScheduleId)
	if err != nil {
		log.Println("Service Error (GetPrice):", err.Error())
		return 0, "", time.Time{}, err
	}

	totalPrice := price * len(req.Seats)

	order := model.Order{
		UserId:        userId,
		ScheduleId:    req.ScheduleId,
		TotalPrice:    totalPrice,
		PaymentStatus: "pending",
	}

	id, bookingCode, createdAt, err := o.orderRepository.CreateOrder(ctx, order, req.Seats)
	if err != nil {
		log.Println("Service Error (CreateOrder):", err.Error())
		return 0, "", time.Time{}, err
	}

	return id, bookingCode, createdAt, nil
}
