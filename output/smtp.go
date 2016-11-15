package output

import (
	"github.com/mgutz/logxi/v1"
	"github.com/tomsquest/go-reddit/assets"
	"github.com/tomsquest/go-reddit/config"
	"github.com/tomsquest/go-reddit/reddit"
	"gopkg.in/gomail.v2"
	"io"
	"strings"
)

type SmtpOutput struct {
	host string
	port int
	user string
	pass string
}

func NewSmtpOutput(cfg config.Config) *SmtpOutput {
	return &SmtpOutput{
		host: cfg.Smtp.Host,
		port: cfg.Smtp.Port,
		user: cfg.Smtp.User,
		pass: cfg.Smtp.Pass,
	}
}

func (out SmtpOutput) Out(subreddit reddit.Subreddit) error {
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
	mail.Embed("reddit_no_thumbnail.png", gomail.SetCopyFunc(func(w io.Writer) error {
		data, err := assets.Asset("assets/reddit_no_thumbnail.png")
		w.Write(data)
		return err
	}))

	log.Info("Sending email", "subject", mail.GetHeader("Subject"))

	return gomail.NewDialer(out.host, out.port, out.user, out.pass).DialAndSend(mail)
}
