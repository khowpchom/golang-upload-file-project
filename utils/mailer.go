package utils

import (
	"log"

	"github.com/khowpchom/golang-upload-file-project/configs"
	"gopkg.in/gomail.v2"
)

type Mailer struct{}

func (m *Mailer) Send(message *gomail.Message) {
	message.SetHeader("From", "panforsendmail@gmail.com")

	if err := configs.Mailer.DialAndSend(message); err != nil {
		log.Panicln("[Mailer] ", err)
	}
}