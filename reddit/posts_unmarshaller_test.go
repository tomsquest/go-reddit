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
							"created_utc": 1451703845.0,
							"ups": 123,
							"num_comments": 345,
							"stickied": true
						}
					}
				]
			}
		}
	`)

	posts, err := postsUnmarshaller{}.UnmarshallPosts(data)

	if assert.NoError(t, err) {
		if assert.Len(t, posts, 1) {
			post := posts[0]
			assert.Equal(t, "title1", post.Title)
			assert.Equal(t, "http://thumbnail1.jpg", post.Thumbnail)
			assert.Equal(t, "/r/perma1", post.Permalink)
			assert.Equal(t, "https://url1", post.Url)
			assert.Equal(t, time.Date(2016, 1, 2, 3, 4, 5, 0, time.UTC), post.Created.Time)
			assert.Equal(t, 123, post.Ups)
			assert.Equal(t, 345, post.NumComments)
			assert.Equal(t, true, post.Stickied)
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
							"title": "title1"
						}
					},
					{
						"data": {
							"title": "title2"
						}
					}
				]
			}
		}
	`)

	posts, err := postsUnmarshaller{}.UnmarshallPosts(data)

	if assert.NoError(t, err) {
		if assert.Len(t, posts, 2) {
			post := posts[1]
			assert.Equal(t, "title2", post.Title)
		}
	}
}

func TestUnmarshallNoPost(t *testing.T) {
	data := []byte(`{}`)

	posts, err := postsUnmarshaller{}.UnmarshallPosts(data)

	if assert.NoError(t, err) {
		assert.Len(t, posts, 0)
	}
}
