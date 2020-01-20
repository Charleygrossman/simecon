package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

// Get makes an HTTP GET request and returns a JSON-encoded response.
func (c Client) Get(url string) ([]byte, error) {
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(body)
}

// Post makes a JSON-encoded HTTP POST request.
func (c Client) Post(url string, body []byte) (err error) {
	return reqWithBody(c.Client, http.MethodPost, url, body)
}

// Put makes a JSON-encoded HTTP PUT request.
func (c Client) Put(url string, body []byte) (err error) {
	return reqWithBody(c.Client, http.MethodPut, url, body)
}

// Delete makes a JSON-encoded HTTP DELETE request.
func (c Client) Delete(url string, body []byte) (err error) {
	return reqWithBody(c.Client, http.MethodDelete, url, body)
}

// reqWithBody make a JSON-encoded HTTP request with a body.
func reqWithBody(c *http.Client, method string, url string, body []byte) (err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")

	_, err = c.Do(req)
	return err
}
