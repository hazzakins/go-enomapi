package domains

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type AddBulkDomainsItem struct {
	Domain   enomapi.Domain
	NumYears *int
}

type AddBulkDomainsRequest struct {
	ProductType string
	Items       []AddBulkDomainsItem
	AutoRenew   *bool
	RegLock     *bool
	UseCart     *bool
}

type CancelOrderRequest struct {
	OrderID    string
	DomainName string
}

type GetNameSuggestionsRequest struct {
	SearchTerm     string
	TldList        string
	OnlyTldList    string
	ExcludeTldList string
	MaxResult      *int
	SpinType       *int
	Adult          *bool
	Premium        *bool
	AllGA          *bool
}

type PreconfigureRequest struct {
	Domain             enomapi.Domain
	Load               int
	ExtendedAttributes map[string]string
	AutoRenew          int
	RegLock            int
	IDNCodes           []string
	OptContactReg      string
	OptTechnical       string
	OptAdministrative  string
	OptContactAux      string
	PreConfigDNS       string
	NameServers        []string
	UseHostRecords     *int
	HostNames          []string
	RecordTypes        []string
	Addresses          []string
	ContactParams      map[string]string
	AccessPassword1    string
	AccessPassword2    string
}

type PurchaseRequest struct {
	Domain             enomapi.Domain
	UseDNS             string
	NameServers        []string
	ExtendedAttributes map[string]string
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

type TMGetNoticeRequest struct {
	Domain    enomapi.Domain
	LookupKey string
}

type TMUpdateCartRequest struct {
	Domain        enomapi.Domain
	TcnID         string
	TcnExpDate    string
	TcnAcceptDate string
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

// Upstream documentation: https://api.enom.com/docs/addbulkdomains
func (c *Client) AddBulkDomains(req AddBulkDomainsRequest) (*response.AddBulkDomains, error) {
	resp := internal.AddBulkDomainsResponse{}

	cmd := c.NewCommand("AddBulkDomains")
	cmd.AddParam("ProductType", req.ProductType)
	cmd.AddParam("ListCount", strconv.Itoa(len(req.Items)))
	for i, item := range req.Items {
		index := strconv.Itoa(i + 1)
		cmd.AddParam("SLD"+index, item.Domain.Name)
		cmd.AddParam("TLD"+index, item.Domain.Extension)
		if item.NumYears != nil {
			cmd.AddParam("NumYears"+index, strconv.Itoa(*item.NumYears))
		}
	}
	addBoolIntParam(cmd, "AutoRenew", req.AutoRenew)
	addBoolIntParam(cmd, "RegLock", req.RegLock)
	addBoolIntParam(cmd, "UseCart", req.UseCart)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
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

// Upstream documentation: https://api.enom.com/docs/get-confirmation-settings
func (c *Client) GetConfirmationSettings() (*response.GetConfirmationSettings, error) {
	resp := internal.GetConfirmationSettingsResponse{}

	cmd := c.NewCommand("GetConfirmationSettings")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/get-confirmation-settings
// Upstream documentation: https://api.enom.com/docs/get-ext-attributes
func (c *Client) GetExtAttributes(tld string) (*response.GetExtAttributes, error) {
	resp := internal.GetExtAttributesResponse{}

	cmd := c.NewCommand("GetExtAttributes")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/get-ext-attributes
// Upstream documentation: https://api.enom.com/docs/get-idn-codes
func (c *Client) GetIDNCodes(tld string) (*response.GetIDNCodes, error) {
	resp := internal.GetIDNCodesResponse{}

	cmd := c.NewCommand("GetIDNCodes")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/get-idn-codes
// Upstream documentation: https://api.enom.com/docs/getnamesuggestions
func (c *Client) GetNameSuggestions(req GetNameSuggestionsRequest) (*response.GetNameSuggestions, error) {
	resp := internal.GetNameSuggestionsResponse{}

	cmd := c.NewCommand("GetNameSuggestions")
	cmd.AddParam("SearchTerm", req.SearchTerm)
	if req.TldList != "" {
		cmd.AddParam("TldList", req.TldList)
	}
	if req.OnlyTldList != "" {
		cmd.AddParam("OnlyTldList", req.OnlyTldList)
	}
	if req.ExcludeTldList != "" {
		cmd.AddParam("ExcludeTldList", req.ExcludeTldList)
	}
	addIntParam(cmd, "MaxResult", req.MaxResult)
	addIntParam(cmd, "SpinType", req.SpinType)
	addBoolParam(cmd, "Adult", req.Adult)
	addBoolParam(cmd, "Premium", req.Premium)
	addBoolParam(cmd, "AllGA", req.AllGA)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/getnamesuggestions

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

// Upstream documentation: https://api.enom.com/docs/namespinner
func (c Client) NameSpinner(domain enomapi.Domain, tldList string) (*response.NameSpinner, error) {
	resp := internal.SpinnerResponse{}

	cmd := c.NewCommand("NameSpinner")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("tldlist", tldList)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/preconfigure
func (c *Client) Preconfigure(req PreconfigureRequest) (*response.Preconfigure, error) {
	resp := internal.PreconfigureResponse{}

	cmd := c.NewCommand("Preconfigure")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)
	cmd.AddParam("Load", strconv.Itoa(req.Load))
	cmd.AddParam("AutoRenew", strconv.Itoa(req.AutoRenew))
	cmd.AddParam("RegLock", strconv.Itoa(req.RegLock))
	if len(req.IDNCodes) > 0 {
		for i, code := range req.IDNCodes {
			cmd.AddParam("IDN"+strconv.Itoa(i+1), code)
		}
	}
	if req.OptContactReg != "" {
		cmd.AddParam("OptContactReg", req.OptContactReg)
	}
	if req.OptTechnical != "" {
		cmd.AddParam("OptTechnical", req.OptTechnical)
	}
	if req.OptAdministrative != "" {
		cmd.AddParam("OptAdministrative", req.OptAdministrative)
	}
	if req.OptContactAux != "" {
		cmd.AddParam("OptContactAux", req.OptContactAux)
	}
	if req.PreConfigDNS != "" {
		cmd.AddParam("PreConfigDNS", req.PreConfigDNS)
	}
	if len(req.NameServers) > 0 {
		for i, ns := range req.NameServers {
			cmd.AddParam("NS"+strconv.Itoa(i+1), ns)
		}
	}
	for key, value := range req.ExtendedAttributes {
		cmd.AddParam(key, value)
	}
	addIntParam(cmd, "UseHostRecords", req.UseHostRecords)
	if len(req.HostNames) > 0 {
		cmd.AddParam("HostName", strings.Join(req.HostNames, ","))
	}
	if len(req.RecordTypes) > 0 {
		cmd.AddParam("RecordType", strings.Join(req.RecordTypes, ","))
	}
	if len(req.Addresses) > 0 {
		cmd.AddParam("Address", strings.Join(req.Addresses, ","))
	}
	if req.AccessPassword1 != "" {
		cmd.AddParam("AccessPassword1", req.AccessPassword1)
	}
	if req.AccessPassword2 != "" {
		cmd.AddParam("AccessPassword2", req.AccessPassword2)
	}
	for key, value := range req.ExtendedAttributes {
		cmd.AddParam(key, value)
	}
	for key, value := range req.ContactParams {
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/preconfigure

// Upstream documentation: https://api.enom.com/docs/purchase
func (c Client) Purchase(domain enomapi.Domain) (*response.DomainPurchase, error) {
	resp := internal.DomainPurchase{}

	tldID, err := c.PEGetTLDID(domain.Extension)
	if err != nil {
		return nil, fmt.Errorf("pe_gettldid: %w", err)
	}
	if tldID.TLDID == "" {
		return nil, fmt.Errorf("pe_gettldid: empty TLDID")
	}

	cmd := c.NewCommand("Purchase")
	cmd.AddParam("SLD", domain.Name)
	cmd.AddParam("TLD", domain.Extension)

	err = c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Decode()
}

// PurchaseWithOptions purchases a domain with optional name server settings.
func (c Client) PurchaseWithOptions(req PurchaseRequest) (*response.DomainPurchase, error) {
	resp := internal.DomainPurchase{}

	tldID, err := c.PEGetTLDID(req.Domain.Extension)
	if err != nil {
		return nil, fmt.Errorf("pe_gettldid: %w", err)
	}
	if tldID.TLDID == "" {
		return nil, fmt.Errorf("pe_gettldid: empty TLDID")
	}

	cmd := c.NewCommand("Purchase")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)
	if req.UseDNS != "" {
		cmd.AddParam("UseDNS", req.UseDNS)
	}
	if len(req.NameServers) > 0 {
		for i, ns := range req.NameServers {
			cmd.AddParam("NS"+strconv.Itoa(i+1), ns)
		}
	}

	err = c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Decode()
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

// Upstream documentation: https://api.enom.com/docs/tm-check
func (c *Client) TMCheck(domain enomapi.Domain) (*response.TMCheck, error) {
	resp := internal.TMCheckResponse{}

	cmd := c.NewCommand("TM_Check")
	cmd.AddParam("SLD", domain.Name)
	cmd.AddParam("TLD", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tm-getnotice
func (c *Client) TMGetNotice(req TMGetNoticeRequest) (*response.TMGetNotice, error) {
	resp := internal.TMGetNoticeResponse{}

	cmd := c.NewCommand("TM_GetNotice")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)
	cmd.AddParam("LookupKey", req.LookupKey)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tm-updatecart
func (c *Client) TMUpdateCart(req TMUpdateCartRequest) (*response.TMUpdateCart, error) {
	resp := internal.TMUpdateCartResponse{}

	cmd := c.NewCommand("TM_UpdateCart")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("tcnID", req.TcnID)
	cmd.AddParam("tcnExpDate", req.TcnExpDate)
	cmd.AddParam("tcnAcceptDate", req.TcnAcceptDate)

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

// Upstream documentation: https://api.enom.com/docs/DeleteRegistration
func (c *Client) DeleteRegistration(domain enomapi.Domain, endUserIP string) (*response.DeleteRegistration, error) {
	resp := internal.DeleteRegistrationResponse{}

	cmd := c.NewCommand("DeleteRegistration")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("EndUserIP", endUserIP)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

func addBoolParam(cmd *enomapi.Command, key string, value *bool) {
	if value == nil {
		return
	}
	cmd.AddParam(key, boolToTitle(*value))
}

func addBoolIntParam(cmd *enomapi.Command, key string, value *bool) {
	if value == nil {
		return
	}
	cmd.AddParam(key, boolToInt(*value))
}

func addIntParam(cmd *enomapi.Command, key string, value *int) {
	if value == nil {
		return
	}
	cmd.AddParam(key, strconv.Itoa(*value))
}

func boolToTitle(value bool) string {
	if value {
		return "True"
	}
	return "False"
}

func boolToInt(value bool) string {
	if value {
		return "1"
	}
	return "0"
}
