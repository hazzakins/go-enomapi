package domains

import (
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type PEGetDomainPricingRequest struct {
	UseQtyEngine *bool
	Years        *int
}

// Upstream documentation: https://api.enom.com/docs/pe_getdomainpricing
func (c *Client) PEGetDomainPricing(req PEGetDomainPricingRequest) (*response.PEGetDomainPricing, error) {
	resp := internal.PEGetDomainPricingResponse{}

	cmd := c.NewCommand("PE_GetDomainPricing")
	addBoolIntParam(cmd, "UseQtyEngine", req.UseQtyEngine)
	addIntParam(cmd, "Years", req.Years)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
