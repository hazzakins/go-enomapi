package enomapi

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hazzakins/go-enomapi/internal"
)

type Client struct {
	resellerURL *url.URL
	resellerID  string
	apikey      string
}

func NewClient(resellerURL string, resellerID string, apikey string) (*Client, error) {
	baseURL, err := url.Parse(resellerURL + "/interface.asp")
	if err != nil {
		return nil, fmt.Errorf("failed to parse reseller URL: %w", err)
	}

	return &Client{
		resellerURL: baseURL,
		resellerID:  resellerID,
		apikey:      apikey,
	}, nil
}

func (c *Client) url(cmd *Command) string {
	c.resellerURL.RawQuery = cmd.toParams()
	return c.resellerURL.String()
}

func (c *Client) Execute(cmd *Command, resp internal.BaseResponse) error {
	httpResp, err := http.Get(c.url(cmd))
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Error: %s, %d", http.StatusText(httpResp.StatusCode), httpResp.StatusCode)
	}

	if err := xml.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if err := resp.Err(); err != nil {
		return fmt.Errorf("API Error: %w", err)
	}

	return nil
}
