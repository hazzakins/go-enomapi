package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

type GetDomainInfoResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	GetDomainInfo  GetDomainInfo `xml:"GetDomainInfo"`
	APIType        string        `xml:"APIType"`
	ResponseCount  int           `xml:"ResponseCount"`
	MinPeriod      int           `xml:"MinPeriod"`
	MaxPeriod      int           `xml:"MaxPeriod"`
	Server         string        `xml:"Server"`
	Site           string        `xml:"Site"`
	IsLockable     bool          `xml:"IsLockable"`
	IsRealTimeTLD  bool          `xml:"IsRealTimeTLD"`
	TimeDifference string        `xml:"TimeDifference"`
	ExecTime       string        `xml:"ExecTime"`
	Done           bool          `xml:"Done"`
	TrackingKey    string        `xml:"TrackingKey"`
	RequestDate    string        `xml:"RequestDateTime"`
}

type GetDomainInfo struct {
	DomainName     DomainInfoDomainName `xml:"domainname"`
	MultyLangSLD   bool                 `xml:"multy-langSLD"`
	Status         DomainInfoStatus     `xml:"status"`
	ParkingEnabled bool                 `xml:"ParkingEnabled"`
	Services       []DomainServiceEntry `xml:"services>entry"`
}

type DomainInfoDomainName struct {
	SLD          string `xml:"sld,attr"`
	TLD          string `xml:"tld,attr"`
	DomainNameID int64  `xml:"domainnameid,attr"`
	Name         string `xml:",chardata"`
}

type DomainInfoStatus struct {
	Expiration            string          `xml:"expiration"`
	DeleteByDate          string          `xml:"deletebydate"`
	DeleteType            string          `xml:"deletetype"`
	Restorable            bool            `xml:"restorable"`
	RenewBeforeExpiration string          `xml:"renewbeforeexpiration"`
	Registrar             string          `xml:"registrar"`
	RegistrationStatus    string          `xml:"registrationstatus"`
	PurchaseStatus        string          `xml:"purchase-status"`
	BelongsTo             DomainBelongsTo `xml:"belongs-to"`
	EscrowHold            bool            `xml:"escrowhold"`
	EscrowLiftDate        string          `xml:"escrowliftdate"`
	AuctionHold           bool            `xml:"auctionhold"`
	AuctionLiftDate       string          `xml:"auctionliftdate"`
}

type DomainBelongsTo struct {
	PartyID string `xml:"party-id,attr"`
	Value   string `xml:",chardata"`
}

type DomainServiceEntry struct {
	Name           string                    `xml:"name,attr"`
	EnomDNS        *DomainServiceEnomDNS     `xml:"enomDNS"`
	Service        *DomainService            `xml:"service"`
	Configuration  *DomainServiceConfig      `xml:"configuration"`
	WBL            *DomainServiceWBL         `xml:"wbl"`
	Mobilizer      *DomainServiceMobilizer   `xml:"mobilizer"`
	RAASetting     *DomainServiceRAASetting  `xml:"raasetting"`
	IRTPSetting    *DomainServiceIRTPSetting `xml:"irtpsetting"`
	WhoisPublicity *DomainServiceWhoisPublic `xml:"whoispublicity"`
}

type DomainServiceEnomDNS struct {
	Value     string `xml:"value,attr"`
	IsDotName string `xml:"isDotName,attr"`
}

type DomainService struct {
	Changable bool   `xml:"changable,attr"`
	Value     string `xml:",chardata"`
}

type DomainServiceConfig struct {
	Changable    bool                `xml:"changable,attr"`
	Type         string              `xml:"type,attr"`
	DNS          []string            `xml:"dns"`
	Hosts        []DomainServiceHost `xml:"host"`
	WSB          string              `xml:"wsb"`
	SiteID       string              `xml:"siteid"`
	ProdType     string              `xml:"prodtype"`
	NextBillDate string              `xml:"nextbilldate"`
	WPPS         *DomainServiceWPPS  `xml:"wpps"`
}

type DomainServiceHost struct {
	Name       string `xml:"name"`
	Type       string `xml:"type"`
	Address    string `xml:"address"`
	MXPref     string `xml:"mxpref"`
	IsEditable bool   `xml:"iseditable"`
}

type DomainServiceWPPS struct {
	CloakedEmail string `xml:"cloakedemail"`
	ForwardTo    string `xml:"forward-to"`
	ExpireDate   string `xml:"expiredate"`
	AutoRenew    string `xml:"autorenew"`
}

