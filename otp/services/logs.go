package services

type OtpLog struct {
	Otp     string `json:"otp"`
	Phone   string `json:"phone"`
	LogType string `json:"logType"`
}
