package db

import (
	"github.com/auth/models"
)

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserOtp{})
}
