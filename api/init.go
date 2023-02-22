package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	HTTPClient    *http.Client
	Host          string
	Username      string
	Password      string
	Authorization string
}

func NewClient(host string, username string, password string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
		Host:          host,
		Username:      username,
		Password:      password,
		Authorization: "",
	}
	return &c, nil
}

func (c *Client) doRequest(method string, url string, body []byte, headers map[string]string) ([]byte, int, string, map[string][]string, error) {
	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, 500, "500 Internal Server Error", nil, err
	}

	if c.Authorization == "" {
		c.Authorization = c.genBasicAuthToken()
	}

	req.Header.Add("Authorization", c.Authorization)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "PCT")
	req.Header.Add("Content-Type", "application/json")

	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 500, "500 Internal Server Error", nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 500, "500 Internal Server Error", nil, err
	}

	defer res.Body.Close()
	return b, res.StatusCode, res.Status, res.Header, nil
}

func (c *Client) genBasicAuthToken() string {
	if c.Username == "" || c.Password == "" {
		return ""
	}
	str := fmt.Sprintf("%s:%s", c.Username, c.Password)
	encStr := base64.StdEncoding.EncodeToString([]byte(str))
	return fmt.Sprintf("Basic %s", encStr)
}
