package reddit

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetTopPosts(t *testing.T) {
	client := Reddit{
		HttpClient:        &fakeClient{"https://www.reddit.com/r/some-sub.json?t=week", "{}", nil},
		PostsUnmarshaller: &fakePostsUnmarshaller{[]byte("{}"), []Post{}, nil},
	}

	subreddit, err := client.GetTopPosts("some-sub")

	if assert.NoError(t, err) {
		assert.Len(t, subreddit.Posts(), 0)
	}
}

func TestGetTopPosts_GivenAHttpError(t *testing.T) {
	client := Reddit{
		HttpClient:        &fakeClient{"", "", errors.New("http error")},
		PostsUnmarshaller: &fakePostsUnmarshaller{[]byte("{}"), []Post{}, nil},
	}

	_, err := client.GetTopPosts("some-sub")

	assert.Error(t, err)
}

func TestGetTopPosts_GivenAnUnmarshallingError(t *testing.T) {
	client := Reddit{
		HttpClient:        &fakeClient{"https://www.reddit.com/r/some-sub.json?t=week", "{}", nil},
		PostsUnmarshaller: &fakePostsUnmarshaller{nil, nil, errors.New("unmarhall error")},
	}

	_, err := client.GetTopPosts("some-sub")

	assert.Error(t, err)
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
