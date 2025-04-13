package config

import (
	"os"
	"strconv"
)

type Config struct {
	Postgres  PostgresConfig
	Redis     RedisConfig
	JWTSecret string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	SSLMode  string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func Load() *Config {
	pgPort, err := strconv.Atoi(getEnv("PG_PORT", "5432"))
	if err != nil {
		pgPort = 5432
	}

	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		redisDB = 0
	}

	return &Config{
		Postgres: PostgresConfig{
			Host:     getEnv("PG_HOST", "localhost"),
			Port:     pgPort,
			User:     getEnv("PG_USER", "defaultUser"),
			DBName:   getEnv("PG_DBNAME", "defaultDB"),
			Password: getEnv("PG_PASSWORD", "defaultPassword"),
			SSLMode:  getEnv("PG_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
