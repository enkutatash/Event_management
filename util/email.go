package util

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func  VerifyEmail(email string,token string) bool{
	
	m := gomail.NewMessage()
	m.SetHeader("From", "enkutatasheshetu96@gmail.com")
	m.SetHeader("To",email)
	m.SetHeader("Subject", "Your Verification Code")
	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"http://172.21.240.1:8081/verify?email=%s?token=%s\">here</a> to activate your account.", email,token))

	d := gomail.NewDialer("smtp.gmail.com", 587, "enkutatasheshetu96@gmail.com", "zksi jxum apfv idoy")

	err := d.DialAndSend(m)
	if err != nil {
		return false
	}
	return true
}