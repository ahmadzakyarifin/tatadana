package infrastructure

import (
	"fmt"
	"log"

	"github.com/ahmadzakyarifin/tatadana/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewGormLogger(cfg.AppEnv),
	})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mengambil instance sql.DB: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Error saat melakukan ping database: %v", err)
	}

	fmt.Println("Connected to the database!")
	
	return db
}