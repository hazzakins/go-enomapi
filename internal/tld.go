package internal

import (
	"encoding/xml"
	"fmt"

	"github.com/hazzakins/go-enomapi/response"
)

type TLDListResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	TLDList  []string `xml:"tldlist>tld>tld"`
	TLDCount int32    `xml:"tldlist>tldcount"`
}

func (t *TLDListResponse) Decode() response.TLDList {
	result := make(response.TLDList, t.TLDCount)
	copy(result, t.TLDList)
	return result
}

type TLDDetailsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	Tlds TLDDetail `xml:"tlds>tld"`
}

// GetExtAttributesResponse wraps the extended attributes metadata.
type GetExtAttributesResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Attributes ExtAttributes `xml:"Attributes"`
}

// ExtAttributes lists the required or optional extended attributes.
type ExtAttributes struct {
	Items []ExtAttribute `xml:"Attribute"`
}

// ExtAttribute defines a single extended attribute and its options.
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

// ExtAttributeOption is a selectable option for an extended attribute.
type ExtAttributeOption struct {
	ID          int    `xml:"ID"`
	Value       string `xml:"Value"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

// GetIDNCodesResponse contains language codes supported for IDN domains.
type GetIDNCodesResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	TLDs IDNCodesTLDs `xml:"tlds"`
}

// IDNCodesTLDs groups IDN language support by TLD.
type IDNCodesTLDs struct {
	TLDs []IDNCodesTLD `xml:"tld"`
}

// IDNCodesTLD lists the IDN languages available for a TLD.
type IDNCodesTLD struct {
	TLD       string        `xml:"tld,attr"`
	Languages []IDNLanguage `xml:"language"`
}

// IDNLanguage represents an IDN language code and name.
type IDNLanguage struct {
	Code string `xml:"code,attr"`
	Name string `xml:"name,attr"`
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
