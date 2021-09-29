package spotify

import (
	"io/ioutil"
	"net/http"
)

type client struct {
	httpClient *http.Client
}

func NewClient() *client {
	return &client{
		httpClient: http.DefaultClient,
	}
}

func (c *client) CreateRequest(method, uri string) (*http.Request, error) {
	req, err := http.NewRequest(method, uri, nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *client) Do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	v, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return v, nil
}
