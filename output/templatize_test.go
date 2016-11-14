package output

import (
	"github.com/stretchr/testify/assert"
	"github.com/tomsquest/go-reddit/reddit"
	"testing"
	"time"
)

func TestTemplatize_singlePost(t *testing.T) {
	subreddit := reddit.NewSubreddit("SubredditName", time.Now(), []reddit.Post{{
		Title:       "Title 1",
		Url:         "http://url1",
		Permalink:   "http://perma1",
		Thumbnail:   "http://thumb1",
		Created:     reddit.PostTime{Time: time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)},
		Ups:         100,
		NumComments: 101,
	}})

	html, err := templatize(subreddit)

	if assert.NoError(t, err) {
		assert.Contains(t, html, "SUBREDDITNAME")
		assert.Contains(t, html, "Title 1")
		assert.Contains(t, html, "src=\"http://thumb1\"")
		assert.Contains(t, html, "href=\"http://url1\"")
		assert.Contains(t, html, "href=\"http://perma1\"")
		assert.Contains(t, html, "2000-01-02 03:04:05")
	}
}

func TestTemplatize_manyPost(t *testing.T) {
	subreddit := reddit.NewSubreddit("SubredditName", time.Now(), []reddit.Post{{
		Title: "Title 1",
	}, {
		Title: "Title 2",
	}})

	html, err := templatize(subreddit)

	if assert.NoError(t, err) {
		assert.Contains(t, html, "Title 1")
		assert.Contains(t, html, "Title 2")
	}
}
