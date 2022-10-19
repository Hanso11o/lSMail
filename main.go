package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	//*config
	hostURL := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "deniserofeev84@gmail.com"
	password := "xxxxxxxxxxxxx"
	emailReceiver := "deniserofeev84@gmail.com"

	//*obj Auth
	emailAuth := smtp.PlainAuth("", emailSender, password, hostURL)

	//*create email
	msg := []byte("To: " + emailReceiver + "\r\n" + "Subject: " + "Hello den" + "\r\n" + "How are you doing")

	//*send mail

	fmt.Println(hostPort, emailReceiver, emailAuth, msg)
}
