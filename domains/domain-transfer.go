package domains

import (
	"strconv"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type PESetPricingRequest struct {
	Parameters map[string]string
}

type PushDomainRequest struct {
	AccountID             string
	PushContact           *bool
	IRTPOptOut            *bool
	IRTPOptOutReason      string
	IRTPEmailLanguageCode string
	AdditionalParams      map[string]string
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

type SynchAuthInfoRequest struct {
	Parameters map[string]string
}

type TPCancelOrderRequest struct {
	OrderID          string
	AdditionalParams map[string]string
}

type TPCreateOrderRequest struct {
	OrderType        string
	DomainCount      int
	SLDs             []string
	TLDs             []string
	AuthInfos        []string
	PreConfig        *bool
	AdditionalParams map[string]string
}

type TPGetDetailsByDomainRequest struct {
	DomainName       string
	AdditionalParams map[string]string
}

type TPGetOrderRequest struct {
	TransferOrderID  string
	AdditionalParams map[string]string
}

type TPGetOrderDetailRequest struct {
	TransferOrderID  string
	AdditionalParams map[string]string
}

type TPGetOrdersByDomainRequest struct {
	DomainName       string
	AdditionalParams map[string]string
}

type TPGetOrderReviewRequest struct {
	TransferOrderID  string
	AdditionalParams map[string]string
}

type TPGetOrderStatusesRequest struct {
	AdditionalParams map[string]string
}

type TPGetTLDInfoRequest struct {
	AdditionalParams map[string]string
}

type TPResendEmailRequest struct {
	TransferOrderID  string
	AdditionalParams map[string]string
}

type TPResubmitLockedRequest struct {
	TransferOrderID  string
	AdditionalParams map[string]string
}

type TPSubmitOrderRequest struct {
	TransferOrderID  string
	AdditionalParams map[string]string
}

type TPUpdateOrderDetailRequest struct {
	TransferOrderID  string
	AdditionalParams map[string]string
}

type UpdateAccountPricingRequest struct {
	Parameters map[string]string
}

type UpdatePushListRequest struct {
	AdditionalParams map[string]string
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
func (c *Client) PESetPricing(req PESetPricingRequest) (*response.PESetPricing, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("PE_SetPricing")
	addParams(cmd, req.Parameters)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/pushdomain
func (c *Client) PushDomain(req PushDomainRequest) (*response.PushDomain, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("PushDomain")
	cmd.AddParam("AccountID", req.AccountID)
	addBoolIntParam(cmd, "PushContact", req.PushContact)
	addBoolIntParam(cmd, "IRTPOptOut", req.IRTPOptOut)
	if req.IRTPOptOutReason != "" {
		cmd.AddParam("IRTPOptOutReason", req.IRTPOptOutReason)
	}
	if req.IRTPEmailLanguageCode != "" {
		cmd.AddParam("IRTPEmailLanguageCode", req.IRTPEmailLanguageCode)
	}
	addParams(cmd, req.AdditionalParams)

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

// Upstream documentation: https://api.enom.com/docs/synchauthinfo
func (c *Client) SynchAuthInfo(req SynchAuthInfoRequest) (*response.SynchAuthInfo, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SynchAuthInfo")
	addParams(cmd, req.Parameters)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-cancelorder
func (c *Client) TPCancelOrder(req TPCancelOrderRequest) (*response.TPCancelOrder, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_CancelOrder")
	cmd.AddParam("TransferOrderID", req.OrderID)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-createorder
func (c *Client) TPCreateOrder(req TPCreateOrderRequest) (*response.TPCreateOrder, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_CreateOrder")
	cmd.AddParam("OrderType", req.OrderType)
	cmd.AddParam("DomainCount", strconv.Itoa(req.DomainCount))
	for i, sld := range req.SLDs {
		cmd.AddParam("SLD"+strconv.Itoa(i+1), sld)
	}
	for i, tld := range req.TLDs {
		cmd.AddParam("TLD"+strconv.Itoa(i+1), tld)
	}
	for i, auth := range req.AuthInfos {
		cmd.AddParam("AuthInfo"+strconv.Itoa(i+1), auth)
	}
	addBoolIntParam(cmd, "PreConfig", req.PreConfig)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getdetailsbydomain
func (c *Client) TPGetDetailsByDomain(req TPGetDetailsByDomainRequest) (*response.TPGetDetailsByDomain, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetDetailsByDomain")
	cmd.AddParam("DomainName", req.DomainName)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorder
func (c *Client) TPGetOrder(req TPGetOrderRequest) (*response.TPGetOrder, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrder")
	cmd.AddParam("TransferOrderID", req.TransferOrderID)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderdetail
func (c *Client) TPGetOrderDetail(req TPGetOrderDetailRequest) (*response.TPGetOrderDetail, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderDetail")
	cmd.AddParam("TransferOrderID", req.TransferOrderID)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getordersbydomain
func (c *Client) TPGetOrdersByDomain(req TPGetOrdersByDomainRequest) (*response.TPGetOrdersByDomain, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrdersByDomain")
	cmd.AddParam("DomainName", req.DomainName)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderreview
func (c *Client) TPGetOrderReview(req TPGetOrderReviewRequest) (*response.TPGetOrderReview, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderReview")
	cmd.AddParam("TransferOrderID", req.TransferOrderID)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderstatuses
func (c *Client) TPGetOrderStatuses(req TPGetOrderStatusesRequest) (*response.TPGetOrderStatuses, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderStatuses")
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-get-tld-info
func (c *Client) TPGetTLDInfo(req TPGetTLDInfoRequest) (*response.TPGetTLDInfo, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetTLDInfo")
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-resendemail
func (c *Client) TPResendEmail(req TPResendEmailRequest) (*response.TPResendEmail, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_ResendEmail")
	cmd.AddParam("TransferOrderID", req.TransferOrderID)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-resubmit-locked
func (c *Client) TPResubmitLocked(req TPResubmitLockedRequest) (*response.TPResubmitLocked, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_ResubmitLocked")
	cmd.AddParam("TransferOrderID", req.TransferOrderID)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-submitorder
func (c *Client) TPSubmitOrder(req TPSubmitOrderRequest) (*response.TPSubmitOrder, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_SubmitOrder")
	cmd.AddParam("TransferOrderID", req.TransferOrderID)
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

// Upstream documentation: https://api.enom.com/docs/tp-updateorderdetail
func (c *Client) TPUpdateOrderDetail(req TPUpdateOrderDetailRequest) (*response.TPUpdateOrderDetail, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_UpdateOrderDetail")
	cmd.AddParam("TransferOrderID", req.TransferOrderID)
	addParams(cmd, req.AdditionalParams)

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

// Upstream documentation: https://api.enom.com/docs/updatepushlist
func (c *Client) UpdatePushList(req UpdatePushListRequest) (*response.UpdatePushList, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("UpdatePushList")
	addParams(cmd, req.AdditionalParams)

	err := c.Execute(cmd, &resp)
	return toTransferResponse(&resp), err
}

func addParams(cmd *enomapi.Command, params map[string]string) {
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}
}

func toTransferResponse(resp *internal.GenericResponse) *response.TransferResponse {
	decoded := resp.Decode()
	return &response.TransferResponse{ResponseMeta: decoded.ResponseMeta}
}
