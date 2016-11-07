package reddit

import (
	"encoding/json"
)

func UnmarshallPosts(data []byte) ([]Post, error) {
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
	Title     string
	Permalink string
	Url       string
	Thumbnail string
	Created   float64 `json:"created_utc"`
}
