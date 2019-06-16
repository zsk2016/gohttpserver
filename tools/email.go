package tools

import (
	"gopkg.in/gomail.v2"
)

type EmailInfo struct {
	SendInfo string
	SendAddr string
	FromAddr string
}

func SendEmail(ei *EmailInfo) bool {
	m := gomail.NewMessage()

	m.SetHeader("From", ei.FromAddr)
	m.SetHeader("To", ei.SendAddr)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "树PDF产品开通信息通知!")
	m.SetBody("text/html", ei.SendInfo)
	//m.Attach("/home/Alex/lolcat.jpg")
	//"qqywlzkjfpqtcaef"
	d := gomail.NewDialer("smtp.qq.com", 587, ei.FromAddr, "jktpkppepbugecci")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return false
	} else {
		return true
	}

}
