package database

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB

	_ = godotenv.Load()
)
