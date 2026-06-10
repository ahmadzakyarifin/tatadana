package main

import (
	"github.com/ahmadzakyarifin/tatadana/config"
	infrastructure "github.com/ahmadzakyarifin/tatadana/internal/Infrastructure"
)

func main(){
	cfg := config.LoadConfig()
	_ = infrastructure.ConnectDB(cfg)
}