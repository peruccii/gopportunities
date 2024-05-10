package config

import (
	"fmt"

	"gorm.io/gorm"
)

var ( // config variable "global"
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize sqlserver
	db, err = InitializeMysql()

	if err != nil {
		return fmt.Errorf("err initializing sqlserver: %v", err)
	}

	return nil
}

func GetSqlServer() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = newLogger(p)
	return logger
}
