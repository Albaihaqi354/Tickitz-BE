package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDb() (*pgxpool.Pool, error) {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, pwd, host, port, dbName)
	return pgxpool.New(context.Background(), connString)
}
