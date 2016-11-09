package reddit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUnmarshallOnePost(t *testing.T) {
	data := []byte(`
		{
			"data": {
				"children": [
					{
						"data": {
							"title": "title1",
							"thumbnail": "http://thumbnail1.jpg",
							"permalink": "/r/perma1",
							"url": "https://url1",
							"created_utc": 1451703845.0
						}
					}
				]
			}
		}
	`)

	posts, err := UnmarshallPosts(data)

	if assert.NoError(t, err) {
		if assert.Len(t, posts, 1) {
			post := posts[0]
			assert.Equal(t, "title1", post.Title)
			assert.Equal(t, "http://thumbnail1.jpg", post.Thumbnail)
			assert.Equal(t, "/r/perma1", post.Permalink)
			assert.Equal(t, "https://url1", post.Url)
			assert.Equal(t, time.Date(2016, 1, 2, 3, 4, 5, 0, time.UTC), post.Created.Time)
		}
	}
}

func TestUnmarshallTwoPosts(t *testing.T) {
	data := []byte(`
		{
			"data": {
				"children": [
					{
						"data": {
							"thumbnail": "http://thumbnail1.jpg",
							"permalink": "/r/perma1",
							"url": "https://url1",
							"title": "title1",
							"created_utc": 1000000001.0
						}
					},
					{
						"data": {
							"thumbnail": "http://thumbnail2.jpg",
							"permalink": "/r/perma2",
							"url": "https://url2",
							"title": "title2",
							"created_utc": 1000000002.0
						}
					}
				]
			}
		}
	`)

	posts, err := UnmarshallPosts(data)

	if assert.NoError(t, err) {
		if assert.Len(t, posts, 2) {
			post := posts[1]
			assert.Equal(t, "title2", post.Title)
		}
	}
}

func TestUnmarshallNoPost(t *testing.T) {
	data := []byte(`{}`)

	posts, err := UnmarshallPosts(data)

	if assert.NoError(t, err) {
		assert.Len(t, posts, 0)
	}
}
