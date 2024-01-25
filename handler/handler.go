package handler

import (
	"github.com/FelipeDaros/gooportunes/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetSQLite()
}
