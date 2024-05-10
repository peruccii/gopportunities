package config

import (
	"os"

	"github.com/peruccii/gopportunities/entity"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitializeMysql() (*gorm.DB, error){
	logger := GetLogger("MYSQL")
	db_path := "./db/main.db"
	// Check if the database file exists

	_, err := os.Stat(db_path)

	if os.IsNotExist(err) {
		logger.Info("DATABASE FILE NOT FOUND, CREATING....")
		// Create the database file and directory

		err = os.MkdirAll("./db", os.ModeAppend)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(db_path)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlserver.Open(db_path), &gorm.Config{})

	if err != nil {
		logger.ErrF("Failed to open database %v", err)
		return nil, err
	}
	// Migrate the entity
	err =db.AutoMigrate(&entity.Opening{})

	return db, err
}