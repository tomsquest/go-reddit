package reddit

import (
	"github.com/hashicorp/errwrap"
	"github.com/mgutz/logxi/v1"
	"github.com/tomsquest/go-reddit/http"
)

var logger log.Logger = log.New("reddit")

type Reddit struct {
	client            http.HttpClient
	postsUnmarshaller PostsUnmarshaller
}

type option func(*Reddit)

func NewClient(opts ...option) Reddit {
	reddit := Reddit{
		client:            http.NewHttpClient("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0"),
		postsUnmarshaller: postsUnmarshaller{},
	}

	for _, opt := range opts {
		opt(&reddit)
	}

	return reddit
}

func (reddit *Reddit) GetTopPosts(subredditName string) (subreddit Subreddit, err error) {
	logger.Info("Getting top posts", "subreddit", subredditName)

	url := "https://www.reddit.com/r/" + subredditName + ".json" + "?t=week"

	data, err := reddit.client.Get(url)
	if err != nil {
		return subreddit, errwrap.Wrapf("Unable to get posts of subreddit '"+subredditName+"': {{err}}", err)
	}

	posts, err := reddit.postsUnmarshaller.UnmarshallPosts(data)
	if err != nil {
		return subreddit, errwrap.Wrapf("Unable to unmarshall posts: {{err}}", err)
	}

	if logger.IsDebug() {
		logger.Debug("Posts", len(posts))
		for i, post := range posts {
			logger.Debug("Post", "index", i, "post", post)
		}
	}

	subreddit = NewSubreddit(subredditName, posts)
	return
}
