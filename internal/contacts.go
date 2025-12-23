package internal

import (
	"encoding/xml"
	"io"
	"strings"

	"github.com/hazzakins/go-enomapi/response"
)

type ContactsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ConsentStatus  string `xml:"ConsentStatus"`
	APIType        string `xml:"APIType"`
	ResponseCount  int    `xml:"ResponseCount"`
	MinPeriod      int    `xml:"MinPeriod"`
	MaxPeriod      int    `xml:"MaxPeriod"`
	Server         string `xml:"Server"`
	Site           string `xml:"Site"`
	IsLockable     bool   `xml:"IsLockable"`
	IsRealTimeTLD  bool   `xml:"IsRealTimeTLD"`
	TimeDifference string `xml:"TimeDifference"`
	ExecTime       string `xml:"ExecTime"`
	Done           bool   `xml:"Done"`
	TrackingKey    string `xml:"TrackingKey"`
	RequestDate    string `xml:"RequestDateTime"`
}

type GetContactsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	GetContacts    GetContacts `xml:"GetContacts"`
	APIType        string      `xml:"APIType"`
	ResponseCount  int         `xml:"ResponseCount"`
	MinPeriod      int         `xml:"MinPeriod"`
	MaxPeriod      int         `xml:"MaxPeriod"`
	Server         string      `xml:"Server"`
	Site           string      `xml:"Site"`
	IsLockable     bool        `xml:"IsLockable"`
	IsRealTimeTLD  bool        `xml:"IsRealTimeTLD"`
	TimeDifference string      `xml:"TimeDifference"`
	ExecTime       string      `xml:"ExecTime"`
	Done           bool        `xml:"Done"`
	TrackingKey    string      `xml:"TrackingKey"`
	RequestDate    string      `xml:"RequestDateTime"`
}

type GetContacts struct {
	DomainName           DomainName                `xml:"domainname"`
	Registrant           ContactData               `xml:"Registrant"`
	AuxBilling           ContactData               `xml:"AuxBilling"`
	Tech                 ContactData               `xml:"Tech"`
	Admin                ContactData               `xml:"Admin"`
	Billing              ContactData               `xml:"Billing"`
	ICANNCompliant       bool                      `xml:"ICANNCompliant"`
	PendingVerification  bool                      `xml:"PendingVerification"`
	RAAStatus            bool                      `xml:"RAA_Status"`
	RAAStatusDesc        string                    `xml:"RAA_StatusDesc"`
	IRTPOptOut           bool                      `xml:"IRTPOptOut"`
	TransferLock         bool                      `xml:"TransferLock"`
	TransferLockExpDate  DomainTransferLockExpDate `xml:"TransferLockExpDate"`
	Nexus                Nexus                     `xml:"Nexus"`
	Purpose              string                    `xml:"Purpose"`
	CurrentAttributes    string                    `xml:"CurrentAttributes"`
	WPPSAllowed          bool                      `xml:"WPPSAllowed"`
	WPPSExists           bool                      `xml:"WPPSExists"`
	WPPSEnabled          bool                      `xml:"WPPSEnabled"`
	WPPSExpDate          string                    `xml:"WPPSExpDate"`
	WPPSAutoRenew        string                    `xml:"WPPSAutoRenew"`
	WPPSContactData      ContactData               `xml:"WPPSContactData"`
	EscrowLiftDate       string                    `xml:"escrowliftdate"`
	EscrowHold           bool                      `xml:"escrowhold"`
	Attributes           ContactAttributes         `xml:"Attributes"`
	IsContactShared      bool                      `xml:"IsContactShared"`
	ContactRestrictedTLD bool                      `xml:"ContactRestrictedTLD"`
	WhoisPublicity       *WhoisPublicity           `xml:"WhoisPublicity"`
}

type GetWhoisContactResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	GetWhoisContacts GetWhoisContacts `xml:"GetWhoisContacts"`
	Success          bool             `xml:"Success"`
	APIType          string           `xml:"APIType"`
	ResponseCount    int              `xml:"ResponseCount"`
	MinPeriod        int              `xml:"MinPeriod"`
	MaxPeriod        int              `xml:"MaxPeriod"`
	Server           string           `xml:"Server"`
	Site             string           `xml:"Site"`
	IsLockable       bool             `xml:"IsLockable"`
	IsRealTimeTLD    bool             `xml:"IsRealTimeTLD"`
	TimeDifference   string           `xml:"TimeDifference"`
	ExecTime         string           `xml:"ExecTime"`
	Done             bool             `xml:"Done"`
	TrackingKey      string           `xml:"TrackingKey"`
	RequestDate      string           `xml:"RequestDateTime"`
}

