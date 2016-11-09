package reddit

import (
	"encoding/json"
	"strconv"
	"time"
)

type PostsUnmarshaller interface {
	UnmarshallPosts(data []byte) ([]Post, error)
}

type postsUnmarshaller struct{}

func (postsUnmarshaller) UnmarshallPosts(data []byte) ([]Post, error) {
	var resp Response
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, 25)
	for _, child := range resp.Data.Children {
		posts = append(posts, child.Post)
	}

	return posts, nil
}

type Response struct {
	Data ResponseData
}

type ResponseData struct {
	Children []PostData
}

type PostData struct {
	Post Post `json:"data"`
}

type Post struct {
	Title       string
	Permalink   string
	Url         string
	Thumbnail   string
	Created     PostTime `json:"created_utc"`
	Ups         int
	NumComments int `json:"num_comments"`
	Stickied    bool
}

type PostTime struct {
	time.Time
}

func (t *PostTime) UnmarshalJSON(b []byte) (err error) {
	unixTimestamp, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(int64(unixTimestamp), 0).UTC()
	return
}
