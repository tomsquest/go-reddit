package http

import (
	"github.com/tomsquest/go-reddit/assets"
)

type StaticResponseHttpClient struct {
}

func (client *StaticResponseHttpClient) Get(url string) ([]byte, error) {
	return assets.Asset("assets/golang_top_week.json")
}
