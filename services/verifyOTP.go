package services

import (
	"errors"
	"strconv"

	"github.com/shivamsouravjha/influenza/helpers/redis"
	requestStruct "github.com/shivamsouravjha/influenza/structure/request"
)

func Verifycode(request requestStruct.OTPRequest) error {
	otp, err := redis.RedisSession().Get(request.Email).Result()
	if err != nil {
		return err
	}

	if otp != strconv.Itoa(request.OTP) {
		return errors.New("OTP didn't match, email isn't verified")
	}

	return nil
}
