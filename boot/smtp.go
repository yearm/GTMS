package boot

import (
	"GTMS/conf"
	"fmt"
	"github.com/go-gomail/gomail"
)

//发送邮件
func SendEmail(receiverAddress string, receiverName string, subject string, body string) (err error) {
	fmt.Println(body)
	cfg := conf.GetSmtpConfig()
	m := gomail.NewMessage()
	m.SetAddressHeader("From", cfg.Address, cfg.UserName) //发件人
	m.SetHeader("To", receiverAddress, receiverName)      //收件人
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewPlainDialer(cfg.Host, cfg.Port, cfg.UserName, cfg.Password)
	if err = d.DialAndSend(m); err != nil {
		return
	}
	return
}
