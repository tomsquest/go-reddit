package http

import (
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/mgutz/logxi/v1"
	"github.com/tomsquest/go-reddit/config"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient interface {
	Get(string) ([]byte, error)
}

type httpClient struct {
	realClient *http.Client
	userAgent  string
}

func NewHttpClient(cfg config.Config) HttpClient {
	var client = &http.Client{
		Timeout: time.Second * 20,
	}

	return &httpClient{
		realClient: client,
		userAgent:  cfg.UserAgent,
	}
}

func (client *httpClient) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errwrap.Wrapf("Error building request: {{err}}", err)
	}

	req.Header.Set("User-Agent", client.userAgent)

	log.Info("Get", "url", url)
	resp, err := client.realClient.Do(req)
	if err != nil {
		return nil, errwrap.Wrapf("Error calling "+url+": {{err}}\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	log.Debug("Go response", "status", resp.Status)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrapf("Error reading response body: {{err}}", err)
	}

	return b, nil
}
