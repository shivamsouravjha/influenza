package email

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendEmail(to, subject string) error {

	// Perform the GET request'
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: transport}
	response, err := client.Get("https://reqres.in/api/users/2")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	// // Print the response body as a string
	fmt.Println("Response:", string(body))
	// otp := services.GenerateOTP()

	// err := redis.RedisSession().Set(to, otp, 4*time.Hour).Err()
	// if err != nil {
	// 	return err
	// }
	// sa, _ := redis.RedisSession().Get(to).Result()
	// fmt.Println(sa)
	// err = services.SendEmail(to, subject, otp)
	// fmt.Println(err)
	return err
}