type GetWhoisContacts struct {
	DomainName      DomainName        `xml:"domainname"`
	DomainExpired   bool              `xml:"DomainExpired"`
	AboutUsLink     string            `xml:"AboutusLink"`
	Row             WhoisRow          `xml:"row"`
	Contacts        []ContactEntry    `xml:"contacts>contact"`
	RRPInfo         WhoisRRPInfo      `xml:"rrp-info"`
	BusinessListing ContactAttributes `xml:"businesslisting"`
	Legal           string            `xml:"legal"`
	WPPSEnabled     bool              `xml:"WPPSEnabled"`
	Version         int               `xml:"Version"`
}

type GetWPPSInfoResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	GetWPPSInfo    GetWPPSInfo `xml:"GetWPPSInfo"`
	APIType        string      `xml:"APIType"`
	ResponseCount  int         `xml:"ResponseCount"`
	MinPeriod      int         `xml:"MinPeriod"`
	MaxPeriod      int         `xml:"MaxPeriod"`
	Server         string      `xml:"Server"`
	Site           string      `xml:"Site"`
	IsLockable     bool        `xml:"IsLockable"`
	IsRealTimeTLD  bool        `xml:"IsRealTimeTLD"`
	TimeDifference string      `xml:"TimeDifference"`
	ExecTime       string      `xml:"ExecTime"`
	Done           bool        `xml:"Done"`
	TrackingKey    string      `xml:"TrackingKey"`
	RequestDate    string      `xml:"RequestDateTime"`
}

type GetWPPSInfo struct {
	DomainName    DomainName     `xml:"domainname"`
	WPPSAllowed   bool           `xml:"WPPSAllowed"`
	WPPSExists    bool           `xml:"WPPSExists"`
	WPPSEnabled   bool           `xml:"WPPSEnabled"`
	WPPSExpDate   string         `xml:"WPPSExpDate"`
	WPPSAutoRenew string         `xml:"WPPSAutoRenew"`
	WPPSPrice     string         `xml:"WPPSPrice"`
	Contacts      []ContactEntry `xml:"contacts>contact"`
}

type DomainName struct {
	SLD          string `xml:"sld,attr"`
	TLD          string `xml:"tld,attr"`
	DomainNameID int64  `xml:"domainnameid,attr"`
	Name         string `xml:",chardata"`
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
	Category string `xml:"category,attr"`
	Value    string `xml:",chardata"`
}

type WhoisPublicity struct {
	VASItemID      int64  `xml:"VASItemID"`
	ProdStatusID   int64  `xml:"ProdStatusID"`
	ProdStatusDesc string `xml:"ProdStatusDesc"`
	ProdEnabled    bool   `xml:"ProdEnabled"`
	ProdConsented  bool   `xml:"ProdConsented"`
	AutoRenew      bool   `xml:"AutoRenew"`
	ExpDate        string `xml:"ExpDate"`
}

type WhoisRow struct {
	RecordFound              string `xml:"RecordFound,attr"`
	IsExpired                string `xml:"IsExpired,attr"`
	RRProcessor              string `xml:"RRProcessor,attr"`
	Registrar                string `xml:"Registrar,attr"`
	LockStatus               string `xml:"LockStatus,attr"`
	AbuseContactEmail        string `xml:"AbuseContactEmail,attr"`
	AbuseContactPhone        string `xml:"AbuseContactPhone,attr"`
	DomainRoid               string `xml:"DomainRoid,attr"`
	ExpDate                  string `xml:"ExpDate,attr"`
	CreateDate               string `xml:"CreateDate,attr"`
	LastUpdatedDate          string `xml:"LastUpdatedDate,attr"`
	NSStatus                 string `xml:"NSStatus,attr"`
	DnsSec                   string `xml:"DnsSec,attr"`
	DnsSecStatus             string `xml:"DnsSecStatus,attr"`
	ResellerOrganizationName string `xml:"ResellerOrganizationName,attr"`
	ResellerURL              string `xml:"ResellerURL,attr"`
	AboutUsLink              string `xml:"AboutUsLink,attr"`
	ResellerEmailAddress     string `xml:"ResellerEmailAddress,attr"`
	RCOMPitch                string `xml:"RCOMPitch,attr"`
	AbuseURL                 string `xml:"AbuseURL,attr"`
	IanaID                   string `xml:"IanaID,attr"`
	RegistrarWhoisServer     string `xml:"RegistrarWhoisServer,attr"`
	RegistrarURL             string `xml:"RegistrarURL,attr"`
	Idp                      string `xml:"Idp,attr"`
	Wps                      string `xml:"Wps,attr"`
}

