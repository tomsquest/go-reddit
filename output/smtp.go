package output

import (
	"bytes"
	"github.com/tomsquest/go-reddit/assets"
	"github.com/tomsquest/go-reddit/reddit"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"strconv"
	"strings"
)

type SmtpOutput struct {
}

func (SmtpOutput) Out(subreddit reddit.Subreddit) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "tom@tomsquest.com")
	m.SetHeader("To", "tom@tomsquest.com")
	m.SetHeader("Subject", "Subreddit "+strings.ToUpper(subreddit.Name)+" - "+subreddit.CrawlDate.Format("2006-01-02"))
	m.SetBody("text/html", prepareHtml(subreddit))

	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	return gomail.NewDialer(host, port, user, pass).DialAndSend(m)
}

func prepareHtml(subreddit reddit.Subreddit) string {
	buf := new(bytes.Buffer)
	data := assets.MustAsset("assets/subreddit.html")

	funcs := template.FuncMap{
		"title":   strings.Title,
		"toUpper": strings.ToUpper,
	}

	tmpl := template.Must(template.New("console").Funcs(funcs).Parse(string(data)))
	tmpl.Execute(buf, subreddit)
	return buf.String()
}
