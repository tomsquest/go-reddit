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
	cfg config.SmtpConfig
}

func NewSmtpOutput(cfg config.SmtpConfig) *SmtpOutput {
	return &SmtpOutput{cfg}
}

func (out SmtpOutput) Out(subreddit reddit.Subreddit) error {
	html, err := templatize(subreddit)
	if err != nil {
		return err
	}

	from := out.cfg.From
	to := out.cfg.To
	subject := "Subreddit " + strings.ToUpper(subreddit.Name) + " - " + toDate(subreddit.CrawlDate)

	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
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

	log.Info("Sending email", "From", from, "To", to, "Subject", subject)

	return gomail.NewDialer(out.cfg.Host, out.cfg.Port, out.cfg.User, out.cfg.Pass).DialAndSend(mail)
}
