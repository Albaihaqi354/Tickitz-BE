package repository

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateNewUser(ctx context.Context, newUser dto.NewUser) (model.User, error) {
	sql := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password, role"
	values := []any{newUser.Email, newUser.Password}

	row := u.db.QueryRow(ctx, sql, values...)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepository) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	sql := "SELECT id, email, password, role FROM users WHERE email = $1"

	row := u.db.QueryRow(ctx, sql, email)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepository) GetProfile(ctx context.Context, userId int) (model.User, error) {
	sqlStr := `
		SELECT 
			id, 
			email, 
			first_name, 
			last_name, 
			phone_number, 
			profile_image, 
			loyalty_points, 
			role,
			created_at
		FROM users
		WHERE id = $1;`

	row := u.db.QueryRow(ctx, sqlStr, userId)

	var user model.User
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.ProfileImage,
		&user.LoyaltyPoints,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		log.Println("Scan error:", err.Error())
		return model.User{}, err
	}

	return user, nil
}
