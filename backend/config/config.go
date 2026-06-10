package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	DBUser   string
	DBPass   string
	DBName   string
	DBHost   string
	DBPort   string
	JWTSecret string
	AccessTokenExpire int
	RefreshTokenExpire int
	AppEnv string	
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPass = os.Getenv("DB_PASS")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.JWTSecret = os.Getenv("JWTSECRET")
	cfg.AccessTokenExpire, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRE"))
	cfg.RefreshTokenExpire, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRE"))
	cfg.AppEnv = os.Getenv("APP_ENV")

	return &cfg
}