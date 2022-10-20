package main

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

func main() {

	msg := gomail.NewMessage()
	msg.SetHeader("From: ", "deniserofeev84@gmail.com")
	msg.SetHeader("To: ", "deniserofeev84@gmail.com")
	msg.SetHeader("Subject: ", "Test Email subject...msg")
	msg.SetBody("text/plain", "This is the test body")

	d := gomail.NewDialer("smtp.gmail.com", 587, "deniserofeev84@gmail.com", "Gd!2413@S%")

	if err := d.DialAndSend(msg); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Print("Email send...")
}

// *config
// hostURL := "smtp.gmail.com"
// hostPort := "587"
// emailSender := "deniserofeev84@gmail.com"
// password := "Gd!2413@S%"
// emailReceiver := "deniserofeev84@gmail.com"

// *obj Auth
// emailAuth := smtp.PlainAuth("", emailSender, password, hostURL)

// *create email
// msg := []byte("To: " + emailReceiver + "\r\n" + "Subject: " + "Hello den" + "\r\n" + "How are you doing")

// *send mail
// err := smtp.SendMail(hostURL+":"+hostPort, emailAuth, emailSender, emailReceiver, msg)
// if err != nil {
// 	fmt.Print("ERROR: ", err)

// }
