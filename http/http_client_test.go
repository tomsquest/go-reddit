package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tomsquest/go-reddit/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet_existingUrl(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `a response`)
	}))
	defer fakeServer.Close()

	client := NewHttpClient(config.Config{})

	resp, err := client.Get(fakeServer.URL + "/some-page")

	if assert.NoError(t, err) {
		assert.EqualValues(t, "a response", string(resp))
	}
}

func TestGet_AddUserAgent(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Header.Get("User-Agent"), "test user agent")
	}))
	defer fakeServer.Close()

	client := NewHttpClient(config.Config{UserAgent: "test user agent"})

	client.Get(fakeServer.URL + "/some-page")
}

func TestGet_given404(t *testing.T) {
	fakeServer := httptest.NewServer(http.NotFoundHandler())
	defer fakeServer.Close()

	client := NewHttpClient(config.Config{})

	resp, err := client.Get(fakeServer.URL + "/unknown")

	assert.Nil(t, resp)
	assert.EqualError(t, err, "404 Not Found")
}
