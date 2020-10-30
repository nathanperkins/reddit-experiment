package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	userAgent = "AuthTest:v0.0 (by /u/wafflezone)"
)

type Client struct {
	token string
}

func New(username, password, clientID, clientSecret string) (*Client, error) {
	const endpoint = "https://www.reddit.com/api/v1/access_token"

	// Prepare request.
	data := url.Values{
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
	}
	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("User-Agent", userAgent)

	// Do request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode Token from resp.Body.
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respJSON struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(b, &respJSON); err != nil {
		return nil, err
	}
	token := respJSON.AccessToken
	if token == "" {
		return nil, fmt.Errorf("failed to decode token from response body: %s", b)
	}
	return &Client{token: token}, nil
}

func (c *Client) Get(endpoint string, data interface{}) error {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", c.token))
	req.Header.Add("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, data); err != nil {
		return err
	}

	return nil
}
