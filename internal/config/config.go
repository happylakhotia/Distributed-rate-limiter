package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port       string
	RedisAddr  string
	RateLimit  int
	RateWindow time.Duration
}

func Load() *Config {
	limit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if limit == 0 {
		limit = 10
	}

	window, _ := strconv.Atoi(os.Getenv("RATE_WINDOW_SECONDS"))
	if window == 0 {
		window = 60
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:       port,
		RedisAddr:  os.Getenv("REDIS_ADDR"),
		RateLimit:  limit,
		RateWindow: time.Duration(window) * time.Second,
	}
}


