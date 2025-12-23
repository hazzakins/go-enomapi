package response

type GetDomainInfo struct {
	DomainName     DomainInfoDomainName
	MultyLangSLD   bool
	Status         DomainInfoStatus
	ParkingEnabled bool
	Services       []DomainServiceEntry
	Command        string
	APIType        string
	Language       string
	ErrCount       int
	ResponseCount  int
	MinPeriod      int
	MaxPeriod      int
	Server         string
	Site           string
	IsLockable     bool
	IsRealTimeTLD  bool
	TimeDifference string
	ExecTime       string
	Done           bool
	TrackingKey    string
	RequestDate    string
}

type DomainInfoDomainName struct {
	SLD          string
	TLD          string
	DomainNameID int64
	Name         string
}

type DomainInfoStatus struct {
	Expiration            string
	DeleteByDate          string
	DeleteType            string
	Restorable            bool
	RenewBeforeExpiration string
	Registrar             string
	RegistrationStatus    string
	PurchaseStatus        string
	BelongsTo             DomainBelongsTo
	EscrowHold            bool
	EscrowLiftDate        string
	AuctionHold           bool
	AuctionLiftDate       string
}

type DomainBelongsTo struct {
	PartyID string
	Value   string
}

type DomainServiceEntry struct {
	Name           string
	EnomDNS        *DomainServiceEnomDNS
	Service        *DomainService
	Configuration  *DomainServiceConfig
	WBL            *DomainServiceWBL
	Mobilizer      *DomainServiceMobilizer
	RAASetting     *DomainServiceRAASetting
	IRTPSetting    *DomainServiceIRTPSetting
	WhoisPublicity *DomainServiceWhoisPublic
}

type DomainServiceEnomDNS struct {
	Value     string
	IsDotName string
}

type DomainService struct {
	Changable bool
	Value     string
}

type DomainServiceConfig struct {
	Changable    bool
	Type         string
	DNS          []string
	Hosts        []DomainServiceHost
	WSB          string
	SiteID       string
	ProdType     string
	NextBillDate string
	WPPS         *DomainServiceWPPS
}

type DomainServiceHost struct {
	Name       string
	Type       string
	Address    string
	MXPref     string
	IsEditable bool
}

type DomainServiceWPPS struct {
	CloakedEmail string
	ForwardTo    string
	ExpireDate   string
	AutoRenew    string
}

type DomainServiceWBL struct {
	WBLID              string
	StatusID           string
	StatusDescr        string
	ExpDate            string
	Enabled            bool
	Renew              bool
	CompanyName        string
	CompanyDescription string
	DomainName         string
	Street             string
	City               string
	PostalCode         string
	Country            string
	CategoryIDX        string
	Fields             []DomainServiceField
	FieldName          string
	Value              string
}

type DomainServiceField struct {
	FieldName string
	Value     string
}

type DomainServiceMobilizer struct{}

type DomainServiceRAASetting struct {
	VerificationStatus string
	DomainSuspended    bool
	SuspensionDate     string
	IsQueuedChange     bool
	StatusExpDate      string
}

type DomainServiceIRTPSetting struct {
	ICANNCompliant      bool
	OptOut              bool
	TransferLock        bool
	TransferLockExpDate DomainTransferLockExpDate
}

type DomainTransferLockExpDate struct {
	DaysRemaining int
	UTC           string
	Epoch         int64
	Value         string
}

type DomainServiceWhoisPublic struct {
	VASItemID int64
	Enabled   bool
}
