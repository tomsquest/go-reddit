package output

import (
	"github.com/mgutz/logxi/v1"
	"github.com/tomsquest/go-reddit/assets"
	"github.com/tomsquest/go-reddit/reddit"
	"gopkg.in/gomail.v2"
	"io"
	"os"
	"strconv"
	"strings"
)

type SmtpOutput struct {
}

func (SmtpOutput) Out(subreddit reddit.Subreddit) error {
	html, err := templatize(subreddit)
	if err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", "tom@tomsquest.com")
	mail.SetHeader("To", "tom@tomsquest.com")
	mail.SetHeader("Subject", "Subreddit "+strings.ToUpper(subreddit.Name)+" - "+toDate(subreddit.CrawlDate))
	mail.SetBody("text/html", html)
	mail.Embed("reddit_logo.png", gomail.SetCopyFunc(func(w io.Writer) error {
		data, err := assets.Asset("assets/reddit_logo.png")
		w.Write(data)
		return err
	}))

	log.Info("Sending email", "subject", mail.GetHeader("Subject"))

	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	return gomail.NewDialer(host, port, user, pass).DialAndSend(mail)
}
