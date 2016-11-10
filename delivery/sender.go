package delivery

import "github.com/tomsquest/go-reddit/reddit"

type Sender interface {
	Send(reddit.Subreddit) error
}
