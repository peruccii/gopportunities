package handler

import (
	"github.com/peruccii/gopportunities/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func InitHandler() {
	Logger = config.GetLogger("handler")
	Db = config.GetSqlServer()
}