package reddit

import (
	"github.com/hashicorp/errwrap"
	"github.com/tomsquest/go-reddit/http"
	"log"
)

type Reddit struct {
	client http.HttpClient
}

type option func(*Reddit)

func New(opts ...option) Reddit {
	reddit := Reddit{
		client: http.NewHttpClient("Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1"),
	}

	for _, opt := range opts {
		opt(&reddit)
	}

	return reddit
}

func (reddit *Reddit) GetTopPosts(subreddit string) ([]Post, error) {

	log.Println("Getting top posts")

	url := "https://www.reddit.com/r/" + subreddit

	data, err := reddit.client.Get(url)
	if err != nil {
		return nil, errwrap.Wrapf("Unable to get posts of subreddit '"+subreddit+"': {{err}}", err)
	}

	posts, err := UnmarshallPosts(data)
	if err != nil {
		return nil, errwrap.Wrapf("Unable to unmarshall posts: {{err}}", err)
	}

	log.Printf("Posts: %#v\n", posts)

	return posts, nil
}
