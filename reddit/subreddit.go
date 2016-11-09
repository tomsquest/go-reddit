package reddit

type Subreddit struct {
	Name  string
	posts []Post
}

func NewSubreddit(name string, posts []Post) Subreddit {
	return Subreddit{name, posts}
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
