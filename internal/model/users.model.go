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
