package internal

import (
	"encoding/xml"
	"fmt"

	"github.com/hazzakins/go-enomapi/response"
)

type TLDDetailsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	Tlds TLDDetail `xml:"tlds>tld"`
}

type TLDDetail struct {
	TLD                     string                `xml:"TLD"`
	AbleToLock              bool                  `xml:"AbleToLock"`
	AllowWPPS               bool                  `xml:"AllowWPPS"`
	AutoRenewOnly           bool                  `xml:"AutoRenewOnly"`
	ExtAttributes           bool                  `xml:"ExtAttributes"`
	AllowWBL                bool                  `xml:"AllowWBL"`
	TransRequiresNewContact bool                  `xml:"TransRequiresNewContact"`
	EncodingType            string                `xml:"EncodingType"`
	ClearPhoneNumber        bool                  `xml:"ClearPhoneNumber"`
	ClearFax                bool                  `xml:"ClearFax"`
	ValidateDNSHosting      bool                  `xml:"ValidateDNSHosting"`
	SupportsDNSSec          bool                  `xml:"SupportsDnsSec"`
	AllowWhoisPublicity     bool                  `xml:"AllowWhoisPublicity"`
	Registration            TLDDetailRegistration `xml:"Registration"`
	Renewal                 TLDDetailRenewal      `xml:"Renewal"`
	Transfer                TLDDetailTransfer     `xml:"Transfer"`
	PremiumNames            TLDDetailPremiumNames `xml:"PremiumNames"`
	NameServers             TLDDetailNameServers  `xml:"NameServers"`
}

type TLDDetailRegistration struct {
	Realtime                     bool   `xml:"Realtime"`
	Unit                         string `xml:"Unit"`
	Minimum                      int    `xml:"Minimum"`
	Maximum                      int    `xml:"Maximum"`
	HasPremiumNames              bool   `xml:"HasPremiumNames"`
	ExtAttributes                bool   `xml:"ExtAttributes"`
	ExtAttributesDomainLevel     bool   `xml:"ExtAttributesDomainLevel"`
	DNSRequired                  bool   `xml:"DNSRequired"`
	DNSMinimum                   int    `xml:"DNSMinimum"`
	DNSMaximum                   int    `xml:"DNSMaximum"`
	GeneralAvailabilityStartDate string `xml:"GeneralAvailabilityStartDate"`
}

type TLDDetailRenewal struct {
	AutoRenewOnly        bool `xml:"AutoRenewOnly"`
	AutoRenewed          bool `xml:"AutoRenewed"`
	RenewBeforeExpMonths int  `xml:"RenewBeforeExpMonths"`
	DeleteType           int  `xml:"DeleteType"`
	DeleteDay            int  `xml:"DeleteDay"`
	GracePeriod          int  `xml:"GracePeriod"`
	Reactivate           bool `xml:"Reactivate"`
	Restorable           bool `xml:"Restorable"`
	RGP                  bool `xml:"RGP"`
	RGPDays              int  `xml:"RGPDays"`
	ExtendedRGP          bool `xml:"ExtendedRGP"`
	TransferPeriod       int  `xml:"TransferPeriod"`
}

type TLDDetailTransfer struct {
	Transferable       bool `xml:"Transferable"`
	AuthInfo           bool `xml:"AuthInfo"`
	Realtime           bool `xml:"Realtime"`
	AutoVerification   bool `xml:"AutoVerification"`
	AutoFax            bool `xml:"AutoFax"`
	TransferByFOA      bool `xml:"TransferByFOA"`
	RequiresNewContact bool `xml:"RequiresNewContact"`
}

type TLDDetailPremiumNames struct {
	HasPremiumNames    bool   `xml:"HasPremiumNames"`
	MaxPremiumRegYears int    `xml:"MaxPremiumRegYears"`
	TrademarkStartDate string `xml:"TrademarkStartDate"`
	TrademarkEndDate   string `xml:"TrademarkEndDate"`
}

type TLDDetailNameServers struct {
	NameServers    bool `xml:"NameServers"`
	MinNameServers int  `xml:"MinNameServers"`
	MaxNameServers int  `xml:"MaxNameServers"`
}

func (t *TLDDetailsResponse) Decode() (*response.TLDDetails, error) {
	result := response.TLDDetails{
		TLD:      t.Tlds.TLD,
		Lockable: t.Tlds.AbleToLock,
		Registration: response.TLDRegistrationDetails{
			IsRealTime:    t.Tlds.Registration.Realtime,
			DNSRequired:   t.Tlds.Registration.DNSRequired,
			DNSMinServers: t.Tlds.Registration.DNSMinimum,
			DNSMaxServers: t.Tlds.Registration.DNSMaximum,
		},
	}
	switch t.Tlds.EncodingType {
	case "PUNY":
		result.SupportsIDN = true
	case "":
		result.SupportsIDN = false
	default:
		return nil, fmt.Errorf("unexpected encoding in TLD: %s", t.Tlds.EncodingType)
	}
	return &result, nil
}
