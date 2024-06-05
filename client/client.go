package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL           string
	TokenExpiry       time.Time
	HttpClient        *http.Client
	APIKey            string
	APIToken          string
	DurantionToExpiry time.Duration
}

func NewClient(baseUrl string, apiKey string, durantionToExpiry time.Duration) (*Client, error) {
	client := &Client{
		BaseURL:           baseUrl,
		HttpClient:        &http.Client{},
		APIKey:            apiKey,
		DurantionToExpiry: durantionToExpiry,
	}
	if err := client.login(); err != nil {
		return nil, err
	}
	return client, nil

}

type LoginResponse struct {
	Status string    `json:"status"`
	Data   DataToken `json:"data"`
}

type DataToken struct {
	Token string `json:"token"`
}

func (c *Client) login() error {
	url := fmt.Sprintf("%s/login", c.BaseURL)

	credentials := map[string]string{
		"apikey": c.APIKey,
		"pin":    "IPQZOCMN",
	}

	reqBody, err := json.Marshal(credentials)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: %s", resp.Status)
	}

	var loginResponse LoginResponse

	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return err
	}

	c.APIToken = loginResponse.Data.Token
	c.TokenExpiry = time.Now().Add(c.DurantionToExpiry)

	return nil
}

func (c *Client) checkToken() error {
	if time.Now().After(c.TokenExpiry) {
		return c.login()
	}
	return nil
}

// define a function to create a request
// why do that? this way we can have a request fully set up with authoriztion and header to be called by the http client
func (c *Client) NewRequest(method string, endpoint string, body map[string]any) (*http.Request, error) {
	if err := c.checkToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	var buffer *bytes.Buffer
	if body == nil {
		buffer = bytes.NewBuffer(nil)
	} else {
		mbody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(mbody)
	}

	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.APIToken))

	return req, nil
}

func (c *Client) Do(req *http.Request, response interface{}) error {
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error found while calling api response code: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}
