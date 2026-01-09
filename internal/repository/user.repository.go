package repository

import (
	"context"

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