type DomainServiceWBL struct {
	WBLID              string               `xml:"wblid"`
	StatusID           string               `xml:"statusid"`
	StatusDescr        string               `xml:"statusdescr"`
	ExpDate            string               `xml:"expdate"`
	Enabled            bool                 `xml:"enabled"`
	Renew              bool                 `xml:"renew"`
	CompanyName        string               `xml:"companyname"`
	CompanyDescription string               `xml:"companydescription"`
	DomainName         string               `xml:"domainname"`
	Street             string               `xml:"street"`
	City               string               `xml:"city"`
	PostalCode         string               `xml:"postalcode"`
	Country            string               `xml:"country"`
	CategoryIDX        string               `xml:"categoryidx"`
	Fields             []DomainServiceField `xml:"field"`
	FieldName          string               `xml:"fieldname"`
	Value              string               `xml:"value"`
}

type DomainServiceField struct {
	FieldName string `xml:"fieldname"`
	Value     string `xml:"value"`
}

type DomainServiceMobilizer struct{}

type DomainServiceRAASetting struct {
	VerificationStatus string `xml:"verificationstatus"`
	DomainSuspended    bool   `xml:"domainsuspended"`
	SuspensionDate     string `xml:"suspensiondate"`
	IsQueuedChange     bool   `xml:"isqueuedchange"`
	StatusExpDate      string `xml:"statusexpdate"`
}

type DomainServiceIRTPSetting struct {
	ICANNCompliant      bool                      `xml:"icanncompliant"`
	OptOut              bool                      `xml:"optout"`
	TransferLock        bool                      `xml:"transferlock"`
	TransferLockExpDate DomainTransferLockExpDate `xml:"transferlockexpdate"`
}

type DomainTransferLockExpDate struct {
	DaysRemaining int    `xml:"daysremaining,attr"`
	UTC           string `xml:"utc,attr"`
	Epoch         int64  `xml:"epoch,attr"`
	Value         string `xml:",chardata"`
}

type DomainServiceWhoisPublic struct {
	VASItemID int64 `xml:"vasitemid"`
	Enabled   bool  `xml:"enabled"`
}

