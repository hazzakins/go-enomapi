package internal

import (
	"encoding/xml"
	"fmt"

	"github.com/hazzakins/go-enomapi/response"
)

// AddBulkDomainsResponse is the raw API response for a bulk domain add request.
type AddBulkDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	AddBulkDomains AddBulkDomains `xml:"AddBulkDomains"`
	Success        string         `xml:"Success"`
	UseCart        string         `xml:"UseCart"`
}

// AddBulkDomains aggregates bulk domain add items and result metadata.
type AddBulkDomains struct {
	Items              []AddBulkDomainsItem `xml:"Item"`
	ListCount          int                  `xml:"ListCount"`
	CartErrors         int                  `xml:"CartErrors"`
	AllItemsSuccessful string               `xml:"AllItemsSuccessful"`
	CartItems          int                  `xml:"CartItems"`
}

// AddBulkDomainsItem captures an individual bulk domain add entry.
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

// CancelOrderResponse is the wire format for an order cancellation.
type CancelOrderResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Order CancelOrder `xml:"Order"`
}

// CancelOrder represents the cancellation details for an order.
type CancelOrder struct {
	OrderID       string             `xml:"OrderID"`
	Success       string             `xml:"Success"`
	Result        string             `xml:"Result"`
	Domains       []string           `xml:"Domains>DomainName"`
	DomainEntries []CancelOrderEntry `xml:"Domains>Domain"`
}

// CancelOrderEntry describes a cancelled domain and any error message.
type CancelOrderEntry struct {
	DomainName  string `xml:"DomainName"`
	Description string `xml:"Description"`
}

// GetConfirmationSettingsResponse holds confirmation email preferences.
type GetConfirmationSettingsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ConfirmationSettings ConfirmationSettings `xml:"ConfirmationSettings"`
}

// ConfirmationSettings describes reseller confirmation email settings.
type ConfirmationSettings struct {
	OrderConfirmation            string `xml:"OrderConfirmation"`
	TransferOrderConfirmation    string `xml:"TransferOrderConfirmation"`
	OrderConfirmationBCC         string `xml:"OrderConfirmationBCC"`
	TransferOrderConfirmationBCC string `xml:"TransferOrderConfirmationBCC"`
	EmailHead                    string `xml:"EmailHead"`
	EmailTail                    string `xml:"EmailTail"`
}

// PreconfigureResponse reports on TLD preconfiguration results.
type PreconfigureResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	PreConfigSuccess string `xml:"PreConfigSuccess"`
	Count            int    `xml:"Count"`
}

type DomainPurchase struct {
	Response
	OrderID           string             `xml:"OrderID"`
	OrderDelayed      bool               `xml:"OrderDelayed"`
	OrderStatus       string             `xml:"OrderStatus"`
	OrderDescription  string             `xml:"OrderDescription"`
	DomainInfo        DomainPurchaseInfo `xml:"Info"`
	TotalCharged      float64            `xml:"TotalCharged"`
	RegistrantPartyID string             `xml:"RegistrantPartyID"`
	IsRealTimeTLD     bool               `xml:"IsRealTimeTLD"`
}

type DomainPurchaseInfo struct {
	RegistryCreateDate string `xml:"RegistryCreateDate"`
	RegistryExpDate    string `xml:"RegistryExpDate"`
}

// QueueDomainPurchaseResponse captures a queued domain purchase result.
type QueueDomainPurchaseResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Purchase QueueDomainPurchase `xml:"Queue_DomainPurchase"`
}

// QueueDomainPurchase summarizes a queued purchase transaction.
type QueueDomainPurchase struct {
	Success              string `xml:"Success"`
	OrderID              string `xml:"OrderID"`
	Amount               string `xml:"Amount"`
	TransactionNumber    string `xml:"TransactionNumber"`
	TransactionReference string `xml:"TransactionReference"`
}

// QueueGetDomainsResponse contains queued domain entries and metrics.
type QueueGetDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Queue QueueGetDomains `xml:"Queue_GetDomains"`
}

// QueueGetDomains lists domains present in a queue response.
type QueueGetDomains struct {
	Metrics   *QueueMetrics `xml:"Metrics"`
	Domains   []QueueDomain `xml:"Domains>Domain"`
	ItemCount int           `xml:"ItemCount"`
	ItemTotal int           `xml:"ItemTotal"`
}

// QueueMetrics provides counts for queue and watchlist statistics.
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

// QueueDomain represents a queued domain along with pricing and status.
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

// QueueGetExtAttributesResponse wraps extended attributes returned in a queue.
type QueueGetExtAttributesResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Queues QueueExtAttributes `xml:"Queue_GetExtAttributes"`
}

