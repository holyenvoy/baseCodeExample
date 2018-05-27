package main

import (
	"gopkg.in/gomail.v2"
)

func main() {

m := gomail.NewMessage()
m.SetHeader("From", "wudebao5220150@126.com")
m.SetHeader("To", "wudebao5220150@163.com", "472119740@qq.com")
m.SetAddressHeader("Cc", "wudebao5220150@126.com", "Dan")
m.SetHeader("Subject", "Hello!")
m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

d := gomail.NewDialer("smtp.163.com", 25, "wudebao5220150@163.com", "wdb5221461121319")

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
    panic(err)
}

}


