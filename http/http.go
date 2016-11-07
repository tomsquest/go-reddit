package http

import (
	"fmt"
	"github.com/hashicorp/errwrap"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	realClient *http.Client
	userAgent  string
}

func NewHttpClient(userAgent string) HttpClient {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	return HttpClient{
		realClient: client,
		userAgent:  userAgent,
	}
}

func (client *HttpClient) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errwrap.Wrapf("Error building request: {{err}}", err)
	}

	req.Header.Set("User-Agent", client.userAgent)

	resp, err := client.realClient.Do(req)
	if err != nil {
		return nil, errwrap.Wrapf("Error calling "+url+": {{err}}\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrapf("Error reading response body: {{err}}", err)
	}

	return b, nil
}
