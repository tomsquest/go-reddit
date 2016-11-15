package output

import (
	"fmt"
	"github.com/tomsquest/go-reddit/reddit"
)

type Console struct {
}

func (Console) Out(subreddit reddit.Subreddit) error {
	fmt.Println("Subreddit", subreddit.Name, "crawled at ", subreddit.CrawlDate.String())

	for idx, post := range subreddit.Posts() {
		fmt.Printf("  -> Post %2d - %v\n", idx, substring(post.Title, 40))
	}

	return nil
}

func substring(s string, size int) string {
	if len(s) < size {
		return s
	}

	return s[:size] + "..."
}
