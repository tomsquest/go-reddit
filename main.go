package main

import (
	"flag"
	"fmt"
	"github.com/tomsquest/go-reddit/delivery"
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

	fmt.Printf("Subreddit '%v', crawled at %v: %v posts fetched\n", subreddit.Name, subreddit.CrawlDate, len(subreddit.Posts()))

	sender := delivery.NewSmtpSender()
	if err = sender.Send(subreddit); err != nil {
		log.Fatalf("Unable to send email: %v", err)
	}

	fmt.Println("Mail sent")
}
