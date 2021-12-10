package tool

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo string, subject string, body string) error {

	mailConn := GetConfig()

	port, _ := strconv.Atoi(mailConn.Email.Port) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn.Email.User, "xmzhang")) //设置邮件发送人别名容易发送入垃圾箱
	// 这种方式可以添加别名，即“go的慢慢学习路”
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo)       //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn.Email.Host, port, mailConn.Email.User, mailConn.Email.Pass)

	err := d.DialAndSend(m)
	return err

}

func main() {
	err := SendMail("25337942@qq.com", "任务错误警报", "<h1>任务:1运行错误,请及时修复</h1>")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("send mail ok!")
	}

}
