package repository

import (
	"context"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a AuthRepository) CreateNewUser(ctx context.Context, newUser dto.NewUser) (model.User, error) {
	sql := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password, role"
	values := []any{newUser.Email, newUser.Password}

	row := a.db.QueryRow(ctx, sql, values...)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a AuthRepository) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	sql := "SELECT id, email, password, role FROM users WHERE email = $1"

	row := a.db.QueryRow(ctx, sql, email)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}
