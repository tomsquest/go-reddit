package reddit

import (
	"github.com/hashicorp/errwrap"
	"github.com/mgutz/logxi/v1"
	"github.com/tomsquest/go-reddit/http"
)

var logger log.Logger = log.New("reddit")

type Reddit struct {
	client http.HttpClient
}

type option func(*Reddit)

func New(opts ...option) Reddit {
	reddit := Reddit{
		client: http.NewHttpClient("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0"),
	}

	for _, opt := range opts {
		opt(&reddit)
	}

	return reddit
}

func (reddit *Reddit) GetTopPosts(subreddit string) ([]Post, error) {

	logger.Info("Getting top posts")

	url := "https://www.reddit.com/r/" + subreddit + ".json" + "?t=week"

	data, err := reddit.client.Get(url)
	if err != nil {
		return nil, errwrap.Wrapf("Unable to get posts of subreddit '"+subreddit+"': {{err}}", err)
	}

	posts, err := UnmarshallPosts(data)
	if err != nil {
		return nil, errwrap.Wrapf("Unable to unmarshall posts: {{err}}", err)
	}

	if logger.IsDebug() {
		logger.Debug("Posts", len(posts))
		for i, post := range posts {
			logger.Debug("Post", "index", i, "post", post)
		}
	}

	return posts, nil
}