type WhoisRRPInfo struct {
	RegistrationExpirationDate string   `xml:"registration-expiration-date"`
	CreatedDate                string   `xml:"created-date"`
	UpdatedDate                string   `xml:"updated-date"`
	Status                     []string `xml:"status>status"`
	Domain                     string   `xml:"domain"`
	NameServers                []string `xml:"nameserver>nameserver"`
}

type ContactAttributes struct {
	Raw string `xml:",innerxml"`
}

func (c *ContactData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return decodeContactFields(d, start, c)
}

func (c *ContactEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if strings.EqualFold(attr.Name.Local, "ContactType") {
			c.ContactType = attr.Value
			break
		}
	}
	return decodeContactFields(d, start, &c.ContactData)
}

func decodeContactFields(d *xml.Decoder, start xml.StartElement, contact *ContactData) error {
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch tok := token.(type) {
		case xml.StartElement:
			var value string
			if err := d.DecodeElement(&value, &tok); err != nil {
				return err
			}
			setContactField(contact, tok.Name.Local, value)
		case xml.EndElement:
			if tok.Name == start.Name {
				return nil
			}
		}
	}
}

func setContactField(contact *ContactData, name, value string) {
	switch strings.ToLower(name) {
	case "organization", "organizationname":
		contact.Organization = value
	case "fname", "firstname":
		contact.FirstName = value
	case "lname", "lastname":
		contact.LastName = value
	case "jobtitle":
		contact.JobTitle = value
	case "address1":
		contact.Address1 = value
	case "address2":
		contact.Address2 = value
	case "city":
		contact.City = value
	case "stateprovincechoice":
		contact.StateProvinceChoice = value
	case "stateprovince":
		contact.StateProvince = value
	case "postalcode":
		contact.PostalCode = value
	case "country":
		contact.Country = value
	case "emailaddress":
		contact.EmailAddress = value
	case "phone":
		contact.Phone = value
	case "phoneext":
		contact.PhoneExt = value
	case "fax":
		contact.Fax = value
	case "consentstatus", "registrantconsentstatus":
		contact.ConsentStatus = value
	}
}

func (r *ContactsResponse) Decode() *response.Contacts {
	return &response.Contacts{
		ConsentStatus:  r.ConsentStatus,
		Command:        r.Command,
		APIType:        r.APIType,
		Language:       r.Language,
		ErrCount:       r.ErrCount,
		ResponseCount:  r.ResponseCount,
		MinPeriod:      r.MinPeriod,
		MaxPeriod:      r.MaxPeriod,
		Server:         r.Server,
		Site:           r.Site,
		IsLockable:     r.IsLockable,
		IsRealTimeTLD:  r.IsRealTimeTLD,
		TimeDifference: r.TimeDifference,
		ExecTime:       r.ExecTime,
		Done:           r.Done,
		TrackingKey:    r.TrackingKey,
		RequestDate:    r.RequestDate,
	}
}

