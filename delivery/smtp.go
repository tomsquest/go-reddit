package delivery

import (
	"github.com/tomsquest/go-reddit/reddit"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
	"strings"
)

func NewSmtpSender() SmtpSender {
	return SmtpSender{}
}

type SmtpSender struct {
}

func (sender *SmtpSender) Send(subreddit reddit.Subreddit) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "tom@tomsquest.com")
	m.SetHeader("To", "tom@tomsquest.com")
	m.SetHeader("Subject", "Subreddit "+strings.ToUpper(subreddit.Name)+" - "+subreddit.CrawlDate.Format("2006-01-02"))
	m.SetBody("text/html", "Hello Subreddit "+subreddit.Name+" - "+subreddit.CrawlDate.String())

	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	return gomail.NewDialer(host, port, user, pass).DialAndSend(m)
}
