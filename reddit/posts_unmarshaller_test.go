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
							"url": "https://url1",
							"permalink": "/r/sub1/comments/1",
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

	posts, err := JsonPostUnmarshaller{}.UnmarshallPosts(data)

	if assert.NoError(t, err) {
		if assert.Len(t, posts, 1) {
			post := posts[0]
			assert.Equal(t, "title1", post.Title)
			assert.Equal(t, "http://thumbnail1.jpg", post.Thumbnail)
			assert.Equal(t, "https://url1", post.Url)
			assert.Equal(t, "https://www.reddit.com/r/sub1/comments/1", post.Permalink)
			assert.Equal(t, time.Date(2016, 1, 2, 3, 4, 5, 0, time.UTC), post.Created)
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

	posts, err := JsonPostUnmarshaller{}.UnmarshallPosts(data)

	if assert.NoError(t, err) {
		if assert.Len(t, posts, 2) {
			post := posts[1]
			assert.Equal(t, "title2", post.Title)
		}
	}
}

func TestUnmarshallNoPost(t *testing.T) {
	data := []byte(`{}`)

	posts, err := JsonPostUnmarshaller{}.UnmarshallPosts(data)

	if assert.NoError(t, err) {
		assert.Len(t, posts, 0)
	}
}

func TestUnmarshallWithNoThumbnail(t *testing.T) {
	data := []byte(`
		{
			"data": {
				"children": [
					{
						"data": {
							"thumbnail": "self"
						}
					}
				]
			}
		}
	`)

	posts, err := JsonPostUnmarshaller{}.UnmarshallPosts(data)

	if assert.NoError(t, err) {
		assert.Equal(t, "", posts[0].Thumbnail)
	}
}
