package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/peruccii/gopportunities/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
		_ "github.com/joho/godotenv/autoload"
)

func InitializeMysql() (*gorm.DB, error){
	err := godotenv.Load()
	if err != nil {
			fmt.Printf("Error loading .env file %v", err)
	}
	logger := GetLogger("MYSQL")

	url := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		logger.ErrF("Failed to open database %v", err)
		return nil, err
	}
	// Migrate the entity
	err =db.AutoMigrate(&entity.Opening{})

	return db, err
}