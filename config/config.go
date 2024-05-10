package config

import "gorm.io/gorm"

var ( // config variable "global"
	db *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger = newLogger(p)
	return logger
}
