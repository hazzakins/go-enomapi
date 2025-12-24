package response

import "time"

type DomainCheck struct {
	FQDN        string
	IsAvailable bool
	IsPremium   bool
	IsPlatinum  bool
	IsEAP       bool
	Prices      DomainPrices
	Properties  DomainProperties
	EAP         DomainEAP
}

type DomainPrices struct {
	Registration  float64
	Renewal       float64
	Restore       float64
	Transfer      float64
	ExpectedPrice float64
}

type DomainInfo struct {
	Domain         string
	Name           string
	Extension      string
	IsRegistered   bool
	ExpirationDate *time.Time
}

type DomainPurchase struct {
	// If this is not "" then the order was placed successfully
	OrderID string
	// If this is false the order was queued. Check OrderStatus and OrderDescription.
	OrderCompleted   bool
	OrderStatus      string
	OrderDescription string
	// Only set if OrderCompleted == true
	RegistrationDate time.Time
	// Only set if OrderCompleted == true
	ExpirationDate time.Time
	// The total amount charged for this purchase
	Price float64
}

type TLDList []string

type TLDRegistrationDetails struct {
	IsRealTime    bool
	DNSRequired   bool
	DNSMinServers int
	DNSMaxServers int
}

type TLDDetails struct {
	TLD          string
	Lockable     bool
	SupportsIDN  bool
	Registration TLDRegistrationDetails
}

type NameSpinner struct {
	SpinCount    int
	TLDList      string
	OriginalSLD  string
	Domains      []SpinnerDomain
	ResponseMeta ResponseMeta
}

type SpinnerDomain struct {
	Name     string
	Com      string
	ComScore string
	Net      string
	NetScore string
	Tv       string
	TvScore  string
	Cc       string
	CcScore  string
}

type DomainProperties struct{}

type DomainEAP struct{}
