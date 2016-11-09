package reddit

import "time"

type Subreddit struct {
	Name      string
	CrawlDate time.Time
	posts     []Post
}

func NewSubreddit(name string, crawlDate time.Time, posts []Post) Subreddit {
	return Subreddit{name, crawlDate, posts}
}

func (sub Subreddit) Posts() []Post {
	filteredPosts := sub.posts[:0] // filter without allocating
	for _, post := range sub.posts {
		if !post.Stickied {
			filteredPosts = append(filteredPosts, post)
		}
	}
	return filteredPosts
}
