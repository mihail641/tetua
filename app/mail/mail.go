package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/logger"
	"github.com/ngocphuongnb/tetua/app/server"
	"mime"
	"net/mail"
	"net/smtp"
)

//channel for recording messages for sending
var messageQueue = make(chan letterInfo)

//letterInfo struct for recording letters
type letterInfo struct {
	receiverName    string
	receiverAddress string
	subject         string
	body            string
	request         string
}

//Send check config -file and returning error if it is empty, calls the goroutine that wrires
//letterInfo struct in chanel messageQueue
func Send(c server.Context, receiverName, receiverAddress, subject, body string) error {
	var err error

	if config.Mail == nil ||
		config.Mail.Sender == "" ||
		config.Mail.Smtp == nil ||
		config.Mail.Smtp.Host == "" ||
		config.Mail.Smtp.Port == 0 ||
		config.Mail.Smtp.User == "" ||
		config.Mail.Smtp.Pass == "" {

		return fmt.Errorf("Mail config is not set")
	}
	go func() {
		letter := letterInfo{
			receiverName:    receiverName,
			receiverAddress: receiverAddress,
			subject:         subject,
			body:            body,
			request:         c.RequestID(),
		}
		messageQueue <- letter
	}()
	return err
}

//startMessageSender-checks chanel messageQueue and takes struct letterInfo
func startMessageSender() {
	for message := range messageQueue {
		err := sendLetter(message)
		if err != nil {
			logger.Get().Error("request_id:", message.request, err)
		}
	}
}

//sendLetter sends letter to new users
func sendLetter(letter letterInfo) error {
	from := mail.Address{
		Name:    config.Setting("app_name"),
		Address: config.Mail.Sender,
	}
	to := mail.Address{
		Name:    letter.receiverName,
		Address: letter.receiverAddress,
	}
	tlsconfig := &tls.Config{
		InsecureSkipVerify: config.DEVELOPMENT,
		ServerName:         config.Mail.Smtp.Host,
	}

	message := fmt.Sprintf("%s: %s\r\n", "From", from.String())
	message += fmt.Sprintf("%s: %s\r\n", "To", to.String())
	message += fmt.Sprintf("%s: %s\r\n", "Subject", mime.QEncoding.Encode("UTF-8", letter.subject))
	message += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	message += "\r\n" + letter.body
	auth := smtp.PlainAuth("", config.Mail.Smtp.User, config.Mail.Smtp.Pass, config.Mail.Smtp.Host)
	address := fmt.Sprintf("%s:%d", config.Mail.Smtp.Host, config.Mail.Smtp.Port)

	tlsConn, err := tls.Dial("tcp", address, tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(tlsConn, config.Mail.Smtp.Host)
	if err != nil {
		return err
	}
	defer client.Quit()
	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(from.Address); err != nil {
		return err
	}

	if err = client.Rcpt(to.Address); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	if _, err = w.Write([]byte(message)); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}
	return nil
}

//initiates the startMessageSender goroutine when the first user registers in the application for the first time
func init() {
	go startMessageSender()
}
