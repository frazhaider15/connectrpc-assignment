package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone        string `gorm:"unique"`
	Name         string
	PhoneVerified bool
}

type UserOtp struct {
	gorm.Model
	Otp        string
	Phone      string
	OtpType    string
	ExpiryTime time.Time
}
