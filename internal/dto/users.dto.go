package dto

import "time"

type User struct {
	Id            int       `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	ProfileImage  string    `json:"profile_image"`
	LoyaltyPoints int       `json:"loyalty_points"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type NewUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type RegisterResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type GetProfile struct {
	Id            int       `json:"id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	ProfileImage  string    `json:"profile_image"`
	LoyaltyPoints int       `json:"loyalty_points"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
}

type GetHistory struct {
	Id            int       `json:"id"`
	BookingCode   string    `json:"booking_code"`
	TotalPrice    int       `json:"total_price"`
	PaymentStatus string    `json:"payment_status"`
	CreatedAt     time.Time `json:"created_at"`
	MovieId       int       `json:"movie_id"`
	Title         string    `json:"title"`
	PosterUrl     string    `json:"poster_url"`
	CinemaName    string    `json:"cinema_name"`
	CinemaLogo    string    `json:"cinema_logo"`
	ShowDate      time.Time `json:"show_date"`
	ShowTime      time.Time `json:"show_time"`
	TicketCount   int       `json:"ticket_count"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
