package main

import (
	"github.com/tomsquest/go-reddit/config"
	"github.com/tomsquest/go-reddit/http"
	"github.com/tomsquest/go-reddit/output"
	"github.com/tomsquest/go-reddit/reddit"
	"log"
)

//go:generate go-bindata -nomemcopy -o assets/assets.go -pkg assets assets

func main() {
	cfg, err := config.Read()

	var httpClient http.HttpClient
	if cfg.FakeReddit {
		httpClient = &http.StaticResponseHttpClient{}
	} else {
		httpClient = http.NewHttpClient(cfg)
	}

	client := reddit.Reddit{
		HttpClient:        httpClient,
		PostsUnmarshaller: reddit.JsonPostUnmarshaller{},
	}

	subredditName := cfg.Subreddits[0]
	subreddit, err := client.GetTopPosts(subredditName)
	if err != nil {
		log.Fatalf("Unable to get posts of subreddit %v: %v", subredditName, err)
	}

	var selectedOutput output.Output
	switch cfg.Output {
	case "mail":
		selectedOutput = output.NewSmtpOutput(cfg)
	default:
		selectedOutput = output.Console{}
	}

	if err := selectedOutput.Out(subreddit); err != nil {
		log.Fatalf("Unable to send : %v", err)
	}
}
