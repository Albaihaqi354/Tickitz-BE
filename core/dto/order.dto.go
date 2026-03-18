package dto

import "time"

type GetSchedules struct {
	Id             int       `json:"id"`
	ShowDate       time.Time `json:"show_date"`
	ShowTime       time.Time `json:"show_time"`
	Price          int       `json:"price"`
	CinemaId       int       `json:"cinema_id"`
	CinemaName     string    `json:"cinema_name"`
	CinemaLogo     string    `json:"cinema_logo"`
	CinemaLocation string    `json:"cinema_location"`
	CinemaCity     string    `json:"cinema_city"`
}

type CreateOrderRequest struct {
	ScheduleId    int    `json:"schedule_id" binding:"required"`
	Seats         []int  `json:"seats" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

type SeatResponse struct {
	SeatId     int    `json:"seat_id"`
	RowLetter  string `json:"row_letter"`
	SeatNumber int    `json:"seat_number"`
	SeatType   string `json:"seat_type"`
	Status     string `json:"status"`
}
type UpdateOrderRequest struct {
	PaymentStatus string `json:"payment_status" binding:"required"`
}
