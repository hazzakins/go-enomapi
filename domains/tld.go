package domains

import (
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

// Upstream documentation: https://api.enom.com/docs/get-ext-attributes
func (c *Client) GetExtAttributes(tld string) (*response.GetExtAttributes, error) {
	resp := internal.GetExtAttributesResponse{}

	cmd := c.NewCommand("GetExtAttributes")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/get-idn-codes
func (c *Client) GetIDNCodes(tld string) (*response.GetIDNCodes, error) {
	resp := internal.GetIDNCodesResponse{}

	cmd := c.NewCommand("GetIDNCodes")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

func (c *Client) GetTLDDetails(tld string) (*response.TLDDetails, error) {
	resp := internal.TLDDetailsResponse{}

	cmd := c.NewCommand("GetTLDDetails")
	cmd.AddParam("tld", tld)

	err := c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Decode()
}

func (c *Client) GetTLDList() (response.TLDList, error) {
	resp := internal.TLDListResponse{}

	cmd := c.NewCommand("GetTLDList")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
