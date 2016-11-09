package reddit

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetTopPosts(t *testing.T) {
	client := NewClient(
		WithFakeClient("https://www.reddit.com/r/some-sub.json?t=week", "{}", nil),
		WithFakePostsUnmarshaller([]byte("{}"), []Post{}, nil),
	)

	subreddit, err := client.GetTopPosts("some-sub")

	if assert.NoError(t, err) {
		assert.Len(t, subreddit.Posts(), 0)
	}
}

func TestGetTopPosts_GivenAHttpError(t *testing.T) {
	client := NewClient(
		WithFakeClient("", "", errors.New("http error")),
	)

	_, err := client.GetTopPosts("some-sub")

	assert.Error(t, err)
}

func TestGetTopPosts_GivenAnUnmarshallingError(t *testing.T) {
	client := NewClient(
		WithFakeClient("https://www.reddit.com/r/some-sub.json?t=week", "{}", nil),
		WithFakePostsUnmarshaller(nil, nil, errors.New("unmarhall error")),
	)

	_, err := client.GetTopPosts("some-sub")

	assert.Error(t, err)
}

func WithFakeClient(givenUrl, thenResponse string, thenError error) option {
	return func(reddit *Reddit) {
		reddit.client = &fakeClient{givenUrl, thenResponse, thenError}
	}
}

type fakeClient struct {
	givenUrl     string
	thenResponse string
	thenError    error
}

func (client *fakeClient) Get(url string) ([]byte, error) {
	if client.thenError != nil {
		return nil, client.thenError
	}

	if client.givenUrl == url {
		return []byte(client.thenResponse), nil
	}

	panic("Fake badly configured")
}

func WithFakePostsUnmarshaller(givenData []byte, thenPosts []Post, thenError error) option {
	return func(reddit *Reddit) {
		reddit.postsUnmarshaller = &fakePostsUnmarshaller{givenData, thenPosts, thenError}
	}
}

type fakePostsUnmarshaller struct {
	givenData []byte
	thenPosts []Post
	thenError error
}

func (fake *fakePostsUnmarshaller) UnmarshallPosts(data []byte) ([]Post, error) {
	if fake.thenError != nil {
		return nil, fake.thenError
	}

	if reflect.DeepEqual(fake.givenData, data) {
		return fake.thenPosts, fake.thenError
	}

	panic("Fake badly configured")
}
