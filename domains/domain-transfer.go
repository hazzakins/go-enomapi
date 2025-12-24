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

// Upstream documentation: https://api.enom.com/docs/pe-set-pricing
func (c *Client) PESetPricing(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "PE_SetPricing", params)
}

// Upstream documentation: https://api.enom.com/docs/pushdomain
func (c *Client) PushDomain(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "PushDomain", params)
}

// Upstream documentation: https://api.enom.com/docs/refill-account
func (c *Client) RefillAccount(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "RefillAccount", params)
}

// Upstream documentation: https://api.enom.com/docs/setresellerservicespricing
func (c *Client) SetResellerServicesPricing(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "SetResellerServicesPricing", params)
}

// Upstream documentation: https://api.enom.com/docs/setresellertldpricing
func (c *Client) SetResellerTLDPricing(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "SetResellerTLDPricing", params)
}

// Upstream documentation: https://api.enom.com/docs/synchauthinfo
func (c *Client) SynchAuthInfo(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "SynchAuthInfo", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-cancelorder
func (c *Client) TPCancelOrder(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_CancelOrder", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-createorder
func (c *Client) TPCreateOrder(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_CreateOrder", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-getdetailsbydomain
func (c *Client) TPGetDetailsByDomain(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_GetDetailsByDomain", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-getorder
func (c *Client) TPGetOrder(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_GetOrder", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderdetail
func (c *Client) TPGetOrderDetail(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_GetOrderDetail", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-getordersbydomain
func (c *Client) TPGetOrdersByDomain(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_GetOrdersByDomain", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderreview
func (c *Client) TPGetOrderReview(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_GetOrderReview", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderstatuses
func (c *Client) TPGetOrderStatuses(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_GetOrderStatuses", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-get-tld-info
func (c *Client) TPGetTLDInfo(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_GetTLDInfo", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-resendemail
func (c *Client) TPResendEmail(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_ResendEmail", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-resubmit-locked
func (c *Client) TPResubmitLocked(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_ResubmitLocked", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-submitorder
func (c *Client) TPSubmitOrder(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_SubmitOrder", params)
}

// Upstream documentation: https://api.enom.com/docs/tp-updateorderdetail
func (c *Client) TPUpdateOrderDetail(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "TP_UpdateOrderDetail", params)
}

// Upstream documentation: https://api.enom.com/docs/updateaccountpricing
func (c *Client) UpdateAccountPricing(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "UpdateAccountPricing", params)
}

// Upstream documentation: https://api.enom.com/docs/updatepushlist
func (c *Client) UpdatePushList(params ParameterSet) (*response.Generic, error) {
	return executeGenericCommand(c, "UpdatePushList", params)
}
