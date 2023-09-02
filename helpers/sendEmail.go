package helpers

import "github.com/shivamsouravjha/influenza/services"

func SendEmail(to, subject string) error {
	otp := services.GenerateOTP()
	services.SendEmail(to, subject, otp)
	return nil
}
