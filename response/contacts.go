package response

type Contacts struct {
	ConsentStatus  string
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

type GetContacts struct {
	DomainName           DomainName
	Registrant           ContactData
	AuxBilling           ContactData
	Tech                 ContactData
	Admin                ContactData
	Billing              ContactData
	ICANNCompliant       bool
	PendingVerification  bool
	RAAStatus            bool
	RAAStatusDesc        string
	IRTPOptOut           bool
	TransferLock         bool
	TransferLockExpDate  DomainTransferLockExpDate
	Nexus                Nexus
	Purpose              string
	CurrentAttributes    string
	WPPSAllowed          bool
	WPPSExists           bool
	WPPSEnabled          bool
	WPPSExpDate          string
	WPPSAutoRenew        string
	WPPSContactData      ContactData
	EscrowLiftDate       string
	EscrowHold           bool
	AttributesXML        string
	IsContactShared      bool
	ContactRestrictedTLD bool
	WhoisPublicity       *WhoisPublicity
	Command              string
	APIType              string
	Language             string
	ErrCount             int
	ResponseCount        int
	MinPeriod            int
	MaxPeriod            int
	Server               string
	Site                 string
	IsLockable           bool
	IsRealTimeTLD        bool
	TimeDifference       string
	ExecTime             string
	Done                 bool
	TrackingKey          string
	RequestDate          string
}

type GetWhoisContact struct {
	DomainName         DomainName
	DomainExpired      bool
	AboutUsLink        string
	Row                WhoisRow
	Contacts           []ContactEntry
	RRPInfo            WhoisRRPInfo
	BusinessListingXML string
	Legal              string
	WPPSEnabled        bool
	Version            int
	Success            bool
	Command            string
	APIType            string
	Language           string
	ErrCount           int
	ResponseCount      int
	MinPeriod          int
	MaxPeriod          int
	Server             string
	Site               string
	IsLockable         bool
	IsRealTimeTLD      bool
	TimeDifference     string
	ExecTime           string
	Done               bool
	TrackingKey        string
	RequestDate        string
}

type GetWPPSInfo struct {
	DomainName     DomainName
	WPPSAllowed    bool
	WPPSExists     bool
	WPPSEnabled    bool
	WPPSExpDate    string
	WPPSAutoRenew  string
	WPPSPrice      string
	Contacts       []ContactEntry
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

type DomainName struct {
	SLD          string
	TLD          string
	DomainNameID int64
	Name         string
}

type ContactData struct {
	Organization        string
	FirstName           string
	LastName            string
	JobTitle            string
	Address1            string
	Address2            string
	City                string
	StateProvinceChoice string
	StateProvince       string
	PostalCode          string
	Country             string
	EmailAddress        string
	Phone               string
	PhoneExt            string
	Fax                 string
	ConsentStatus       string
}

type ContactEntry struct {
	ContactType string
	ContactData ContactData
}

type Nexus struct {
	Category string
	Value    string
}

type WhoisPublicity struct {
	VASItemID      int64
	ProdStatusID   int64
	ProdStatusDesc string
	ProdEnabled    bool
	ProdConsented  bool
	AutoRenew      bool
	ExpDate        string
}

type WhoisRow struct {
	RecordFound              string
	IsExpired                string
	RRProcessor              string
	Registrar                string
	LockStatus               string
	AbuseContactEmail        string
	AbuseContactPhone        string
	DomainRoid               string
	ExpDate                  string
	CreateDate               string
	LastUpdatedDate          string
	NSStatus                 string
	DnsSec                   string
	DnsSecStatus             string
	ResellerOrganizationName string
	ResellerURL              string
	AboutUsLink              string
	ResellerEmailAddress     string
	RCOMPitch                string
	AbuseURL                 string
	IanaID                   string
	RegistrarWhoisServer     string
	RegistrarURL             string
	Idp                      string
	Wps                      string
}

type WhoisRRPInfo struct {
	RegistrationExpirationDate string
	CreatedDate                string
	UpdatedDate                string
	Status                     []string
	Domain                     string
	NameServers                []string
}
