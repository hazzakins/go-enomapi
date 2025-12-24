package domains

import (
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)


// ParameterSet represents a collection of optional parameters to include with a
// command. Only entries with non-empty values are transmitted.
type ParameterSet map[string]string

func addParameters(cmd *enomapi.Command, params ParameterSet) {
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}
}
// Upstream documentation: https://api.enom.com/docs/pe-get-tld-id
func (c *Client) PEGetTLDID(tld string) (*response.PEGetTLDID, error) {
	resp := internal.PEGetTLDIDResponse{}

	cmd := c.NewCommand("PE_GetTLDID")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/pe-set-pricing
func (c *Client) PESetPricing(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("PE_SetPricing")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/pushdomain
func (c *Client) PushDomain(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("PushDomain")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/refill-account
func (c *Client) RefillAccount(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("RefillAccount")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/setresellerservicespricing
func (c *Client) SetResellerServicesPricing(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetResellerServicesPricing")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/setresellertldpricing
func (c *Client) SetResellerTLDPricing(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetResellerTLDPricing")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/synchauthinfo
func (c *Client) SynchAuthInfo(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SynchAuthInfo")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-cancelorder
func (c *Client) TPCancelOrder(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_CancelOrder")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-createorder
func (c *Client) TPCreateOrder(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_CreateOrder")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getdetailsbydomain
func (c *Client) TPGetDetailsByDomain(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetDetailsByDomain")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorder
func (c *Client) TPGetOrder(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrder")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderdetail
func (c *Client) TPGetOrderDetail(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderDetail")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getordersbydomain
func (c *Client) TPGetOrdersByDomain(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrdersByDomain")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderreview
func (c *Client) TPGetOrderReview(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderReview")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderstatuses
func (c *Client) TPGetOrderStatuses(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderStatuses")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-get-tld-info
func (c *Client) TPGetTLDInfo(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetTLDInfo")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-resendemail
func (c *Client) TPResendEmail(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_ResendEmail")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-resubmit-locked
func (c *Client) TPResubmitLocked(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_ResubmitLocked")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-submitorder
func (c *Client) TPSubmitOrder(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_SubmitOrder")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-updateorderdetail
func (c *Client) TPUpdateOrderDetail(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_UpdateOrderDetail")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/updateaccountpricing
func (c *Client) UpdateAccountPricing(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("UpdateAccountPricing")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/updatepushlist
func (c *Client) UpdatePushList(params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("UpdatePushList")
	addParameters(cmd, params)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
