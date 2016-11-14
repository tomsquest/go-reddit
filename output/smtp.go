package output

import (
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

	m := gomail.NewMessage()
	m.SetHeader("From", "tom@tomsquest.com")
	m.SetHeader("To", "tom@tomsquest.com")
	m.SetHeader("Subject", "Subreddit "+strings.ToUpper(subreddit.Name)+" - "+toDate(subreddit.CrawlDate))
	m.SetBody("text/html", html)
	m.Embed("reddit_logo.png", gomail.SetCopyFunc(func(w io.Writer) error {
		data, err := assets.Asset("assets/reddit_logo.png")
		w.Write(data)
		return err
	}))

	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	return gomail.NewDialer(host, port, user, pass).DialAndSend(m)
}
