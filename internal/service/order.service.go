package service

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/Albaihaqi354/Tickitz-BE/internal/repository"
)

type OrderService struct {
	orderRepository repository.OrderRepo
	db              *pgxpool.Pool
}

func NewOrderService(orderRepository repository.OrderRepo, db *pgxpool.Pool) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		db:              db,
	}
}

func (o OrderService) GetSchedules(ctx context.Context, movieId int, showDate *string, city *string) ([]dto.GetSchedules, error) {
	schedules, err := o.orderRepository.GetSchedules(ctx, o.db, movieId, showDate, city)
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
	seats, err := o.orderRepository.GetSeatsByScheduleID(ctx, o.db, scheduleId)
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
	price, err := o.orderRepository.GetPriceFromSchedule(ctx, o.db, req.ScheduleId)
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

	tx, err := o.db.Begin(ctx)
	if err != nil {
		log.Println("Service Error (Begin Tx):", err.Error())
		return 0, "", time.Time{}, err
	}
	defer tx.Rollback(ctx)

	id, bookingCode, createdAt, err := o.orderRepository.InsertOrder(ctx, tx, order)
	if err != nil {
		log.Println("Service Error (InsertOrder):", err.Error())
		return 0, "", time.Time{}, err
	}

	for _, seatId := range req.Seats {
		err := o.orderRepository.InsertOrderDetail(ctx, tx, id, seatId)
		if err != nil {
			log.Println("Service Error (InsertOrderDetail):", err.Error())
			return 0, "", time.Time{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Service Error (Commit Tx):", err.Error())
		return 0, "", time.Time{}, err
	}

	return id, bookingCode, createdAt, nil
}

func (o OrderService) UpdatePaymentStatus(ctx context.Context, orderId int, status string) error {
	err := o.orderRepository.UpdatePaymentStatus(ctx, o.db, orderId, status)
	if err != nil {
		log.Println("Service Error (UpdatePaymentStatus):", err.Error())
		return err
	}
	return nil
}
