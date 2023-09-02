package helpers

import (
	"time"

	"github.com/shivamsouravjha/influenza/services"
)

func SendEmail(to, subject string) error {
	otp := services.GenerateOTP()
	RedisSession().Set(to, otp, 4*time.Hour)
	services.SendEmail(to, subject, otp)
	return nil
}
