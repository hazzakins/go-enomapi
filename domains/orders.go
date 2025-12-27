package domains

import (
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type CancelOrderRequest struct {
	OrderID    string
	DomainName string
}

type QueueDomainPurchaseRequest struct {
	ItemList        string
	UseCreditCard   *bool
	Amount          string
	EndUserIP       string
	CCName          string
	CCExpMonth      string
	CCNumber        string
	CCExpYear       string
	CVV2            string
	CCAddress       string
	CCCity          string
	CCStateProvince string
	CCZip           string
	CCCountry       string
}

type QueueGetDomainsRequest struct {
	DisplayMetrics   *bool
	DisplayDomains   *bool
	RecordStart      *int
	PagingSize       *int
	SortBy           string
	SortOrder        string
	FilterTLD        string
	FilterQueue      string
	FilterStatus     string
	FilterStatusDesc string
	DomainNameFilter string
}

type QueueGetExtAttributesRequest struct {
	QIDList string
}

type QueueGetOrderDetailRequest struct {
	OrderID           string
	PortalUserPartyID string
}

type QueueGetOrdersRequest struct {
	PortalUserPartyID string
}

type GetAgreementPageRequest struct {
	Page     string
	Language string
}

type QueueGetInfoRequest struct {
	Category            string
	FilterTLD           string
	FilterTLDStatusDesc string
	FilterQStatusName   string
	FilterExtAttributes string
	FilterStartDate     string
	FilterEndDate       string
	DisplayComingSoon   string
}

// Upstream documentation: https://api.enom.com/docs/cancel-order
func (c *Client) CancelOrder(req CancelOrderRequest) (*response.CancelOrder, error) {
	resp := internal.CancelOrderResponse{}

	cmd := c.NewCommand("CancelOrder")
	cmd.AddParam("OrderID", req.OrderID)
	cmd.AddParam("DomainName", req.DomainName)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/get-confirmation-settings
func (c *Client) GetConfirmationSettings() (*response.GetConfirmationSettings, error) {
	resp := internal.GetConfirmationSettingsResponse{}

	cmd := c.NewCommand("GetConfirmationSettings")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/queue-domain-purchase
func (c *Client) QueueDomainPurchase(req QueueDomainPurchaseRequest) (*response.QueueDomainPurchase, error) {
	resp := internal.QueueDomainPurchaseResponse{}

	cmd := c.NewCommand("Queue_DomainPurchase")
	cmd.AddParam("ItemList", req.ItemList)
	addBoolParam(cmd, "UseCreditCard", req.UseCreditCard)
	if req.Amount != "" {
		cmd.AddParam("Amount", req.Amount)
	}
	if req.EndUserIP != "" {
		cmd.AddParam("EndUserIP", req.EndUserIP)
	}
	if req.CCName != "" {
		cmd.AddParam("CCName", req.CCName)
	}
	if req.CCExpMonth != "" {
		cmd.AddParam("CCExpMonth", req.CCExpMonth)
	}
	if req.CCNumber != "" {
		cmd.AddParam("CCNumber", req.CCNumber)
	}
	if req.CCExpYear != "" {
		cmd.AddParam("CCExpYear", req.CCExpYear)
	}
	if req.CVV2 != "" {
		cmd.AddParam("CVV2", req.CVV2)
	}
	if req.CCAddress != "" {
		cmd.AddParam("CCAddress", req.CCAddress)
	}
	if req.CCCity != "" {
		cmd.AddParam("CCCity", req.CCCity)
	}
	if req.CCStateProvince != "" {
		cmd.AddParam("CCStateProvince", req.CCStateProvince)
	}
	if req.CCZip != "" {
		cmd.AddParam("CCZip", req.CCZip)
	}
	if req.CCCountry != "" {
		cmd.AddParam("CCCountry", req.CCCountry)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/Queue_GetDomains
func (c *Client) QueueGetDomains(req QueueGetDomainsRequest) (*response.QueueGetDomains, error) {
	resp := internal.QueueGetDomainsResponse{}

	cmd := c.NewCommand("Queue_GetDomains")
	addBoolParam(cmd, "DisplayMetrics", req.DisplayMetrics)
	addBoolParam(cmd, "DisplayDomains", req.DisplayDomains)
	addIntParam(cmd, "RecordStart", req.RecordStart)
	addIntParam(cmd, "PagingSize", req.PagingSize)
	if req.SortBy != "" {
		cmd.AddParam("SortBy", req.SortBy)
	}
	if req.SortOrder != "" {
		cmd.AddParam("SortOrder", req.SortOrder)
	}
	if req.FilterTLD != "" {
		cmd.AddParam("FilterTLD", req.FilterTLD)
	}
	if req.FilterQueue != "" {
		cmd.AddParam("FilterQueue", req.FilterQueue)
	}
	if req.FilterStatus != "" {
		cmd.AddParam("FilterStatus", req.FilterStatus)
	}
	if req.FilterStatusDesc != "" {
		cmd.AddParam("FilterStatusDesc", req.FilterStatusDesc)
	}
	if req.DomainNameFilter != "" {
		cmd.AddParam("DomainNameFilter", req.DomainNameFilter)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/queue-get-ext-attributes
func (c *Client) QueueGetExtAttributes(req QueueGetExtAttributesRequest) (*response.QueueGetExtAttributes, error) {
	resp := internal.QueueGetExtAttributesResponse{}

	cmd := c.NewCommand("Queue_GetExtAttributes")
	cmd.AddParam("QIDList", req.QIDList)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/queue-get-order-detail
func (c *Client) QueueGetOrderDetail(req QueueGetOrderDetailRequest) (*response.QueueGetOrderDetail, error) {
	resp := internal.QueueGetOrderDetailResponse{}

	cmd := c.NewCommand("Queue_GetOrderDetail")
	cmd.AddParam("OrderID", req.OrderID)
	if req.PortalUserPartyID != "" {
		cmd.AddParam("PortalUserPartyID", req.PortalUserPartyID)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/queue-get-orders
func (c *Client) QueueGetOrders(req QueueGetOrdersRequest) (*response.QueueGetOrders, error) {
	resp := internal.QueueGetOrdersResponse{}

	cmd := c.NewCommand("Queue_GetOrders")
	if req.PortalUserPartyID != "" {
		cmd.AddParam("PortalUserPartyID", req.PortalUserPartyID)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/GetAgreementPage
func (c *Client) GetAgreementPage(req GetAgreementPageRequest) (*response.GetAgreementPage, error) {
	resp := internal.GetAgreementPageResponse{}

	cmd := c.NewCommand("GetAgreementPage")
	if req.Page != "" {
		cmd.AddParam("Page", req.Page)
	}
	if req.Language != "" {
		cmd.AddParam("Language", req.Language)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/Queue_GetInfo
func (c *Client) QueueGetInfo(req QueueGetInfoRequest) (*response.QueueGetInfo, error) {
	resp := internal.QueueGetInfoResponse{}

	cmd := c.NewCommand("Queue_GetInfo")
	if req.Category != "" {
		cmd.AddParam("Category", req.Category)
	}
	if req.FilterTLD != "" {
		cmd.AddParam("FilterTLD", req.FilterTLD)
	}
	if req.FilterTLDStatusDesc != "" {
		cmd.AddParam("FilterTLDStatusDesc", req.FilterTLDStatusDesc)
	}
	if req.FilterQStatusName != "" {
		cmd.AddParam("FilterQStatusName", req.FilterQStatusName)
	}
	if req.FilterExtAttributes != "" {
		cmd.AddParam("FilterExtAttributes", req.FilterExtAttributes)
	}
	if req.FilterStartDate != "" {
		cmd.AddParam("FilterStartDate", req.FilterStartDate)
	}
	if req.FilterEndDate != "" {
		cmd.AddParam("FilterEndDate", req.FilterEndDate)
	}
	if req.DisplayComingSoon != "" {
		cmd.AddParam("DisplayComingSoon", req.DisplayComingSoon)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
