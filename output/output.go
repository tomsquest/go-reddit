package output

import "github.com/tomsquest/go-reddit/reddit"

type Output interface {
	Out(reddit.Subreddit) error
}
