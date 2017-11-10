package networking

import (
	"net/smtp"
	"log"
)



func SendTestMail() {
	send("<b>https://support.google.com/mail/answer/78754</b>")
}

func send(body string) {
	from := "ovasp.net@gmail.com"
	pass := "getsetgo"
	to := "ovais.koti@continuum.net"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := "From: " + from + "\n" + mime +
		"To: " + to + "\n" +
		"Subject: Hello New there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://ovais.xyz")
}