package email

import (
	"gopkg.in/gomail.v2"
	"server/global"
)

type Subject string

const (
	Code  Subject = "验证码"
	Note  Subject = "操作通知"
	Alert Subject = "警告通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{Subject: Code}
}
func NewNote() Api {
	return Api{Subject: Note}
}

func NewAlert() Api {
	return Api{Subject: Alert}
}
func send(name, subject, body string) error {
	e := global.Config.Email
	return sendMail(e.Host, e.Port, e.User, e.Password, e.DefaultFromEmail, name, subject, body, e.UseSSL, e.UserTLs)
}
func sendMail(host string, port int, user, password, from, to, subject, body string, useSSL, userTLs bool) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(user, from)) // 发件人名称和发件人邮箱
	m.SetHeader("To", to)                            // 收件人
	m.SetHeader("Subject", subject)                  // 主题
	m.SetBody("text/html", body)                     // 正文
	d := gomail.NewDialer(host, port, user, password)
	err := d.DialAndSend(m)
	return err
}
