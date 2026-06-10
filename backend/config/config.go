package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	DB_HOST   string
	DB_PORT   string
	JWTSECRET string
	ACCESS_TOKEN_EXPIRE int
	REFRESH_TOKEN_EXPIRE int	
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	cfg.DB_USER = os.Getenv("DB_USER")
	cfg.DB_PASS = os.Getenv("DB_PASS")
	cfg.DB_NAME = os.Getenv("DB_NAME")
	cfg.DB_HOST = os.Getenv("DB_HOST")
	cfg.DB_PORT = os.Getenv("DB_PORT")
	cfg.JWTSECRET = os.Getenv("JWTSECRET")
	cfg.ACCESS_TOKEN_EXPIRE, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRE"))
	cfg.REFRESH_TOKEN_EXPIRE, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRE"))

	return &cfg
}