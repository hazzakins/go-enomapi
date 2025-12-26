package response

type ResponseMeta struct {
	Command        string
	Language       string
	ErrCount       int
	RRPCode        int
	RRPText        string
	APIType        string
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

type AdvancedDomainSearch struct {
	SearchParams  AdvancedDomainSearchParams
	TotalResults  int
	StartPosition int
	NextPosition  int
	MultiRRP      bool
	TLDOverride   string
	Domains       []AdvancedDomainSearchDomain
	ResponseMeta  ResponseMeta
}

type AdvancedDomainSearchParams struct {
	TLDList             string
	SLD                 string
	SearchCriteria      string
	ParkingStatus       string
	RegistrationStatus  string
	AutoRenew           string
	Locked              string
	DaysTillExpires     string
	DaysExpired         string
	NSStatus            string
	NameServer          string
	HasIDProtect        string
	HasPOPMail          string
	EmailForwarding     string
	HasWebHosting       string
	ExcludeNumbers      string
	ExcludeDashes       string
	IncludeSubAccounts  string
	SubAccountLogin     string
	RecordsToReturn     string
	StartPosition       string
	OrderBy             string
	FolderOption        string
	FolderSyncStatus    string
	FolderName          string
	CreationDate        string
	DaysUntilIDPExpires string
	WBLStatusID         string
	WBLEnabled          string
	WBLAutoRenew        string
	ContactXML          string
	XMLResponse         string
	HostRecordType      string
	HostName            string
	HostAddress         string
	EmailResultsOnly    string
	XML                 string
}

type AdvancedDomainSearchDomain struct {
	DomainNameID               int64
	SLD                        string
	TLD                        string
	AutoRenew                  string
	ExpDate                    string
	DomainRegistrationStatus   string
	DeleteType                 string
	LoginID                    string
	AccountID                  string
	NSStatus                   string
	FolderStatus               string
	RRProcessor                string
	RRCompanyName              string
	HasIDProtect               string
	IDProtectExpires           string
	DomainFolderStatus         string
	AbleToReactivate           string
	IsPremiumName              string
	PremiumPrice               string
	PremiumAboveThresholdPrice string
	PremiumCategory            string
	ReactivatePrice            string
	WBLStatusID                string
	WBLStatus                  string
	WBLExpDate                 string
	WBLAutoRenew               string
	NameServers                string
	Vas                        string
}

type GetAllDomains struct {
	Domains           []GetAllDomainsDetail
	DomainCount       int
	UserRequestStatus string
	ResponseMeta      ResponseMeta
}

type GetAllDomainsDetail struct {
	DomainName     string
	DomainNameID   int64
	ExpirationDate string
	LockStatus     string
	AutoRenew      string
}

type GetDomainCount struct {
	RegisteredCount     int
	HostCount           int
	ExpiringCount       int
	ExpiredDomainsCount int
	RGP                 int
	ExtendedRGP         int
	KeywordCount        int
	ProcessCount        int
	WatchlistCount      int
	CartItemCount       int
	TrafficMsg          *DomainTraffic
	ResponseMeta        ResponseMeta
}

type GetDomainExp struct {
	ExpirationDate string
	ResponseMeta   ResponseMeta
}

type DomainTraffic struct {
	VistaCustomer  bool
	RedirectorData bool
	Month          string
	PageViews      string
	Visitors       string
	FreeTrial      bool
	PDQVista       string
}

type GetDomainNameID struct {
	DomainRRP    string
	SLD          string
	TLD          string
	DomainNameID int64
	ResponseMeta ResponseMeta
}

type GetDomains struct {
	Tab              string
	MultiRRP         bool
	DomainListType   string
	Domains          []GetDomainsItem
	EndPosition      int
	PreviousRecords  int
	NextRecords      int
	OrderBy          string
	Result           bool
	StartPosition    int
	DomainCount      int
	TotalDomainCount int
	StartLetter      string
	ResponseMeta     ResponseMeta
}

type GetDomainsItem struct {
	DomainNameID   int64
	SLD            string
	TLD            string
	NSStatus       string
	ExpirationDate string
	AutoRenew      string
	WPPSStatus     string
	WPPSExpDate    string
	RRProcessor    string
}

type GetDomainStatus struct {
	DomainName   string
	Registrar    string
	InAccount    string
	StatusDesc   string
	ExpDate      string
	OrderID      string
	ResponseMeta ResponseMeta
}

type GetDomainSldTld struct {
	DomainRRP    string
	SLD          string
	TLD          string
	DomainNameID int64
	ResponseMeta ResponseMeta
}

type GetExpiredDomains struct {
	Domains      []ExpiredDomainDetail
	DomainCount  int
	ResponseMeta ResponseMeta
}

type ExpiredDomainDetail struct {
	DomainName     string
	DomainNameID   int64
	Status         string
	ExpirationDate string
	LockStatus     string
}

type GetExtendInfo struct {
	RegistrarHold       bool
	Expiration          string
	MaxExtension        int
	MinAllowed          int
	CCAuthorized        bool
	Price               string
	Balance             string
	AvailableBalance    string
	CustomerPrefs       ExtendInfoCustomerPrefs
	CustomerInformation ExtendInfoCustomerInformation
	ResponseMeta        ResponseMeta
}

type ExtendInfoCustomerPrefs struct {
	DefPeriod            int
	AllowDNS             bool
	ShowPopups           bool
	AutoRenew            bool
	RegLock              bool
	AutoPakRenew         bool
	UseDNS               bool
	ResellerStatus       string
	RenewalSetting       int
	RenewalBCC           int
	RenewalURLForward    bool
	RenewalEmailForward  bool
	MailNumLimit         int
	IDProtect            bool
	DefIDProtectRenew    bool
	DefWBLRenew          bool
	NameJetSales         bool
	DefaultHostRecords   ExtendInfoDefaultHostRecords
	DefaultHostRecordOwn bool
	UseOurDNS            bool
	NameServers          ExtendInfoNameServers
}

type ExtendInfoDefaultHostRecords struct {
	HostRecords []ExtendInfoHostRecord
}

type ExtendInfoHostRecord struct {
	HostName   string
	Address    string
	RecordType string
}

type ExtendInfoNameServers struct {
	DNS1 string
	DNS2 string
	DNS3 string
	DNS4 string
	DNS5 string
}

type ExtendInfoCustomerInformation struct {
	AcceptTerms   bool
	URL           string
	ParentAccount string
	ParentLogin   string
	NoService     bool
	BulkRegLimit  int
	Account       string
}

type GetNews struct {
	Alerts       Alerts
	ResponseMeta ResponseMeta
}

type GetProductNews struct {
	Alerts       Alerts
	ResponseMeta ResponseMeta
}

type Alerts struct {
	Entries []Alert
	Total   int
}

type Alert struct {
	Product string
	Message string
}

type GetPasswordBit struct {
	DomainPassword string
	PasswordSet    int
	ResponseMeta   ResponseMeta
}

type GetRegistrationStatus struct {
	RegistrationStatus string
	PurchaseStatus     string
	ResponseMeta       ResponseMeta
}

type GetRegLock struct {
	RegLock      string
	Registrar    string
	ResponseMeta ResponseMeta
}

type GetRenew struct {
	RenewName         bool
	PakExist          bool
	AutoPakRenew      bool
	EmailFwdExists    bool
	EmailForwardRenew bool
	URLFwdExists      bool
	URLForwardRenew   bool
	IDProtectRenew    bool
	IDProtectExists   bool
	MobilizerRenew    bool
	ResponseMeta      ResponseMeta
}

type GetSubAccountPassword struct {
	ResponseMeta ResponseMeta
}

type ParseDomain struct {
	Host         string
	SLD          string
	TLD          string
	ResponseMeta ResponseMeta
}

type PortalGetAwardedDomains struct {
	Domains           []PortalAwardedDomain
	DomainCount       int
	TotalDomainCount  int
	NextStartPosition int
	Success           string
	ResponseMeta      ResponseMeta
}

type PortalAwardedDomain struct {
	DomainName              string
	EmailAddress            string
	ExpirationDate          string
	RegisterDate            string
	ForeignLoginID          string
	PortalDomainID          string
	RegisterPrice           string
	RenewPrice              string
	RegistrationPeriod      string
	ResellerProvisioned     string
	ResellerProvisionedDate string
	IsPremium               string
	RegistrationStatus      string
}

type PortalGetDomainInfo struct {
	Domain          PortalDomainInfo
	TransactionsXML string
	ServicesXML     string
	RAASettingsXML  string
	PaymentXML      string
	ContactsXML     string
	ResponseMeta    ResponseMeta
}

type PortalDomainInfo struct {
	DomainName              string
	EmailAddress            string
	ExpirationDate          string
	RegisterDate            string
	ForeignLoginID          string
	PortalDomainID          string
	RegisterPrice           string
	RenewPrice              string
	RegistrationPeriod      string
	ResellerProvisioned     string
	ResellerProvisionedDate string
	IsPremium               string
	RegistrationStatus      string
}

type PortalGetToken struct {
	Token        string
	ResponseMeta ResponseMeta
}

type PortalUpdateAwardedDomains struct {
	Success      string
	ResponseMeta ResponseMeta
}

type RPTGetReport struct {
	Report       Report
	ResponseMeta ResponseMeta
}

type Report struct {
	RptString             string
	ThreeMoPast           string
	SixMoPast             string
	BeginDate             string
	EndDate               string
	Options               ReportOptions
	ReportType            string
	ReportName            string
	QueueReportResults    string
	ResultsXML            string
	NewStartPosition      string
	ThisStartPosition     string
	ThisEndPosition       string
	PreviousStartPosition string
	TotalReturned         int
	TotalRows             int
	ShowPrevious          string
	ShowNext              string
	ShowPaging            string
	Version               string
}

type ReportOptions struct {
	ReportTypes []ReportTypeOption
}

type ReportTypeOption struct {
	ID   string
	Name string
}

type SetPassword struct {
	ResponseMeta ResponseMeta
}

type SetRegLock struct {
	RegLock      string
	RRPCodeSR    string
	RRPText      string
	ResponseMeta ResponseMeta
}

type SetRenew struct {
	RenewName         bool
	AutoPakRenew      bool
	EmailForwardRenew bool
	URLForwardRenew   bool
	WPPSRenew         bool
	ResponseMeta      ResponseMeta
}

type StatusDomain struct {
	DomainName   string
	Registrar    string
	InAccount    string
	ExpDate      string
	OrderID      string
	ResponseMeta ResponseMeta
}

type ValidatePassword struct {
	ResponseMeta ResponseMeta
}
