package model

import "time"

type User struct {
	Id            int       `db:"id"`
	Email         string    `db:"email"`
	Password      string    `db:"password"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	PhoneNumber   string    `db:"phone_number"`
	ProfileImage  string    `db:"profile_image"`
	LoyaltyPoints int       `db:"loyalty_points"`
	Role          string    `db:"role"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type GetHistory struct {
	Id            int       `db:"id"`
	BookingCode   string    `db:"booking_code"`
	TotalPrice    int       `db:"total_price"`
	PaymentStatus string    `db:"payment_status"`
	CreatedAt     time.Time `db:"created_at"`
	MovieId       int       `db:"movie_id"`
	Title         string    `db:"title"`
	PosterUrl     string    `db:"poster_url"`
	CinemaName    string    `db:"cinema_name"`
	CinemaLogo    string    `db:"cinema_logo"`
	ShowDate      time.Time `db:"show_date"`
	ShowTime      time.Time `db:"show_time"`
	TicketCount   int       `db:"ticket_count"`
}
