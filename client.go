/*
   client.go
*/
package idex

import (
	"fmt"
	"net/http"
    "encoding/json"
)

type Client struct {
	address    string
	privKey    string
	httpClient *http.Client
}

func NewClient(address, key string) (c *Client) {
	client := &Client{
		address:    address,
		privKey:    key,
		httpClient: &http.Client{},
	}
	return client
}

func (c *Client) do(method, resource, payload string, auth bool, result interface{}) (resp *http.Response, err error) {

	fullUrl := fmt.Sprintf("%s/%s", BaseUrl, resource)

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return
	}

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp != nil {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(result)
	}
	return
}
