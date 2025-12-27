package domains

import (
	"fmt"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

// Upstream documentation: https://api.enom.com/docs/check
func (c *Client) Check(domain enomapi.Domain) (*response.DomainCheck, error) {
	resp := internal.DomainCheckResponse{}

	cmd := c.NewCommand("Check")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("Version", "2")
	cmd.AddParam("IncludePrice", "true")
	cmd.AddParam("IncludeProperties", "true")

	err := c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}

	if resp.RRPCode > 300 {
		return nil, fmt.Errorf("RRP error %d: %s", resp.RRPCode, resp.RRPText)
	}

	return resp.Decode()
}
