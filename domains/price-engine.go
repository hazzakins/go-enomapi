package domains

import (
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

// Upstream documentation: https://api.enom.com/docs/pe-get-tld-id
func (c *Client) PEGetTLDID(tld string) (*response.PEGetTLDID, error) {
	resp := internal.PEGetTLDIDResponse{}

	cmd := c.NewCommand("PE_GetTLDID")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
