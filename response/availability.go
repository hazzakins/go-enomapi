package response

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

type DomainProperties struct{}

type DomainEAP struct{}
