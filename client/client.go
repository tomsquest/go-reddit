package client

import (
	"fmt"
	"github.com/hashicorp/errwrap"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	realClient *http.Client
	userAgent  string
}

func New(userAgent string) Client {
	var realClient = &http.Client{
		Timeout: time.Second * 10,
	}

	return Client{
		realClient: realClient,
		userAgent:  userAgent,
	}
}

func (c *Client) Get(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errwrap.Wrapf("Error building request: {{err}}", err)
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.realClient.Do(req)
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
