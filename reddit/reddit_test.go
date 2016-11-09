package reddit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTopPosts(t *testing.T) {
	client := New(WithFakeClient("https://www.reddit.com/r/some-sub.json?t=week", "{}"))

	subreddit, err := client.GetTopPosts("some-sub")

	if assert.NoError(t, err) {
		assert.Len(t, subreddit.Posts(), 0)
	}
}

func WithFakeClient(answerUrl, answerResponse string) option {
	return func(reddit *Reddit) {
		reddit.client = &fakeClient{answerUrl, answerResponse}
	}
}

type fakeClient struct {
	answerUrl      string
	answerResponse string
}

func (client *fakeClient) Get(url string) ([]byte, error) {
	if client.answerUrl == url {
		return []byte(client.answerResponse), nil
	}

	return nil, fmt.Errorf("url unknown %v", url)
}
