package main

import (
	"flag"
	"fmt"
	"github.com/tomsquest/go-reddit/reddit"
	"log"
)

func main() {
	subredditName := flag.String("subreddit", "golang", "Subreddit to fetch")
	flag.Parse()

	client := reddit.NewClient()

	subreddit, err := client.GetTopPosts(*subredditName)
	if err != nil {
		log.Fatalf("Unable to get posts of subreddit %v: %v", subredditName, err)
	}

	for idx, post := range subreddit.Posts() {
		fmt.Printf("Post %2d - %v\n", idx, post.Title)
	}
}
