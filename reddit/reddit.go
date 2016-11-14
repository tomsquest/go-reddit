package reddit

import (
	"github.com/hashicorp/errwrap"
	"github.com/mgutz/logxi/v1"
	"github.com/tomsquest/go-reddit/http"
	"time"
)

var logger log.Logger = log.New("reddit")

type Reddit struct {
	HttpClient        http.HttpClient
	PostsUnmarshaller PostsUnmarshaller
}

func (reddit *Reddit) GetTopPosts(subredditName string) (subreddit Subreddit, err error) {
	url := "https://www.reddit.com/r/" + subredditName + ".json" + "?t=week"
	logger.Info("Getting top posts", "subreddit", subredditName, "url", url)

	data, err := reddit.HttpClient.Get(url)
	if err != nil {
		return subreddit, errwrap.Wrapf("Unable to get posts of subreddit '"+subredditName+"': {{err}}", err)
	}

	posts, err := reddit.PostsUnmarshaller.UnmarshallPosts(data)
	if err != nil {
		return subreddit, errwrap.Wrapf("Unable to unmarshall posts: {{err}}", err)
	}

	if logger.IsDebug() {
		logger.Debug("Posts", len(posts))
		for i, post := range posts {
			logger.Debug("Post", "index", i, "post", post)
		}
	}

	subreddit = NewSubreddit(subredditName, time.Now(), posts)
	return
}
