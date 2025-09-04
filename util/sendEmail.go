package util

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

var maxTrys = 3

func SendOTPEmail(to, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("FROM_EMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", "Your OTP code is: "+otp+"\nThis code expires in 5 minutes.")

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("FROM_EMAIL"), os.Getenv("EM_PASSWORD"))
	for i := 0; i < maxTrys; i++ {
		if err := d.DialAndSend(m); err == nil {
			log.Println("✅ Email sent successfully to", to)
			return nil
		}
	}
	return fmt.Errorf("failed to send email to %s after %d attempts", to, maxTrys)
}
