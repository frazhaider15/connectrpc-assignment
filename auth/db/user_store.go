package db

import (
	"github.com/auth/models"
)

func CreateUser(user *models.User) error {
	return DB.Create(&user).Error
}

func GetUserByPhone(phone string) (models.User, error) {
	var user models.User
	err := DB.First(&user, "phone = ?", phone).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func CreateUserOtpRequest(otp *models.UserOtp) error {
	return DB.Create(&otp).Error
}

func GetAllUserSignupOtps(phone string) ([]models.UserOtp, error) {
	otps := make([]models.UserOtp, 0)
	err := DB.Where("phone = ? and otp_type = ?", phone, "USER_SIGNUP").Find(&otps).Error
	if err != nil {
		return otps, err
	}
	return otps, nil
}

func GetAllUserLoginOtps(phone string) ([]models.UserOtp, error) {
	otps := make([]models.UserOtp, 0)
	err := DB.Where("phone = ? and otp_type = ?", phone, "USER_LOGIN").Find(&otps).Error
	if err != nil {
		return otps, err
	}
	return otps, nil
}

func VerifyUserPhone(phone string) error {
	err := DB.Model(&models.User{}).Where("phone = ?", phone).Update("phone_verified", true).Error
	if err != nil {
		return err
	}
	return nil
}

