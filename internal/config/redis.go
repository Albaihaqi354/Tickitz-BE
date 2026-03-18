package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	// Jika RDS_URL tersedia, gunakan ParseURL (Sangat disarankan untuk Upstash/TLS)
	rdsURL := os.Getenv("RDS_URL")
	if rdsURL != "" {
		opt, err := redis.ParseURL(rdsURL)
		if err == nil {
			return redis.NewClient(opt)
		}
	}

	// Fallback ke manual config jika RDS_URL tidak ada
	user := os.Getenv("RDS_USER")
	pass := os.Getenv("RDS_PASS")
	host := os.Getenv("RDS_HOST")
	port := os.Getenv("RDS_PORT")
	dtbs := os.Getenv("RDS_DTBS")

	db, _ := strconv.Atoi(dtbs)

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: user,
		Password: pass,
		DB:       db,
	})
}
