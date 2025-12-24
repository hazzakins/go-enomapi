package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

type AddBulkDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	AddBulkDomains AddBulkDomains `xml:"AddBulkDomains"`
	Success        string         `xml:"Success"`
	UseCart        string         `xml:"UseCart"`
}

type AddBulkDomains struct {
	Items              []AddBulkDomainsItem `xml:"Item"`
	ListCount          int                  `xml:"ListCount"`
	CartErrors         int                  `xml:"CartErrors"`
	AllItemsSuccessful string               `xml:"AllItemsSuccessful"`
	CartItems          int                  `xml:"CartItems"`
}

type AddBulkDomainsItem struct {
	WscAccountOverride string `xml:"WscAccountOverride"`
	ItemName           string `xml:"ItemName"`
	ItemID             string `xml:"ItemId"`
	Price              string `xml:"Price"`
	ICANNFees          string `xml:"ICANNFees"`
	CartItemID         string `xml:"CartItemID"`
	NewDomainNameID    string `xml:"NewDomainNameID"`
	ItemAdded          string `xml:"ItemAdded"`
	ItemError          string `xml:"ItemError"`
	DomainName         string `xml:"DomainName"`
}

type CancelOrderResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Order CancelOrder `xml:"Order"`
}

type CancelOrder struct {
	OrderID       string             `xml:"OrderID"`
	Success       string             `xml:"Success"`
	Result        string             `xml:"Result"`
	Domains       []string           `xml:"Domains>DomainName"`
	DomainEntries []CancelOrderEntry `xml:"Domains>Domain"`
}

type CancelOrderEntry struct {
	DomainName  string `xml:"DomainName"`
	Description string `xml:"Description"`
}

type GetConfirmationSettingsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ConfirmationSettings ConfirmationSettings `xml:"ConfirmationSettings"`
}

type ConfirmationSettings struct {
	OrderConfirmation            string `xml:"OrderConfirmation"`
	TransferOrderConfirmation    string `xml:"TransferOrderConfirmation"`
	OrderConfirmationBCC         string `xml:"OrderConfirmationBCC"`
	TransferOrderConfirmationBCC string `xml:"TransferOrderConfirmationBCC"`
	EmailHead                    string `xml:"EmailHead"`
	EmailTail                    string `xml:"EmailTail"`
}

type GetExtAttributesResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Attributes ExtAttributes `xml:"Attributes"`
}

type ExtAttributes struct {
	Items []ExtAttribute `xml:"Attribute"`
}

type ExtAttribute struct {
	ID          int                  `xml:"ID"`
	Name        string               `xml:"Name"`
	Application int                  `xml:"Application"`
	UserDefined bool                 `xml:"UserDefined"`
	Required    int                  `xml:"Required"`
	Description string               `xml:"Description"`
	IsChild     int                  `xml:"IsChild"`
	Options     []ExtAttributeOption `xml:"Options>Option"`
}