func (r *GetContactsResponse) Decode() *response.GetContacts {
	result := response.GetContacts{
		DomainName:          decodeDomainName(r.GetContacts.DomainName),
		Registrant:          decodeContactData(r.GetContacts.Registrant),
		AuxBilling:          decodeContactData(r.GetContacts.AuxBilling),
		Tech:                decodeContactData(r.GetContacts.Tech),
		Admin:               decodeContactData(r.GetContacts.Admin),
		Billing:             decodeContactData(r.GetContacts.Billing),
		ICANNCompliant:      r.GetContacts.ICANNCompliant,
		PendingVerification: r.GetContacts.PendingVerification,
		RAAStatus:           r.GetContacts.RAAStatus,
		RAAStatusDesc:       r.GetContacts.RAAStatusDesc,
		IRTPOptOut:          r.GetContacts.IRTPOptOut,
		TransferLock:        r.GetContacts.TransferLock,
		TransferLockExpDate: response.DomainTransferLockExpDate{
			DaysRemaining: r.GetContacts.TransferLockExpDate.DaysRemaining,
			UTC:           r.GetContacts.TransferLockExpDate.UTC,
			Epoch:         r.GetContacts.TransferLockExpDate.Epoch,
			Value:         r.GetContacts.TransferLockExpDate.Value,
		},
		Nexus: response.Nexus{
			Category: r.GetContacts.Nexus.Category,
			Value:    r.GetContacts.Nexus.Value,
		},
		Purpose:              r.GetContacts.Purpose,
		CurrentAttributes:    r.GetContacts.CurrentAttributes,
		WPPSAllowed:          r.GetContacts.WPPSAllowed,
		WPPSExists:           r.GetContacts.WPPSExists,
		WPPSEnabled:          r.GetContacts.WPPSEnabled,
		WPPSExpDate:          r.GetContacts.WPPSExpDate,
		WPPSAutoRenew:        r.GetContacts.WPPSAutoRenew,
		WPPSContactData:      decodeContactData(r.GetContacts.WPPSContactData),
		EscrowLiftDate:       r.GetContacts.EscrowLiftDate,
		EscrowHold:           r.GetContacts.EscrowHold,
		AttributesXML:        r.GetContacts.Attributes.Raw,
		IsContactShared:      r.GetContacts.IsContactShared,
		ContactRestrictedTLD: r.GetContacts.ContactRestrictedTLD,
		Command:              r.Command,
		APIType:              r.APIType,
		Language:             r.Language,
		ErrCount:             r.ErrCount,
		ResponseCount:        r.ResponseCount,
		MinPeriod:            r.MinPeriod,
		MaxPeriod:            r.MaxPeriod,
		Server:               r.Server,
		Site:                 r.Site,
		IsLockable:           r.IsLockable,
		IsRealTimeTLD:        r.IsRealTimeTLD,
		TimeDifference:       r.TimeDifference,
		ExecTime:             r.ExecTime,
		Done:                 r.Done,
		TrackingKey:          r.TrackingKey,
		RequestDate:          r.RequestDate,
	}
	if r.GetContacts.WhoisPublicity != nil {
		result.WhoisPublicity = &response.WhoisPublicity{
			VASItemID:      r.GetContacts.WhoisPublicity.VASItemID,
			ProdStatusID:   r.GetContacts.WhoisPublicity.ProdStatusID,
			ProdStatusDesc: r.GetContacts.WhoisPublicity.ProdStatusDesc,
			ProdEnabled:    r.GetContacts.WhoisPublicity.ProdEnabled,
			ProdConsented:  r.GetContacts.WhoisPublicity.ProdConsented,
			AutoRenew:      r.GetContacts.WhoisPublicity.AutoRenew,
			ExpDate:        r.GetContacts.WhoisPublicity.ExpDate,
		}
	}
	return &result
}

func (r *GetWhoisContactResponse) Decode() *response.GetWhoisContact {
	result := response.GetWhoisContact{
		DomainName:    decodeDomainName(r.GetWhoisContacts.DomainName),
		DomainExpired: r.GetWhoisContacts.DomainExpired,
		AboutUsLink:   r.GetWhoisContacts.AboutUsLink,
		Row: response.WhoisRow{
			RecordFound:              r.GetWhoisContacts.Row.RecordFound,
			IsExpired:                r.GetWhoisContacts.Row.IsExpired,
			RRProcessor:              r.GetWhoisContacts.Row.RRProcessor,
			Registrar:                r.GetWhoisContacts.Row.Registrar,
			LockStatus:               r.GetWhoisContacts.Row.LockStatus,
			AbuseContactEmail:        r.GetWhoisContacts.Row.AbuseContactEmail,
			AbuseContactPhone:        r.GetWhoisContacts.Row.AbuseContactPhone,
			DomainRoid:               r.GetWhoisContacts.Row.DomainRoid,
			ExpDate:                  r.GetWhoisContacts.Row.ExpDate,
			CreateDate:               r.GetWhoisContacts.Row.CreateDate,
			LastUpdatedDate:          r.GetWhoisContacts.Row.LastUpdatedDate,
			NSStatus:                 r.GetWhoisContacts.Row.NSStatus,
			DnsSec:                   r.GetWhoisContacts.Row.DnsSec,
			DnsSecStatus:             r.GetWhoisContacts.Row.DnsSecStatus,
			ResellerOrganizationName: r.GetWhoisContacts.Row.ResellerOrganizationName,
			ResellerURL:              r.GetWhoisContacts.Row.ResellerURL,
			AboutUsLink:              r.GetWhoisContacts.Row.AboutUsLink,
			ResellerEmailAddress:     r.GetWhoisContacts.Row.ResellerEmailAddress,
			RCOMPitch:                r.GetWhoisContacts.Row.RCOMPitch,
			AbuseURL:                 r.GetWhoisContacts.Row.AbuseURL,
			IanaID:                   r.GetWhoisContacts.Row.IanaID,
			RegistrarWhoisServer:     r.GetWhoisContacts.Row.RegistrarWhoisServer,
			RegistrarURL:             r.GetWhoisContacts.Row.RegistrarURL,
			Idp:                      r.GetWhoisContacts.Row.Idp,
			Wps:                      r.GetWhoisContacts.Row.Wps,
		},
		Contacts: decodeContactEntries(r.GetWhoisContacts.Contacts),
		RRPInfo: response.WhoisRRPInfo{
			RegistrationExpirationDate: r.GetWhoisContacts.RRPInfo.RegistrationExpirationDate,
			CreatedDate:                r.GetWhoisContacts.RRPInfo.CreatedDate,
			UpdatedDate:                r.GetWhoisContacts.RRPInfo.UpdatedDate,
			Status:                     append([]string(nil), r.GetWhoisContacts.RRPInfo.Status...),
			Domain:                     r.GetWhoisContacts.RRPInfo.Domain,
			NameServers:                append([]string(nil), r.GetWhoisContacts.RRPInfo.NameServers...),
		},
		BusinessListingXML: r.GetWhoisContacts.BusinessListing.Raw,
		Legal:              r.GetWhoisContacts.Legal,
		WPPSEnabled:        r.GetWhoisContacts.WPPSEnabled,
		Version:            r.GetWhoisContacts.Version,
		Success:            r.Success,
		Command:            r.Command,
		APIType:            r.APIType,
		Language:           r.Language,
		ErrCount:           r.ErrCount,
		ResponseCount:      r.ResponseCount,
		MinPeriod:          r.MinPeriod,
		MaxPeriod:          r.MaxPeriod,
		Server:             r.Server,
		Site:               r.Site,
		IsLockable:         r.IsLockable,
		IsRealTimeTLD:      r.IsRealTimeTLD,
		TimeDifference:     r.TimeDifference,
		ExecTime:           r.ExecTime,
		Done:               r.Done,
		TrackingKey:        r.TrackingKey,
		RequestDate:        r.RequestDate,
	}
	return &result
}

