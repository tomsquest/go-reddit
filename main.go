package main

import (
	"flag"
	"github.com/tomsquest/go-reddit/http"
	"github.com/tomsquest/go-reddit/output"
	"github.com/tomsquest/go-reddit/reddit"
	"log"
)

//go:generate go-bindata -nomemcopy -o assets/assets.go -pkg assets assets

func main() {
	subredditNameParam := flag.String("subreddit", "golang", "Subreddit to fetch")
	outputParam := flag.String("output", "console", "Select the output: console, mail")
	fakeRedditParam := flag.Bool("fakeReddit", false, "Use predefined posts instead of calling Reddit")
	flag.Parse()

	var httpClient http.HttpClient
	if *fakeRedditParam {
		httpClient = &http.StaticResponseHttpClient{}
	} else {
		httpClient = http.NewHttpClient("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0")
	}

	client := reddit.Reddit{
		HttpClient:        httpClient,
		PostsUnmarshaller: reddit.JsonPostUnmarshaller{},
	}

	subreddit, err := client.GetTopPosts(*subredditNameParam)
	if err != nil {
		log.Fatalf("Unable to get posts of subreddit %v: %v", subredditNameParam, err)
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
