package internal

import (
	"fmt"

	"github.com/hazzakins/go-enomapi/response"
)

type DomainCheckResponse struct {
	Response
	Name       string                `xml:"Domains>Domain>Name"`
	RRPCode    int32                 `xml:"Domains>Domain>RRPCode"`
	RRPText    string                `xml:"Domains>Domain>RRPText"`
	IsPremium  bool                  `xml:"Domains>Domain>IsPremium"`
	IsPlatinum bool                  `xml:"Domains>Domain>IsPlatinum"`
	IsEAP      bool                  `xml:"Domains>Domain>IsEAP"`
	Prices     DomainCheckPrices     `xml:"Domains>Domain>Prices"`
	Properties DomainCheckProperties `xml:"Domains>Domain>Properties"`
	EAP        DomainCheckEAP        `xml:"Domains>Domain>EAP"`
}

type DomainCheckPrices struct {
	Currency      string  `xml:"Currency"`
	Registration  float64 `xml:"Registration"`
	Renewal       float64 `xml:"Renewal"`
	Restore       float64 `xml:"Restore"`
	Transfer      float64 `xml:"Transfer"`
	ExpectedPrice float64 `xml:"ExpectedCustomerSuppliedPrice"`
}

type DomainCheckProperties struct {
	NativeSLD      string `xml:"NativeSLD"`
	MinRegYear     int    `xml:"MinRegYear"`
	MaxRegYear     int    `xml:"MaxRegYear"`
	AbleToLock     bool   `xml:"AbleToLock"`
	ExtAttributes  bool   `xml:"ExtAttributes"`
	Transferable   bool   `xml:"Transferable"`
	AllowWPPS      bool   `xml:"AllowWPPS"`
	TrademarkStart string `xml:"TrademarkStart"`
	TrademarkEnd   string `xml:"TrademarkEnd"`
}

type DomainCheckEAP struct {
	CurrentDay int     `xml:"CurrentDay"`
	Day        int     `xml:"Day"`
	Fee        float64 `xml:"Fee"`
	Open       bool    `xml:"Open"`
	DateStart  string  `xml:"DateStart"`
	DateEnd    string  `xml:"DateEnd"`
}

func (r *DomainCheckResponse) Decode() (*response.DomainCheck, error) {
	if r.Prices.Currency != "" {
		return nil, fmt.Errorf("unexpected currency: %s", r.Prices.Currency)
	}
	return &response.DomainCheck{
		FQDN:        r.Name,
		IsAvailable: r.RRPCode == 210, //TODO remove magic number, see response.go
		IsPremium:   r.IsPremium,
		IsPlatinum:  r.IsPlatinum,
		IsEAP:       r.IsEAP,
		Prices: response.DomainPrices{
			Registration:  r.Prices.Registration,
			Renewal:       r.Prices.Renewal,
			Restore:       r.Prices.Restore,
			Transfer:      r.Prices.Transfer,
			ExpectedPrice: r.Prices.ExpectedPrice,
		},
		Properties: response.DomainProperties{},
		EAP:        response.DomainEAP{},
	}, nil
}
