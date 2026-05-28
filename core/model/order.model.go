package model

import "time"

type GetSchedules struct {
	Id             int       `db:"id" json:"id"`
	ShowDate       time.Time `db:"show_date" json:"show_date"`
	ShowTime       time.Time `db:"show_time" json:"show_time"`
	Price          int       `db:"price" json:"price"`
	CinemaId       int       `db:"cinema_id" json:"cinema_id"`
	CinemaName     string    `db:"cinema_name" json:"cinema_name"`
	CinemaLogo     string    `db:"cinema_logo" json:"cinema_logo"`
	CinemaLocation string    `db:"cinema_location" json:"cinema_location"`
	CinemaCity     string    `db:"cinema_city" json:"cinema_city"`
}

type Order struct {
	Id            int       `db:"id" json:"id"`
	UserId        int       `db:"user_id" json:"user_id"`
	ScheduleId    int       `db:"schedule_id" json:"schedule_id"`
	BookingCode   string    `db:"booking_code" json:"booking_code"`
	TotalPrice    int       `db:"total_price" json:"total_price"`
	PaymentMethod string    `db:"payment_method" json:"payment_method"`
	PaymentStatus string    `db:"payment_status" json:"payment_status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type OrderDetail struct {
	Id        int       `db:"id" json:"id"`
	OrderId   int       `db:"order_id" json:"order_id"`
	SeatId    int       `db:"seat_id" json:"seat_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Seat struct {
	SeatId     int    `db:"seat_id" json:"seat_id"`
	RowLetter  string `db:"row_letter" json:"row_letter"`
	SeatNumber int    `db:"seat_number" json:"seat_number"`
	SeatType   string `db:"seat_type" json:"seat_type"`
	Status     string `db:"status" json:"status"`
}
