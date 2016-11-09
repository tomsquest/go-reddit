package reddit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterOutStickies(t *testing.T) {
	notSticky := Post{Title: "not sticky"}
	stickied := Post{Title: "stickied", Stickied: true}

	subreddit := NewSubreddit("a-sub", []Post{notSticky, stickied})

	assert.Contains(t, subreddit.Posts(), notSticky)
	assert.NotContains(t, subreddit.Posts(), stickied)
}