type ExtAttributeOption struct {
	ID          int    `xml:"ID"`
	Value       string `xml:"Value"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

type GetIDNCodesResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	TLDs IDNCodesTLDs `xml:"tlds"`
}

type IDNCodesTLDs struct {
	TLDs []IDNCodesTLD `xml:"tld"`
}

type IDNCodesTLD struct {
	TLD       string        `xml:"tld,attr"`
	Languages []IDNLanguage `xml:"language"`
}

type IDNLanguage struct {
	Code string `xml:"code,attr"`
	Name string `xml:"name,attr"`
}

type GetNameSuggestionsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Suggestions RawXML `xml:"suggestions"`
}

type PreconfigureResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	PreConfigSuccess string `xml:"PreConfigSuccess"`
	Count            int    `xml:"Count"`
}

type QueueDomainPurchaseResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Purchase QueueDomainPurchase `xml:"Queue_DomainPurchase"`
}

type QueueDomainPurchase struct {
	Success              string `xml:"Success"`
	OrderID              string `xml:"OrderID"`
	Amount               string `xml:"Amount"`
	TransactionNumber    string `xml:"TransactionNumber"`
	TransactionReference string `xml:"TransactionReference"`
}

type QueueGetDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Queue QueueGetDomains `xml:"Queue_GetDomains"`
}

type QueueGetDomains struct {
	Metrics   *QueueMetrics `xml:"Metrics"`
	Domains   []QueueDomain `xml:"Domains>Domain"`
	ItemCount int           `xml:"ItemCount"`
	ItemTotal int           `xml:"ItemTotal"`
}

type QueueMetrics struct {
	WatchlistTotal         int    `xml:"WatchlistTotal"`
	WatchlistPreorder      int    `xml:"WatchlistPreorder"`
	WatchlistDomainWatched int    `xml:"WathclistDomainWatched"`
	WatchlistTLDWatched    int    `xml:"WatchlistTLDWatched"`
	NewTLDOrders           int    `xml:"NewTLDOrders"`
	NewTLDValidated        int    `xml:"NewTLDValidated"`
	NewTLDPending          int    `xml:"NewTLDPending"`
	NewTLDRejected         int    `xml:"NewTLDRejected"`
	QueuesDisabled         string `xml:"QueuesDisabled"`
	PreregDisabled         string `xml:"PreregDisabled"`
}

type QueueDomain struct {
	OrderID    string `xml:"OrderID"`
	DomainName string `xml:"DomainName"`
	TLDID      string `xml:"TLDID"`
	TLDName    string `xml:"TLDName"`
	TLDDesc    string `xml:"TLDDesc"`
	DateStart  string `xml:"DateStart"`
	DateEnd    string `xml:"DateEnd"`
	Price      string `xml:"Price"`
	QID        string `xml:"QID"`
	QName      string `xml:"QName"`
	StatusID   string `xml:"StatusID"`
	StatusName string `xml:"StatusName"`
}

type QueueGetExtAttributesResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Queues QueueExtAttributes `xml:"Queue_GetExtAttributes"`
}

type QueueExtAttributes struct {
	Queues []QueueExtAttributeEntry `xml:"Queues>Queue"`
}

type QueueExtAttributeEntry struct {
	ID         string `xml:"ID"`
	TLD        string `xml:"TLD"`
	Attributes RawXML `xml:"Attributes"`
}

type QueueGetOrderDetailResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Detail QueueOrderDetail `xml:"Queue_GetOrderDetail"`
}

type QueueOrderDetail struct {
	Orders       []QueueOrderDetailEntry `xml:"Orders>Order"`
	RecordStart  int                     `xml:"RecordStart"`
	PagingSize   int                     `xml:"PagingSize"`
	NextRecord   int                     `xml:"NextRecord"`
	TotalResults int                     `xml:"TotalResults"`
}

type QueueOrderDetailEntry struct {
	OrderID         string `xml:"OrderID"`
	OrderDate       string `xml:"OrderDate"`
	PaidAmount      string `xml:"PaidAmount"`
	DomainName      string `xml:"DomainName"`
	OrderType       string `xml:"OrderType"`
	ExtTime         string `xml:"ExtTime"`
	OrderStatus     string `xml:"OrderStatus"`
	RegistrationFee string `xml:"RegistrationFee"`
	ApplicationFee  string `xml:"ApplicationFee"`
	TotalFee        string `xml:"TotalFee"`
}

type QueueGetOrdersResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Orders QueueOrders `xml:"Queue_GetOrders"`
}

type QueueOrders struct {
	Orders    []QueueOrderEntry `xml:"Orders>Order"`
	ItemCount int               `xml:"ItemCount"`
	ItemTotal int               `xml:"ItemTotal"`
}

type QueueOrderEntry struct {
	OrderID    string `xml:"OrderID"`
	StatusID   string `xml:"StatusID"`
	StatusName string `xml:"StatusName"`
	StatusDesc string `xml:"StatusDesc"`
	Success    string `xml:"Success"`
	Failed     string `xml:"Failed"`
	Pending    string `xml:"Pending"`
	OrderDate  string `xml:"OrderDate"`
}

type TMCheckResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	LookupKey string `xml:"LookupKey"`
}

type TMGetNoticeResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Notice TMNotice `xml:"TMNotice"`
}

type TMNotice struct {
	TcnID        string `xml:"tcnID"`
	TcnStartDate string `xml:"tcnStartDate"`
	TcnExpDate   string `xml:"tcnExpDate"`
	SLD          string `xml:"Sld"`
	Claims       RawXML `xml:"Claims"`
}

type TMUpdateCartResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Success string `xml:"Success"`
}

type GetAgreementPageResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Content RawXML `xml:"content"`
}

type QueueGetInfoResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Info QueueInfo `xml:"Queue_GetInfo"`
}

type QueueInfo struct {
	Queues       []QueueInfoEntry `xml:"Queue"`
	TotalRecords int              `xml:"TotalRecords"`
}

type QueueInfoEntry struct {
	TLD              string `xml:"TLD"`
	NativeIDN        string `xml:"NativeIDN"`
	CategoryName     string `xml:"CategoryName"`
	CategoryDesc     string `xml:"CategoryDesc"`
	TLDDescription   string `xml:"TLDDescription"`
	TLDStatusID      string `xml:"TLDStatusID"`
	TLDStatusDesc    string `xml:"TLDStatusDesc"`
	QID              string `xml:"QID"`
	QName            string `xml:"QName"`
	QStatusID        string `xml:"QStatusID"`
	QStatusName      string `xml:"QStatusName"`
	ExtAttributes    string `xml:"ExtAttributes"`
	StartDate        string `xml:"StartDate"`
	EndDate          string `xml:"EndDate"`
	AnnouncementDate string `xml:"AnnouncementDate"`
	Price            string `xml:"Price"`
	InEAP            string `xml:"InEAP"`
}

type DeleteRegistrationResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DeleteDomain DeleteDomain `xml:"deletedomain"`
	ErrString    string       `xml:"ErrString"`
	ErrSource    string       `xml:"ErrSource"`
	ErrSection   string       `xml:"ErrSection"`
}

type DeleteDomain struct {
	DomainDeleted string `xml:"domaindeleted"`
}

func (r *AddBulkDomainsResponse) Decode() *response.AddBulkDomains {
	return &response.AddBulkDomains{
		Items:              decodeAddBulkItems(r.AddBulkDomains.Items),
		ListCount:          r.AddBulkDomains.ListCount,
		CartErrors:         r.AddBulkDomains.CartErrors,
		AllItemsSuccessful: r.AddBulkDomains.AllItemsSuccessful,
		CartItems:          r.AddBulkDomains.CartItems,
		Success:            r.Success,
		UseCart:            r.UseCart,
		ResponseMeta:       decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *CancelOrderResponse) Decode() *response.CancelOrder {
	domains := decodeCancelOrderDomains(r.Order.DomainEntries)
	if len(domains) == 0 && len(r.Order.Domains) > 0 {
		for _, name := range r.Order.Domains {
			domains = append(domains, response.CancelOrderDomain{
				DomainName: name,
			})
		}
	}
	return &response.CancelOrder{
		OrderID:      r.Order.OrderID,
		Success:      r.Order.Success,
		Result:       r.Order.Result,
		Domains:      domains,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetConfirmationSettingsResponse) Decode() *response.GetConfirmationSettings {
	return &response.GetConfirmationSettings{
		OrderConfirmation:            r.ConfirmationSettings.OrderConfirmation,
		TransferOrderConfirmation:    r.ConfirmationSettings.TransferOrderConfirmation,
		OrderConfirmationBCC:         r.ConfirmationSettings.OrderConfirmationBCC,
		TransferOrderConfirmationBCC: r.ConfirmationSettings.TransferOrderConfirmationBCC,
		EmailHead:                    r.ConfirmationSettings.EmailHead,
		EmailTail:                    r.ConfirmationSettings.EmailTail,
		ResponseMeta:                 decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetExtAttributesResponse) Decode() *response.GetExtAttributes {
	return &response.GetExtAttributes{
		Attributes:   decodeExtAttributes(r.Attributes.Items),
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetIDNCodesResponse) Decode() *response.GetIDNCodes {
	return &response.GetIDNCodes{
		TLDs:         decodeIDNCodesTLDs(r.TLDs.TLDs),
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetNameSuggestionsResponse) Decode() *response.GetNameSuggestions {
	return &response.GetNameSuggestions{
		SuggestionsXML: r.Suggestions.Raw,
		ResponseMeta:   decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *PreconfigureResponse) Decode() *response.Preconfigure {
	return &response.Preconfigure{
		PreConfigSuccess: r.PreConfigSuccess,
		Count:            r.Count,
		ResponseMeta:     decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *QueueDomainPurchaseResponse) Decode() *response.QueueDomainPurchase {
	return &response.QueueDomainPurchase{
		Success:              r.Purchase.Success,
		OrderID:              r.Purchase.OrderID,
		Amount:               r.Purchase.Amount,
		TransactionNumber:    r.Purchase.TransactionNumber,
		TransactionReference: r.Purchase.TransactionReference,
		ResponseMeta:         decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *QueueGetDomainsResponse) Decode() *response.QueueGetDomains {
	return &response.QueueGetDomains{
		Metrics:      decodeQueueMetrics(r.Queue.Metrics),
		Domains:      decodeQueueDomains(r.Queue.Domains),
		ItemCount:    r.Queue.ItemCount,
		ItemTotal:    r.Queue.ItemTotal,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *QueueGetExtAttributesResponse) Decode() *response.QueueGetExtAttributes {
	return &response.QueueGetExtAttributes{
		Queues:       decodeQueueExtAttributes(r.Queues.Queues),
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *QueueGetOrderDetailResponse) Decode() *response.QueueGetOrderDetail {
	return &response.QueueGetOrderDetail{
		Orders:       decodeQueueOrderDetails(r.Detail.Orders),
		RecordStart:  r.Detail.RecordStart,
		PagingSize:   r.Detail.PagingSize,
		NextRecord:   r.Detail.NextRecord,
		TotalResults: r.Detail.TotalResults,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *QueueGetOrdersResponse) Decode() *response.QueueGetOrders {
	return &response.QueueGetOrders{
		Orders:       decodeQueueOrders(r.Orders.Orders),
		ItemCount:    r.Orders.ItemCount,
		ItemTotal:    r.Orders.ItemTotal,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *TMCheckResponse) Decode() *response.TMCheck {
	return &response.TMCheck{
		LookupKey:    r.LookupKey,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *TMGetNoticeResponse) Decode() *response.TMGetNotice {
	return &response.TMGetNotice{
		TcnID:        r.Notice.TcnID,
		TcnStartDate: r.Notice.TcnStartDate,
		TcnExpDate:   r.Notice.TcnExpDate,
		SLD:          r.Notice.SLD,
		ClaimsXML:    r.Notice.Claims.Raw,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *TMUpdateCartResponse) Decode() *response.TMUpdateCart {
	return &response.TMUpdateCart{
		Success:      r.Success,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetAgreementPageResponse) Decode() *response.GetAgreementPage {
	return &response.GetAgreementPage{
		ContentXML:   r.Content.Raw,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *QueueGetInfoResponse) Decode() *response.QueueGetInfo {
	return &response.QueueGetInfo{
		Queues:       decodeQueueInfoEntries(r.Info.Queues),
		TotalRecords: r.Info.TotalRecords,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *DeleteRegistrationResponse) Decode() *response.DeleteRegistration {
	return &response.DeleteRegistration{
		DomainDeleted: r.DeleteDomain.DomainDeleted,
		ErrString:     r.ErrString,
		ErrSource:     r.ErrSource,
		ErrSection:    r.ErrSection,
		ResponseMeta:  decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func decodeAddBulkItems(items []AddBulkDomainsItem) []response.AddBulkDomainsItem {
	if len(items) == 0 {
		return nil
	}
	result := make([]response.AddBulkDomainsItem, 0, len(items))
	for _, item := range items {
		result = append(result, response.AddBulkDomainsItem{
			WscAccountOverride: item.WscAccountOverride,
			ItemName:           item.ItemName,
			ItemID:             item.ItemID,
			Price:              item.Price,
			ICANNFees:          item.ICANNFees,
			CartItemID:         item.CartItemID,
			NewDomainNameID:    item.NewDomainNameID,
			ItemAdded:          item.ItemAdded,
			ItemError:          item.ItemError,
			DomainName:         item.DomainName,
		})
	}
	return result
}

func decodeCancelOrderDomains(entries []CancelOrderEntry) []response.CancelOrderDomain {
	if len(entries) == 0 {
		return nil
	}
	result := make([]response.CancelOrderDomain, 0, len(entries))
	for _, entry := range entries {
		result = append(result, response.CancelOrderDomain{
			DomainName:  entry.DomainName,
			Description: entry.Description,
		})
	}
	return result
}

func decodeExtAttributes(attrs []ExtAttribute) []response.ExtAttribute {
	if len(attrs) == 0 {
		return nil
	}
	result := make([]response.ExtAttribute, 0, len(attrs))
	for _, attr := range attrs {
		result = append(result, response.ExtAttribute{
			ID:          attr.ID,
			Name:        attr.Name,
			Application: attr.Application,
			UserDefined: attr.UserDefined,
			Required:    attr.Required,
			Description: attr.Description,
			IsChild:     attr.IsChild,
			Options:     decodeExtAttributeOptions(attr.Options),
		})
	}
	return result
}

func decodeExtAttributeOptions(options []ExtAttributeOption) []response.ExtAttributeOption {
	if len(options) == 0 {
		return nil
	}
	result := make([]response.ExtAttributeOption, 0, len(options))
	for _, option := range options {
		result = append(result, response.ExtAttributeOption{
			ID:          option.ID,
			Value:       option.Value,
			Title:       option.Title,
			Description: option.Description,
		})
	}
	return result
}

func decodeIDNCodesTLDs(tlds []IDNCodesTLD) []response.IDNCodesTLD {
	if len(tlds) == 0 {
		return nil
	}
	result := make([]response.IDNCodesTLD, 0, len(tlds))
	for _, tld := range tlds {
		result = append(result, response.IDNCodesTLD{
			TLD:       tld.TLD,
			Languages: decodeIDNLanguages(tld.Languages),
		})
	}
	return result
}

func decodeIDNLanguages(langs []IDNLanguage) []response.IDNLanguage {
	if len(langs) == 0 {
		return nil
	}
	result := make([]response.IDNLanguage, 0, len(langs))
	for _, lang := range langs {
		result = append(result, response.IDNLanguage{
			Code: lang.Code,
			Name: lang.Name,
		})
	}
	return result
}

func decodeQueueMetrics(metrics *QueueMetrics) *response.QueueMetrics {
	if metrics == nil {
		return nil
	}
	return &response.QueueMetrics{
		WatchlistTotal:         metrics.WatchlistTotal,
		WatchlistPreorder:      metrics.WatchlistPreorder,
		WatchlistDomainWatched: metrics.WatchlistDomainWatched,
		WatchlistTLDWatched:    metrics.WatchlistTLDWatched,
		NewTLDOrders:           metrics.NewTLDOrders,
		NewTLDValidated:        metrics.NewTLDValidated,
		NewTLDPending:          metrics.NewTLDPending,
		NewTLDRejected:         metrics.NewTLDRejected,
		QueuesDisabled:         metrics.QueuesDisabled,
		PreregDisabled:         metrics.PreregDisabled,
	}
}

func decodeQueueDomains(domains []QueueDomain) []response.QueueDomain {
	if len(domains) == 0 {
		return nil
	}
	result := make([]response.QueueDomain, 0, len(domains))
	for _, domain := range domains {
		result = append(result, response.QueueDomain{
			OrderID:    domain.OrderID,
			DomainName: domain.DomainName,
			TLDID:      domain.TLDID,
			TLDName:    domain.TLDName,
			TLDDesc:    domain.TLDDesc,
			DateStart:  domain.DateStart,
			DateEnd:    domain.DateEnd,
			Price:      domain.Price,
			QID:        domain.QID,
			QName:      domain.QName,
			StatusID:   domain.StatusID,
			StatusName: domain.StatusName,
		})
	}
	return result
}

func decodeQueueExtAttributes(queues []QueueExtAttributeEntry) []response.QueueExtAttributeEntry {
	if len(queues) == 0 {
		return nil
	}
	result := make([]response.QueueExtAttributeEntry, 0, len(queues))
	for _, entry := range queues {
		result = append(result, response.QueueExtAttributeEntry{
			ID:            entry.ID,
			TLD:           entry.TLD,
			AttributesXML: entry.Attributes.Raw,
		})
	}
	return result
}

func decodeQueueOrderDetails(entries []QueueOrderDetailEntry) []response.QueueOrderDetailEntry {
	if len(entries) == 0 {
		return nil
	}
	result := make([]response.QueueOrderDetailEntry, 0, len(entries))
	for _, entry := range entries {
		result = append(result, response.QueueOrderDetailEntry{
			OrderID:         entry.OrderID,
			OrderDate:       entry.OrderDate,
			PaidAmount:      entry.PaidAmount,
			DomainName:      entry.DomainName,
			OrderType:       entry.OrderType,
			ExtTime:         entry.ExtTime,
			OrderStatus:     entry.OrderStatus,
			RegistrationFee: entry.RegistrationFee,
			ApplicationFee:  entry.ApplicationFee,
			TotalFee:        entry.TotalFee,
		})
	}
	return result
}

func decodeQueueOrders(entries []QueueOrderEntry) []response.QueueOrderEntry {
	if len(entries) == 0 {
		return nil
	}
	result := make([]response.QueueOrderEntry, 0, len(entries))
	for _, entry := range entries {
		result = append(result, response.QueueOrderEntry{
			OrderID:    entry.OrderID,
			StatusID:   entry.StatusID,
			StatusName: entry.StatusName,
			StatusDesc: entry.StatusDesc,
			Success:    entry.Success,
			Failed:     entry.Failed,
			Pending:    entry.Pending,
			OrderDate:  entry.OrderDate,
		})
	}
	return result
}

func decodeQueueInfoEntries(entries []QueueInfoEntry) []response.QueueInfoEntry {
	if len(entries) == 0 {
		return nil
	}
	result := make([]response.QueueInfoEntry, 0, len(entries))
	for _, entry := range entries {
		result = append(result, response.QueueInfoEntry{
			TLD:              entry.TLD,
			NativeIDN:        entry.NativeIDN,
			CategoryName:     entry.CategoryName,
			CategoryDesc:     entry.CategoryDesc,
			TLDDescription:   entry.TLDDescription,
			TLDStatusID:      entry.TLDStatusID,
			TLDStatusDesc:    entry.TLDStatusDesc,
			QID:              entry.QID,
			QName:            entry.QName,
			QStatusID:        entry.QStatusID,
			QStatusName:      entry.QStatusName,
			ExtAttributes:    entry.ExtAttributes,
			StartDate:        entry.StartDate,
			EndDate:          entry.EndDate,
			AnnouncementDate: entry.AnnouncementDate,
			Price:            entry.Price,
			InEAP:            entry.InEAP,
		})
	}
	return result
}
