package main

import (
	"fmt"
	"github.com/tomsquest/go-reddit/reddit"
	"log"
)

func main() {
	subreddit := "gifs"
	client := reddit.New()

	posts, err := client.GetTopPosts(subreddit)
	if err != nil {
		log.Fatalf("Unable to get posts: %v", err)
	}

	for idx, post := range posts {
		fmt.Printf("Post %2d - %v\n", idx, post.Title)
	}
}
