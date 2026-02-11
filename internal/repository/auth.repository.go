package repository

import (
	"context"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewAuthRepository(db *pgxpool.Pool, rdb *redis.Client) *AuthRepository {
	return &AuthRepository{
		db:    db,
		redis: rdb,
	}
}

func (a AuthRepository) CreateNewUser(ctx context.Context, newUser dto.NewUser, hashedPwd string) (model.User, error) {
	sql := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password, role"
	values := []any{newUser.Email, hashedPwd}

	row := a.db.QueryRow(ctx, sql, values...)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a AuthRepository) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	sql := "SELECT id, email, password, role FROM users WHERE email = $1"

	var user model.User
	if err := a.db.QueryRow(ctx, sql, email).Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a AuthRepository) SaveToken(ctx context.Context, token string, ttl time.Duration) error {
	rkey := "bian:tickitz:whitelist:" + token
	return a.redis.Set(ctx, rkey, "active", ttl).Err()
}

func (a AuthRepository) DeleteToken(ctx context.Context, token string) error {
	rkey := "bian:tickitz:whitelist:" + token
	return a.redis.Del(ctx, rkey).Err()
}

func (a AuthRepository) TokenWhitelist(ctx context.Context, token string) (bool, error) {
	rkey := "bian:tickitz:whitelist:" + token
	rsc, err := a.redis.Exists(ctx, rkey).Result()
	if err != nil {
		return false, err
	}
	return rsc > 0, nil
}
