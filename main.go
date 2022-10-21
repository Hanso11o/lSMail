package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type Person struct {
	Name string
}

func main() {

	fmt.Println("Email template app running...")
	To := os.Getenv("xxxx")
	emailTemplate(To)
}

func emailTemplate(To []string) (err error) {
	addr := "smtp.gmail.com:587"
	from := os.Getenv("xxxx")
	password := os.Getenv("")

	//*authentication data
	//*func PlainAuth(indentify, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, addr)

	var P Person
	P.Name = "Den"
	var t *template.Template

	t, err = t.ParseFiles("templates/bodyTemplate.html")
	if err != nil {
		fmt.Println((err))
		return
	}

	buff := new(bytes.Buffer)
	t.Execute(buff, P)
	fmt.Println(buff.String())

	subject := "Subject: HTML Template Email\n"
	mime := "Mime version: 1.0;\nContent-Type : text/html; charset=\"UTF-8\";\n\n"
	body := "<html><h1>Golang lsMail</h1><ul><li>Denis Erofeev</li>"
	msg := []byte(subject + mime + body)
	fmt.Println("message:", string(msg))

	//*send email
	//*func SendMail(address, auth, from,to, message)
	err = smtp.SendMail(addr, auth, from, To, msg)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("Go check your email")

}

// msg := gomail.NewMessage()
// 	msg.SetHeader("From: ", "")
// 	msg.SetHeader("To: ", "xxx")
// 	msg.SetHeader("Subject: ", "Test Email subject...msg")
// 	msg.SetBody("text/plain", "This is the test body")

// 	d := gomail.NewDialer("smtp.gmail.com", 587, "xxx", "")

// 	if err := d.DialAndSend(msg); err != nil {
// 		fmt.Println(err)
// 		log.Fatal(err)
// 	}

// 	fmt.Print("Email send...")
// *config
// hostURL :=
// hostPort := "587"
// emailSender := "xxx"
// password := ""
// emailReceiver := "xxx"

// *obj Auth
// emailAuth := smtp.PlainAuth("", emailSender, password, hostURL)

// *create email
// msg := []byte("To: " + emailReceiver + "\r\n" + "Subject: " + "Hello " + "\r\n" + "How are you doing")

// *send mail
// err := smtp.SendMail(hostURL+":"+hostPort, emailAuth, emailSender, emailReceiver, msg)
// if err != nil {
// 	fmt.Print("ERROR: ", err)

// }
