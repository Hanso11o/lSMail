package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Golang email app running...")
	email()
}

func email() {
	//*sender data
	from := os.Getenv("deniserofeev84@gmail.com")
	password := os.Getenv("")

	//*receiver address
	toEmail := os.Getenv("deniserofeev84@gmail.com")
	to := []string{"deniserofeev84@gmail.com"}

	//*smtp Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	//*message
	subject := "Subject:  Our Golang Email\n"
	body := "Our first email!"
	message := []byte(subject + body)

	//*authentication data

}

// msg := gomail.NewMessage()
// 	msg.SetHeader("From: ", "")
// 	msg.SetHeader("To: ", "deniserofeev84@gmail.com")
// 	msg.SetHeader("Subject: ", "Test Email subject...msg")
// 	msg.SetBody("text/plain", "This is the test body")

// 	d := gomail.NewDialer("smtp.gmail.com", 587, "deniserofeev84@gmail.com", "")

// 	if err := d.DialAndSend(msg); err != nil {
// 		fmt.Println(err)
// 		log.Fatal(err)
// 	}

// 	fmt.Print("Email send...")
// *config
// hostURL :=
// hostPort := "587"
// emailSender := "deniserofeev84@gmail.com"
// password := ""
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
