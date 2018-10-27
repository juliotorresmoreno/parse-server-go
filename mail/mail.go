package mail

import (
	"bytes"
	"html/template"
	"log"
)

type Mail struct {
	From     string
	To       []string
	Attach   []byte
	CC       []string
	CCO      []string
	Template string
	Data     interface{}
}

func NewMail() Mail {
	return Mail{}
}

func (mail Mail) SendMail() error {
	t, err := template.ParseGlob("templates/mail/*.html")
	if err != nil {
		log.Println(err)
		return err
	}
	buff := bytes.Buffer{}
	err = t.ExecuteTemplate(&buff, mail.Template, mail.Data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
