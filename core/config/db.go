package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDb membuat connection pool ke PostgreSQL.
// Mendukung Neon (serverless PostgreSQL) dengan dua cara:
//  1. DATABASE_URL — langsung pakai connection string dari Neon (sudah include sslmode=require)
//  2. Variabel terpisah DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME + SSL_MODE
func InitDb() (*pgxpool.Pool, error) {
	connString := buildConnString()

	// Konfigurasi pool — penting untuk serverless (Neon/Vercel)
	maxConns := getEnvInt("DB_MAX_CONNS", 10)
	minConns := getEnvInt("DB_MIN_CONNS", 2)
	maxConnLifetime := getEnvDuration("DB_MAX_CONN_LIFETIME", 30*time.Minute)
	maxConnIdleTime := getEnvDuration("DB_MAX_CONN_IDLE_TIME", 5*time.Minute)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse db config: %w", err)
	}

	config.MaxConns = int32(maxConns)
	config.MinConns = int32(minConns)
	config.MaxConnLifetime = maxConnLifetime
	config.MaxConnIdleTime = maxConnIdleTime
	config.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("create db pool: %w", err)
	}

	return pool, nil
}

// buildConnString menyusun connection string.
// Prioritas: DATABASE_URL > variabel terpisah (DB_USER dll).
func buildConnString() string {
	// Opsi 1: DATABASE_URL (cara yang disarankan Neon & Vercel)
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return dbURL
	}

	// Opsi 2: variabel terpisah
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		// Default: require untuk produksi (Neon butuh SSL)
		// Set SSL_MODE=disable di .env untuk development lokal
		sslMode = "require"
	}

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", user, pwd, host, port, dbName, sslMode)
}

func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func getEnvDuration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}