func (r *GetWPPSInfoResponse) Decode() *response.GetWPPSInfo {
	result := response.GetWPPSInfo{
		DomainName:     decodeDomainName(r.GetWPPSInfo.DomainName),
		WPPSAllowed:    r.GetWPPSInfo.WPPSAllowed,
		WPPSExists:     r.GetWPPSInfo.WPPSExists,
		WPPSEnabled:    r.GetWPPSInfo.WPPSEnabled,
		WPPSExpDate:    r.GetWPPSInfo.WPPSExpDate,
		WPPSAutoRenew:  r.GetWPPSInfo.WPPSAutoRenew,
		WPPSPrice:      r.GetWPPSInfo.WPPSPrice,
		Contacts:       decodeContactEntries(r.GetWPPSInfo.Contacts),
		Command:        r.Command,
		APIType:        r.APIType,
		Language:       r.Language,
		ErrCount:       r.ErrCount,
		ResponseCount:  r.ResponseCount,
		MinPeriod:      r.MinPeriod,
		MaxPeriod:      r.MaxPeriod,
		Server:         r.Server,
		Site:           r.Site,
		IsLockable:     r.IsLockable,
		IsRealTimeTLD:  r.IsRealTimeTLD,
		TimeDifference: r.TimeDifference,
		ExecTime:       r.ExecTime,
		Done:           r.Done,
		TrackingKey:    r.TrackingKey,
		RequestDate:    r.RequestDate,
	}
	return &result
}

func decodeContactData(contact ContactData) response.ContactData {
	return response.ContactData{
		Organization:        contact.Organization,
		FirstName:           contact.FirstName,
		LastName:            contact.LastName,
		JobTitle:            contact.JobTitle,
		Address1:            contact.Address1,
		Address2:            contact.Address2,
		City:                contact.City,
		StateProvinceChoice: contact.StateProvinceChoice,
		StateProvince:       contact.StateProvince,
		PostalCode:          contact.PostalCode,
		Country:             contact.Country,
		EmailAddress:        contact.EmailAddress,
		Phone:               contact.Phone,
		PhoneExt:            contact.PhoneExt,
		Fax:                 contact.Fax,
		ConsentStatus:       contact.ConsentStatus,
	}
}

func decodeContactEntries(entries []ContactEntry) []response.ContactEntry {
	if len(entries) == 0 {
		return nil
	}
	result := make([]response.ContactEntry, 0, len(entries))
	for _, entry := range entries {
		result = append(result, response.ContactEntry{
			ContactType: entry.ContactType,
			ContactData: decodeContactData(entry.ContactData),
		})
	}
	return result
}

func decodeDomainName(name DomainName) response.DomainName {
	return response.DomainName{
		SLD:          name.SLD,
		TLD:          name.TLD,
		DomainNameID: name.DomainNameID,
		Name:         name.Name,
	}
}
