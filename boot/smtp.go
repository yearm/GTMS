package boot

import (
	"GTMS/conf"
	"gopkg.in/gomail.v2"
)

//发送邮件
func SendEmail(receiverAddress string, receiverName string, subject string, body string) (err error) {
	cfg := conf.GetSmtpConfig()
	m := gomail.NewMessage()
	m.SetAddressHeader("From", cfg.Address, cfg.UserName)             //发件人
	m.SetHeader("To", m.FormatAddress(receiverAddress, receiverName)) //收件人
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.UserName, cfg.Password)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true} //取消证书的验证(阿里云默认禁用25端口)
	if err = d.DialAndSend(m); err != nil {
		return
	}
	return
}
