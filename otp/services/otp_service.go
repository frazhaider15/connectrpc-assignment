package services

import "fmt"

func SendOtp(log OtpLog) {
	var body string
	if log.LogType == "USER_SIGNUP" {
		body = fmt.Sprintf("Enter this otp to signup: %v", log.Otp)
	}
	if log.LogType == "USER_LOGIN" {
		body = fmt.Sprintf("Enter this otp to login: %v", log.Otp)
	}
	err := SendTwilioSms(log.Phone, body)
	if err != nil {
		fmt.Println(err)
	}
}
