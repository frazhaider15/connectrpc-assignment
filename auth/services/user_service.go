package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/auth/db"
	"github.com/auth/models"
	"github.com/auth/rabbitmq"
)

func UserSignup(name, phone string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if phone == "" {
		return fmt.Errorf("phone number is required")
	}
	newUser := models.User{
		Phone:         phone,
		Name:          name,
		PhoneVerified: false,
	}
	err := db.CreateUser(&newUser)
	if err != nil {
		return err
	}
	otp, _ := generateOtp(6)
	err = db.CreateUserOtpRequest(&models.UserOtp{
		Otp:        otp,
		Phone:      phone,
		OtpType:    "USER_SIGNUP",
		ExpiryTime: time.Now().Add(10 * time.Minute),
	})
	if err != nil {
		return err
	}
	log := rabbitmq.OtpLog{
		Otp:     otp,
		Phone:   phone,
		LogType: "USER_SIGNUP",
	}
	rabbitmq.PublishMessage(log)
	return nil
}

func VerfifySignupOtp(givenOtp, phone string) error {
	if givenOtp == "" {
		return fmt.Errorf("otp is required")
	}
	if phone == "" {
		return fmt.Errorf("phone number is required")
	}
	otps, err := db.GetAllUserSignupOtps(phone)
	if err != nil {
		return err
	}
	if len(otps) == 0 {
		return fmt.Errorf("no otp generated")
	}
	for _, otp := range otps {
		if otp.Otp == givenOtp {
			if time.Now().After(otp.ExpiryTime) {
				return fmt.Errorf("otp has expired")
			}
			err = db.VerifyUserPhone(phone)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("invalid otp")
}

func UserLogin(phone string) error {
	if phone == "" {
		return fmt.Errorf("phone number is required")
	}

	user, err := db.GetUserByPhone(phone)
	if err != nil {
		return err
	}
	if !user.PhoneVerified {
		return fmt.Errorf("user phone number not verified")
	}
	otp, _ := generateOtp(6)
	err = db.CreateUserOtpRequest(&models.UserOtp{
		Otp:        otp,
		Phone:      phone,
		OtpType:    "USER_LOGIN",
		ExpiryTime: time.Now().Add(10 * time.Minute),
	})
	if err != nil {
		return err
	}
	log := rabbitmq.OtpLog{
		Otp:     otp,
		Phone:   phone,
		LogType: "USER_LOGIN",
	}
	rabbitmq.PublishMessage(log)
	return nil
}

func VerifyLoginOtp(givenOtp, phone string) error {
	if givenOtp == "" {
		return fmt.Errorf("otp is required")
	}
	if phone == "" {
		return fmt.Errorf("phone number is required")
	}
	otps, err := db.GetAllUserLoginOtps(phone)
	if err != nil {
		return err
	}
	if len(otps) == 0 {
		return fmt.Errorf("no otp generated")
	}
	for _, otp := range otps {
		if otp.Otp == givenOtp {
			if time.Now().After(otp.ExpiryTime) {
				return fmt.Errorf("otp has expired")
			}
			return nil
		}
	}
	return fmt.Errorf("invalid otp")
}

func GetProfile(phone string) (models.User, error) {
	return db.GetUserByPhone(phone)
}

///////////////////////////////////////////////

func generateOtp(length int) (string, error) {
	charSet := "0123456789"
	var err error

	if len(charSet) == 0 {
		return "", errors.New("no character set selected")
	}

	b := make([]byte, length)
	_, err = rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("error generating random bytes: %w", err)
	}

	for i := range b {
		b[i] = charSet[int(b[i])%len(charSet)]
	}
	return string(b), nil
}
