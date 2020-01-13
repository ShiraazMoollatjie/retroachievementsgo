package retroachievementsgo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client makes all the API calls to dev.to.
type Client struct {
	apiKey   string
	baseURL  string
	http     *http.Client
	username string
}

// Arguments are used for passing query parameters to the dev.to api.
type Arguments map[string]string

// Defaults returns an empty map of arguments.
func Defaults() Arguments {
	return make(map[string]string)
}

func (a Arguments) toQueryParams() url.Values {
	res := make(url.Values)
	for k, v := range a {
		res.Add(k, v)
	}
	return res
}

// Option allows the client to be configured with different options.
type Option func(*Client)

func withBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithApiKey sets the retroachievements api key to use for this client.
func WithApiKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithUsername sets the retroachievements username to use for this client.
func WithUsername(username string) Option {
	return func(c *Client) {
		c.username = username
	}
}

// NewClient creates a dev.to client with the provided options.
func NewClient(opts ...Option) *Client {
	res := &Client{
		baseURL: "https://ra.hfc-essentials.com",
		http:    &http.Client{},
	}
	for _, o := range opts {
		o(res)
	}

	return res
}

// get returns an error if the http client cannot perform a HTTP GET for the provided URL.
func (c *Client) get(url string, target interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error from retroachievements api")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &target)
}

func (c *Client) buildReqQueryParams() url.Values {
	v := make(url.Values)
	v.Set("mode", "json")
	v.Set("user", c.username)
	v.Set("key", c.apiKey)

	return v
}
