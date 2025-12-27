package response

type CancelOrder struct {
	OrderID      string
	Success      string
	Result       string
	Domains      []CancelOrderDomain
	ResponseMeta ResponseMeta
}

type CancelOrderDomain struct {
	DomainName  string
	Description string
}

type GetConfirmationSettings struct {
	OrderConfirmation            string
	TransferOrderConfirmation    string
	OrderConfirmationBCC         string
	TransferOrderConfirmationBCC string
	EmailHead                    string
	EmailTail                    string
	ResponseMeta                 ResponseMeta
}

type QueueDomainPurchase struct {
	Success              string
	OrderID              string
	Amount               string
	TransactionNumber    string
	TransactionReference string
	ResponseMeta         ResponseMeta
}

type QueueGetDomains struct {
	Metrics      *QueueMetrics
	Domains      []QueueDomain
	ItemCount    int
	ItemTotal    int
	ResponseMeta ResponseMeta
}

type QueueMetrics struct {
	WatchlistTotal         int
	WatchlistPreorder      int
	WatchlistDomainWatched int
	WatchlistTLDWatched    int
	NewTLDOrders           int
	NewTLDValidated        int
	NewTLDPending          int
	NewTLDRejected         int
	QueuesDisabled         string
	PreregDisabled         string
}

type QueueDomain struct {
	OrderID    string
	DomainName string
	TLDID      string
	TLDName    string
	TLDDesc    string
	DateStart  string
	DateEnd    string
	Price      string
	QID        string
	QName      string
	StatusID   string
	StatusName string
}

type QueueGetExtAttributes struct {
	Queues       []QueueExtAttributeEntry
	ResponseMeta ResponseMeta
}

type QueueExtAttributeEntry struct {
	ID            string
	TLD           string
	AttributesXML string
}

type QueueGetOrderDetail struct {
	Orders       []QueueOrderDetailEntry
	RecordStart  int
	PagingSize   int
	NextRecord   int
	TotalResults int
	ResponseMeta ResponseMeta
}

type QueueOrderDetailEntry struct {
	OrderID         string
	OrderDate       string
	PaidAmount      string
	DomainName      string
	OrderType       string
	ExtTime         string
	OrderStatus     string
	RegistrationFee string
	ApplicationFee  string
	TotalFee        string
}

type QueueGetOrders struct {
	Orders       []QueueOrderEntry
	ItemCount    int
	ItemTotal    int
	ResponseMeta ResponseMeta
}

type QueueOrderEntry struct {
	OrderID    string
	StatusID   string
	StatusName string
	StatusDesc string
	Success    string
	Failed     string
	Pending    string
	OrderDate  string
}

type GetAgreementPage struct {
	ContentXML   string
	ResponseMeta ResponseMeta
}

type QueueGetInfo struct {
	Queues       []QueueInfoEntry
	TotalRecords int
	ResponseMeta ResponseMeta
}

type QueueInfoEntry struct {
	TLD              string
	NativeIDN        string
	CategoryName     string
	CategoryDesc     string
	TLDDescription   string
	TLDStatusID      string
	TLDStatusDesc    string
	QID              string
	QName            string
	QStatusID        string
	QStatusName      string
	ExtAttributes    string
	StartDate        string
	EndDate          string
	AnnouncementDate string
	Price            string
	InEAP            string
}
