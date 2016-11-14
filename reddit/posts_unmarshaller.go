package reddit

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type PostsUnmarshaller interface {
	UnmarshallPosts(data []byte) ([]Post, error)
}

type JsonPostUnmarshaller struct{}

func (JsonPostUnmarshaller) UnmarshallPosts(data []byte) ([]Post, error) {
	var resp response
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(resp.Data.Children))
	for _, child := range resp.Data.Children {
		post := child.RawPost.ToPost()
		posts = append(posts, post)
	}

	return posts, nil
}

type response struct {
	Data responseData
}

type responseData struct {
	Children []postData
}

type postData struct {
	RawPost rawPost `json:"data"`
}

type rawPost struct {
	Title       string
	Url         string
	Permalink   permalink
	Thumbnail   string
	Created     postTime `json:"created_utc"`
	Ups         int
	NumComments int `json:"num_comments"`
	Stickied    bool
}

func (raw rawPost) ToPost() Post {
	thumbnail := raw.Thumbnail
	if !strings.HasPrefix(raw.Thumbnail, "http") {
		thumbnail = ""
	}

	return Post{
		Title:       raw.Title,
		Url:         raw.Url,
		Permalink:   string(raw.Permalink),
		Thumbnail:   thumbnail,
		Created:     time.Time(raw.Created),
		Ups:         raw.Ups,
		NumComments: raw.NumComments,
		Stickied:    raw.Stickied,
	}
}

type Post struct {
	Title       string
	Url         string
	Permalink   string
	Thumbnail   string
	Created     time.Time
	Ups         int
	NumComments int
	Stickied    bool
}

type permalink string

func (p *permalink) UnmarshalJSON(b []byte) (err error) {
	path, err := strconv.Unquote(string(b))
	*p = permalink("https://www.reddit.com" + path)
	return
}

type postTime time.Time

func (t *postTime) UnmarshalJSON(b []byte) (err error) {
	unixTimestamp, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}
	*t = postTime(time.Unix(int64(unixTimestamp), 0).UTC())
	return
}