func (r *GetDomainInfoResponse) Decode() *response.GetDomainInfo {
	result := response.GetDomainInfo{
		DomainName: response.DomainInfoDomainName{
			SLD:          r.GetDomainInfo.DomainName.SLD,
			TLD:          r.GetDomainInfo.DomainName.TLD,
			DomainNameID: r.GetDomainInfo.DomainName.DomainNameID,
			Name:         r.GetDomainInfo.DomainName.Name,
		},
		MultyLangSLD:   r.GetDomainInfo.MultyLangSLD,
		Status:         decodeDomainInfoStatus(r.GetDomainInfo.Status),
		ParkingEnabled: r.GetDomainInfo.ParkingEnabled,
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

	if len(r.GetDomainInfo.Services) > 0 {
		result.Services = make([]response.DomainServiceEntry, 0, len(r.GetDomainInfo.Services))
		for _, entry := range r.GetDomainInfo.Services {
			result.Services = append(result.Services, decodeDomainServiceEntry(entry))
		}
	}

	return &result
}

func decodeDomainInfoStatus(status DomainInfoStatus) response.DomainInfoStatus {
	return response.DomainInfoStatus{
		Expiration:            status.Expiration,
		DeleteByDate:          status.DeleteByDate,
		DeleteType:            status.DeleteType,
		Restorable:            status.Restorable,
		RenewBeforeExpiration: status.RenewBeforeExpiration,
		Registrar:             status.Registrar,
		RegistrationStatus:    status.RegistrationStatus,
		PurchaseStatus:        status.PurchaseStatus,
		BelongsTo: response.DomainBelongsTo{
			PartyID: status.BelongsTo.PartyID,
			Value:   status.BelongsTo.Value,
		},
		EscrowHold:      status.EscrowHold,
		EscrowLiftDate:  status.EscrowLiftDate,
		AuctionHold:     status.AuctionHold,
		AuctionLiftDate: status.AuctionLiftDate,
	}
}

func decodeDomainServiceEntry(entry DomainServiceEntry) response.DomainServiceEntry {
	result := response.DomainServiceEntry{
		Name: entry.Name,
	}

	if entry.EnomDNS != nil {
		result.EnomDNS = &response.DomainServiceEnomDNS{
			Value:     entry.EnomDNS.Value,
			IsDotName: entry.EnomDNS.IsDotName,
		}
	}

	if entry.Service != nil {
		result.Service = &response.DomainService{
			Changable: entry.Service.Changable,
			Value:     entry.Service.Value,
		}
	}

	if entry.Configuration != nil {
		result.Configuration = decodeDomainServiceConfig(entry.Configuration)
	}

	if entry.WBL != nil {
		result.WBL = decodeDomainServiceWBL(entry.WBL)
	}

	if entry.Mobilizer != nil {
		result.Mobilizer = &response.DomainServiceMobilizer{}
	}

	if entry.RAASetting != nil {
		result.RAASetting = &response.DomainServiceRAASetting{
			VerificationStatus: entry.RAASetting.VerificationStatus,
			DomainSuspended:    entry.RAASetting.DomainSuspended,
			SuspensionDate:     entry.RAASetting.SuspensionDate,
			IsQueuedChange:     entry.RAASetting.IsQueuedChange,
			StatusExpDate:      entry.RAASetting.StatusExpDate,
		}
	}

	if entry.IRTPSetting != nil {
		result.IRTPSetting = &response.DomainServiceIRTPSetting{
			ICANNCompliant: entry.IRTPSetting.ICANNCompliant,
			OptOut:         entry.IRTPSetting.OptOut,
			TransferLock:   entry.IRTPSetting.TransferLock,
			TransferLockExpDate: response.DomainTransferLockExpDate{
				DaysRemaining: entry.IRTPSetting.TransferLockExpDate.DaysRemaining,
				UTC:           entry.IRTPSetting.TransferLockExpDate.UTC,
				Epoch:         entry.IRTPSetting.TransferLockExpDate.Epoch,
				Value:         entry.IRTPSetting.TransferLockExpDate.Value,
			},
		}
	}

	if entry.WhoisPublicity != nil {
		result.WhoisPublicity = &response.DomainServiceWhoisPublic{
			VASItemID: entry.WhoisPublicity.VASItemID,
			Enabled:   entry.WhoisPublicity.Enabled,
		}
	}

	return result
}

func decodeDomainServiceConfig(config *DomainServiceConfig) *response.DomainServiceConfig {
	result := &response.DomainServiceConfig{
		Changable:    config.Changable,
		Type:         config.Type,
		DNS:          append([]string(nil), config.DNS...),
		WSB:          config.WSB,
		SiteID:       config.SiteID,
		ProdType:     config.ProdType,
		NextBillDate: config.NextBillDate,
	}

	if len(config.Hosts) > 0 {
		result.Hosts = make([]response.DomainServiceHost, 0, len(config.Hosts))
		for _, host := range config.Hosts {
			result.Hosts = append(result.Hosts, response.DomainServiceHost{
				Name:       host.Name,
				Type:       host.Type,
				Address:    host.Address,
				MXPref:     host.MXPref,
				IsEditable: host.IsEditable,
			})
		}
	}

	if config.WPPS != nil {
		result.WPPS = &response.DomainServiceWPPS{
			CloakedEmail: config.WPPS.CloakedEmail,
			ForwardTo:    config.WPPS.ForwardTo,
			ExpireDate:   config.WPPS.ExpireDate,
			AutoRenew:    config.WPPS.AutoRenew,
		}
	}

	return result
}

func decodeDomainServiceWBL(wbl *DomainServiceWBL) *response.DomainServiceWBL {
	result := &response.DomainServiceWBL{
		WBLID:              wbl.WBLID,
		StatusID:           wbl.StatusID,
		StatusDescr:        wbl.StatusDescr,
		ExpDate:            wbl.ExpDate,
		Enabled:            wbl.Enabled,
		Renew:              wbl.Renew,
		CompanyName:        wbl.CompanyName,
		CompanyDescription: wbl.CompanyDescription,
		DomainName:         wbl.DomainName,
		Street:             wbl.Street,
		City:               wbl.City,
		PostalCode:         wbl.PostalCode,
		Country:            wbl.Country,
		CategoryIDX:        wbl.CategoryIDX,
		FieldName:          wbl.FieldName,
		Value:              wbl.Value,
	}

	if len(wbl.Fields) > 0 {
		result.Fields = make([]response.DomainServiceField, 0, len(wbl.Fields))
		for _, field := range wbl.Fields {
			result.Fields = append(result.Fields, response.DomainServiceField{
				FieldName: field.FieldName,
				Value:     field.Value,
			})
		}
	}

	return result
}
