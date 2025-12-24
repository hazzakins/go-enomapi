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

// TODO: PE_SetPricing
// TODO: PushDomain
// TODO: RefillAccount
// TODO: SetResellerServicesPricing
// TODO: SetResellerTLDPricing
// TODO: SynchAuthInfo
// TODO: TP_CancelOrder
// TODO: TP_CreateOrder
// TODO: TP_GetDetailsByDomain
// TODO: TP_GetOrder
// TODO: TP_GetOrderDetail
// TODO: TP_GetOrdersByDomain
// TODO: TP_GetOrderReview
// TODO: TP_GetOrderStatuses
// TODO: TP_GetTLDInfo
// TODO: TP_ResendEmail
// TODO: TP_ResubmitLocked
// TODO: TP_SubmitOrder
// TODO: TP_UpdateOrderDetail
// TODO: UpdateAccountPricing
// TODO: UpdatePushList
