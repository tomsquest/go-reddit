package reddit

import (
	"encoding/json"
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

type SubredditResponse struct {
	Data struct {
		Children []struct {
			Post Post `json:"data"`
		}
	}
}

type Post struct {
	Title     string
	Permalink string
	Url       string
	Thumbnail string
	Created   float64 `json:"created_utc"`
}

func (reddit *Reddit) GetTopPosts(subreddit string) ([]Post, error) {

	log.Println("Getting top posts")

	url := "https://www.reddit.com/r/" + subreddit

	data, err := reddit.client.Get(url)
	if err != nil {
		return nil, errwrap.Wrapf("Unable to get posts of subreddit '"+subreddit+"': {{err}}", err)
	}

	var subs SubredditResponse
	err = json.Unmarshal(data, &subs)
	if err != nil {
		log.Fatalf("Unable to read response: %v\n", err)
	}

	log.Printf("Subreddits: %#v\n", subs)

	log.Println("# Done")

	return nil, nil
}
