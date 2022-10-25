package main

import (
	"fmt"
	"net/smtp"
)

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth("",
		"deniserofeev84@gmail.com", "xxxxxxxxx", "smtp.gmail.com")

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail("smtp.gmail.com:587", auth, "deniserofeev84@gmail.com",
		[]string{"deniserofeev84@gmail.com"}, []byte(msg))

	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func main() {
	fmt.Println("Golang app sending email runing...")
	sendMailSimple("Another subject", "Another body", []string{"deniserofeev84@gmail.com"})
	fmt.Println("Check your email...")

}
