package output

import (
	"bytes"
	"github.com/tomsquest/go-reddit/assets"
	"github.com/tomsquest/go-reddit/reddit"
	"gopkg.in/gomail.v2"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type SmtpOutput struct {
}

func (SmtpOutput) Out(subreddit reddit.Subreddit) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "tom@tomsquest.com")
	m.SetHeader("To", "tom@tomsquest.com")
	m.SetHeader("Subject", "Subreddit "+strings.ToUpper(subreddit.Name)+" - "+toDate(subreddit.CrawlDate))
	m.SetBody("text/html", prepareHtml(subreddit))
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

func prepareHtml(subreddit reddit.Subreddit) string {
	buf := new(bytes.Buffer)
	data := assets.MustAsset("assets/subreddit.html")

	funcs := template.FuncMap{
		"title":   strings.Title,
		"toUpper": strings.ToUpper,
		"toDate":  toDate,
	}

	tmpl := template.Must(template.New("html").Funcs(funcs).Parse(string(data)))
	tmpl.Execute(buf, subreddit)
	return buf.String()
}

func toDate(date interface{}) string {
	var t time.Time
	switch date := date.(type) {
	case time.Time:
		t = date
	case reddit.PostTime:
		t = date.Time
	default:
		t = time.Now()
	}

	return t.Format("2006-01-02 15:04:05")
}
