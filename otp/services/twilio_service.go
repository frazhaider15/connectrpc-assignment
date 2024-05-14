package services

import (
	"fmt"
	"os"

	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

var TwilioClient *twilio.RestClient


func ConnectTwilio() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	TwilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func SendTwilioSms(toNumber, body string) error {
	fromNumber := os.Getenv("TWILIO_FROM_NUMBER")
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(fromNumber)
	params.SetBody(body)

	_, err := TwilioClient.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("Error sending SMS message: " + err.Error())
	}
	return nil
}
