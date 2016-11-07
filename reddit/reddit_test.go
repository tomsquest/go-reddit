package reddit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTopPosts_GivenSomePosts(t *testing.T) {
	//fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, `
	//	{
	//		"data": {
	//			"children": [
	//				{
	//					"data": {
	//						"thumbnail": "http://thumbnail1.jpg",
	//						"permalink": "/r/perma1",
	//						"url": "https://url1",
	//						"title": "title1",
	//						"created_utc": 1000000001.0,
	//					}
	//				},
	//				{
	//					"data": {
	//						"thumbnail": "http://thumbnail2.jpg",
	//						"permalink": "/r/perma2",
	//						"url": "https://url2",
	//						"title": "title2",
	//						"created_utc": 1000000002.0,
	//					}
	//				}
	//			]
	//		}
	//	}
	//	`)
	//}))
	//defer fakeServer.Close()

	client := New(FakeClientWithResponse("a response"))

	posts, _ := client.GetTopPosts("some-sub")

	assert.Len(t, posts, 2)
}

func FakeClientWithResponse(resp string) option {
	return func(reddit *Reddit) {
		//reddit.Client :=
	}
}

type fakeClient struct{}

func (c *fakeClient) Get(url string) ([]byte, error) {
	return []byte("a response"), nil
}
