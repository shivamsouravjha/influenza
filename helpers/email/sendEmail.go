package email

import (
	"time"

	"github.com/shivamsouravjha/influenza/helpers/redis"
	"github.com/shivamsouravjha/influenza/services"
)

func SendEmail(to, subject string) error {
	otp := services.GenerateOTP()

	err := redis.RedisSession().Set(to, otp, 4*time.Hour).Err()
	if err != nil {
		return err
	}

	err = services.SendEmail(to, subject, otp)
	return err
}
