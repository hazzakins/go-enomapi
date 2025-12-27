package domains

import (
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type PEGetDomainPricingRequest struct {
	UseQtyEngine *bool
	Years        *int
}

type PESetPricingRequest struct {
	Parameters map[string]string
}

type RefillAccountRequest struct {
	AdditionalParams map[string]string
}

type SetResellerServicesPricingRequest struct {
	Parameters map[string]string
}

type SetResellerTLDPricingRequest struct {
	Parameters map[string]string
}

type UpdateAccountPricingRequest struct {
	Parameters map[string]string
}

// Upstream documentation: https://api.enom.com/docs/pe-get-tld-id
func (c *Client) PEGetTLDID(tld string) (*response.PEGetTLDID, error) {
	resp := internal.PEGetTLDIDResponse{}

	cmd := c.NewCommand("PE_GetTLDID")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
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

// Upstream documentation: https://api.enom.com/docs/pe-set-pricing
func (c *Client) PESetPricing(req PESetPricingRequest) (*response.PESetPricing, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("PE_SetPricing")
	addParams(cmd, req.Parameters)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/refill-account
func (c *Client) RefillAccount(req RefillAccountRequest) (*response.RefillAccount, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("RefillAccount")
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/setresellerservicespricing
func (c *Client) SetResellerServicesPricing(req SetResellerServicesPricingRequest) (*response.SetResellerServicesPricing, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetResellerServicesPricing")
	addParams(cmd, req.Parameters)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/setresellertldpricing
func (c *Client) SetResellerTLDPricing(req SetResellerTLDPricingRequest) (*response.SetResellerTLDPricing, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetResellerTLDPricing")
	addParams(cmd, req.Parameters)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/updateaccountpricing
func (c *Client) UpdateAccountPricing(req UpdateAccountPricingRequest) (*response.UpdateAccountPricing, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("UpdateAccountPricing")
	addParams(cmd, req.Parameters)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}