// QueueExtAttributes groups queued extended attribute payloads.
type QueueExtAttributes struct {
	Queues []QueueExtAttributeEntry `xml:"Queues>Queue"`
}

// QueueExtAttributeEntry holds extended attribute data for a queued TLD.
type QueueExtAttributeEntry struct {
	ID         string `xml:"ID"`
	TLD        string `xml:"TLD"`
	Attributes RawXML `xml:"Attributes"`
}

// QueueGetOrderDetailResponse contains detailed order history for queues.
type QueueGetOrderDetailResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Detail QueueOrderDetail `xml:"Queue_GetOrderDetail"`
}

// QueueOrderDetail lists the returned order details for a queue query.
type QueueOrderDetail struct {
	Orders       []QueueOrderDetailEntry `xml:"Orders>Order"`
	RecordStart  int                     `xml:"RecordStart"`
	PagingSize   int                     `xml:"PagingSize"`
	NextRecord   int                     `xml:"NextRecord"`
	TotalResults int                     `xml:"TotalResults"`
}

// QueueOrderDetailEntry describes a single order returned in queue details.
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

// QueueGetOrdersResponse holds paged queue order summaries.
type QueueGetOrdersResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Orders QueueOrders `xml:"Queue_GetOrders"`
}

// QueueOrders encapsulates a collection of queue order summaries.
type QueueOrders struct {
	Orders    []QueueOrderEntry `xml:"Orders>Order"`
	ItemCount int               `xml:"ItemCount"`
	ItemTotal int               `xml:"ItemTotal"`
}

// QueueOrderEntry summarizes the status for an order in the queue.
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

// TMCheckResponse contains the lookup key for a trademark check.
type TMCheckResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	LookupKey string `xml:"LookupKey"`
}

// TMGetNoticeResponse wraps the trademark notice returned for a lookup.
type TMGetNoticeResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Notice TMNotice `xml:"TMNotice"`
}

// TMNotice contains trademark claims and metadata for a domain.
type TMNotice struct {
	TcnID        string `xml:"tcnID"`
	TcnStartDate string `xml:"tcnStartDate"`
	TcnExpDate   string `xml:"tcnExpDate"`
	SLD          string `xml:"Sld"`
	Claims       RawXML `xml:"Claims"`
}

// TMUpdateCartResponse indicates whether a trademark cart update succeeded.
type TMUpdateCartResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Success string `xml:"Success"`
}

// GetAgreementPageResponse returns the rendered agreement content.
type GetAgreementPageResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Content RawXML `xml:"content"`
}

// QueueGetInfoResponse provides queue metadata and available TLDs.
type QueueGetInfoResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Info QueueInfo `xml:"Queue_GetInfo"`
}

// QueueInfo aggregates queue-level details for preregistration.
type QueueInfo struct {
	Queues       []QueueInfoEntry `xml:"Queue"`
	TotalRecords int              `xml:"TotalRecords"`
}

// QueueInfoEntry describes a single queue entry and its properties.
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

// DeleteRegistrationResponse captures the result of deleting a domain.
type DeleteRegistrationResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DeleteDomain DeleteDomain `xml:"deletedomain"`
	ErrString    string       `xml:"ErrString"`
	ErrSource    string       `xml:"ErrSource"`
	ErrSection   string       `xml:"ErrSection"`
}

// DeleteDomain indicates whether the domain deletion was processed.
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

func (r *PreconfigureResponse) Decode() *response.Preconfigure {
	return &response.Preconfigure{
		PreConfigSuccess: r.PreConfigSuccess,
		Count:            r.Count,
		ResponseMeta:     decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (d *DomainPurchase) Decode() (*response.DomainPurchase, error) {
	if d.OrderID == "" {
		return nil, fmt.Errorf("no OrderID received - RRPCode: %d", d.ResponseCode)
	}
	result := response.DomainPurchase{
		OrderID:          d.OrderID,
		OrderCompleted:   d.DomainInfo.RegistryCreateDate != "",
		OrderStatus:      d.OrderStatus,
		OrderDescription: d.OrderDescription,
		Price:            d.TotalCharged,
	}
	var err error
	if d.DomainInfo.RegistryCreateDate != "" {
		result.RegistrationDate, err = ParseDate(d.DomainInfo.RegistryCreateDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RegistryCreateDate: %w", err)
		}
	}
	if d.DomainInfo.RegistryExpDate != "" {
		result.ExpirationDate, err = ParseDate(d.DomainInfo.RegistryExpDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RegistryExpDate: %w", err)
		}
	}
	return &result, nil
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
