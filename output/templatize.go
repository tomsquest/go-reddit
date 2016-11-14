package output

import (
	"bytes"
	"github.com/tomsquest/go-reddit/assets"
	"github.com/tomsquest/go-reddit/reddit"
	"html/template"
	"strings"
	"time"
)

func templatize(subreddit reddit.Subreddit) (string, error) {
	buf := new(bytes.Buffer)
	data, err := assets.Asset("assets/subreddit.html")
	if err != nil {
		return "", err
	}

	funcs := template.FuncMap{
		"title":   strings.Title,
		"toUpper": strings.ToUpper,
		"toDate":  toDate,
	}

	tmpl, err := template.New("html").Funcs(funcs).Parse(string(data))
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(buf, subreddit); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func toDate(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
