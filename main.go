package main

import (
	"flag"
	"github.com/tomsquest/go-reddit/output"
	"github.com/tomsquest/go-reddit/reddit"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//go:generate go-bindata -nomemcopy -o assets/assets.go -pkg assets assets

func main() {
	subredditNameParam := flag.String("subreddit", "golang", "Subreddit to fetch")
	outputParam := flag.String("output", "console", "Select the output: console, mail")
	fakeRedditParam := flag.Bool("fakeReddit", false, "Use predefined posts instead of calling Reddit")
	flag.Parse()

	client := reddit.NewClient()

	var subreddit reddit.Subreddit
	if *fakeRedditParam {
		posts := []reddit.Post{}
		for i := 0; i < 10; i++ {
			posts = append(posts, reddit.Post{Title: "Title " + strconv.Itoa(i), Url: "http://lorempixel.com/400/400/", Thumbnail: "http://lorempixel.com/400/400/", Created: reddit.PostTime{Time: time.Now()}, Ups: rand.Int(), NumComments: rand.Int()})
		}
		subreddit = reddit.NewSubreddit(*subredditNameParam, time.Now(), posts)
	} else {
		sub, err := client.GetTopPosts(*subredditNameParam)
		subreddit = sub
		if err != nil {
			log.Fatalf("Unable to get posts of subreddit %v: %v", subredditNameParam, err)
		}
	}

	var selectedOutput output.Output
	switch *outputParam {
	case "mail":
		selectedOutput = output.SmtpOutput{}
	default:
		selectedOutput = output.Console{}
	}

	if err := selectedOutput.Out(subreddit); err != nil {
		log.Fatalf("Unable to send : %v", err)
	}
}
