package schedules

import (
	"bytes"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)

func WeeklyReport() {
	t, _ := template.ParseFiles("mail.html")
	var body bytes.Buffer

	err := t.Execute(&body, struct {
		Name    string
		Email   string
		Address string
	}{
		Name:    "zww",
		Email:   "zhiweiwang2001@gmail.com",
		Address: "中山大学",
	})
	if err != nil {
		fmt.Println("发生错误")
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "18827189341@163.com")
	m.SetHeader("To", "3440771474@qq.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(
		"smtp.163.com",
		465,
		"18827189341@163.com",
		"XMPNOYVLVYBMEASJ",
	)

	fmt.Println("ready to send!")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
