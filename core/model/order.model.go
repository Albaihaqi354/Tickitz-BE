package model

import "time"

type GetSchedules struct {
	Id             int       `db:"id"`
	ShowDate       time.Time `db:"show_date"`
	ShowTime       time.Time `db:"show_time"`
	Price          int       `db:"price"`
	CinemaId       int       `db:"cinema_id"`
	CinemaName     string    `db:"cinema_name"`
	CinemaLogo     string    `db:"cinema_logo"`
	CinemaLocation string    `db:"cinema_location"`
	CinemaCity     string    `db:"cinema_city"`
}

type Order struct {
	Id            int       `db:"id"`
	UserId        int       `db:"user_id"`
	ScheduleId    int       `db:"schedule_id"`
	BookingCode   string    `db:"booking_code"`
	TotalPrice    int       `db:"total_price"`
	PaymentMethod string    `db:"payment_method"`
	PaymentStatus string    `db:"payment_status"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type OrderDetail struct {
	Id        int       `db:"id"`
	OrderId   int       `db:"order_id"`
	SeatId    int       `db:"seat_id"`
	CreatedAt time.Time `db:"created_at"`
}

type Seat struct {
	SeatId     int    `db:"seat_id"`
	RowLetter  string `db:"row_letter"`
	SeatNumber int    `db:"seat_number"`
	SeatType   string `db:"seat_type"`
	Status     string `db:"status"`
}
