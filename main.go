package main

import (
	"log"
	"encoding/json"
	"os"
)

func main() {
	log.Println("# Reddit")

	//subreddit := "https://www.reddit.com/r/todayilearned.json"
	//log.Println("Listing", subreddit)
	//
	//var client = &http.Client{
	//	Timeout: time.Second * 10,
	//}
	//
	//req, err := http.NewRequest("GET", subreddit, nil)
	//if err != nil {
	//	log.Fatalf("Unable to build request: %v\n", err)
	//}
	//
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.1")
	//
	//resp, err := client.Do(req)
	//
	//if err != nil {
	//	log.Fatalf("Error calling %v: %v\n", subreddit, err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	log.Fatalf(resp.Status)
	//}

	//data, err := ioutil.ReadFile("subreddit.json")
	file, err := os.Open("/home/tom/Dev/go/src/github.com/tomsquest/go-reddit/subreddit.json")
	if err != nil {
		log.Fatal(err)
	}

	//subs := new(Subreddits)
	var subs SubredditResponse
	//err = json.NewDecoder(resp.Body).Decode(&subs)
	err = json.NewDecoder(file).Decode(&subs)


	//_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatalf("Unable to read response: %v\n", err)
	}

	log.Printf("Subreddits: %#v\n", subs)

	log.Println("# Done")
}

type SubredditResponse struct {
	Data struct {
			 Children []struct {
				 Post Post    `json:"data"`
			 }
		 }
}

type Post struct {
	Title     string
	Permalink string
	Url       string
	Thumbnail string
	Created   float64 `json:"created_utc"`
}