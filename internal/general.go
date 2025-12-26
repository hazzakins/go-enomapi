package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

type ResponseMeta struct {
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

type AdvancedDomainSearchResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DomainSearch AdvancedDomainSearch `xml:"DomainSearch"`
}

type AdvancedDomainSearch struct {
	SearchParams  AdvancedDomainSearchParams   `xml:"SearchParams"`
	TotalResults  int                          `xml:"TotalResults"`
	StartPosition int                          `xml:"StartPosition"`
	NextPosition  int                          `xml:"NextPosition"`
	MultiRRP      bool                         `xml:"MultiRRP"`
	TLDOverride   string                       `xml:"TLDOverride"`
	Domains       []AdvancedDomainSearchDomain `xml:"Domains>Domain"`
}

type AdvancedDomainSearchParams struct {
	TLDList             string                  `xml:"SP-TLDList"`
	SLD                 string                  `xml:"SP-SLD"`
	SearchCriteria      string                  `xml:"SP-SearchCriteria"`
	ParkingStatus       string                  `xml:"SP-ParkingStatus"`
	RegistrationStatus  string                  `xml:"SP-RegistrationStatus"`
	AutoRenew           string                  `xml:"SP-AutoRenew"`
	Locked              string                  `xml:"SP-Locked"`
	DaysTillExpires     string                  `xml:"SP-DaysTillExpires"`
	DaysExpired         string                  `xml:"SP-DaysExpired"`
	NSStatus            string                  `xml:"SP-NsStatus"`
	NameServer          string                  `xml:"SP-NameServer"`
	HasIDProtect        string                  `xml:"SP-HasIDProtect"`
	HasPOPMail          string                  `xml:"SP-HasPOPMail"`
	EmailForwarding     string                  `xml:"SP-EmailForwarding"`
	HasWebHosting       string                  `xml:"SP-HasWebHosting"`
	ExcludeNumbers      string                  `xml:"SP-ExcludeNumbers"`
	ExcludeDashes       string                  `xml:"SP-ExcludeDashes"`
	IncludeSubAccounts  string                  `xml:"SP-IncludeSubAccounts"`
	SubAccountLogin     string                  `xml:"SP-SubAccountLogin"`
	RecordsToReturn     string                  `xml:"SP-RecordsToReturn"`
	StartPosition       string                  `xml:"SP-StartPosition"`
	OrderBy             string                  `xml:"SP-OrderBy"`
	FolderOption        string                  `xml:"SP-FolderOption"`
	FolderSyncStatus    string                  `xml:"SP-FolderSyncStatus"`
	FolderName          string                  `xml:"SP-FolderName"`
	CreationDate        string                  `xml:"SP-CreationDate"`
	DaysUntilIDPExpires string                  `xml:"SP-DaysUntilIDProtectExpires"`
	WBLStatusID         string                  `xml:"SP-WBLStatusID"`
	WBLEnabled          string                  `xml:"SP-WBLEnabled"`
	WBLAutoRenew        string                  `xml:"SP-WBLAutoRenew"`
	ContactXML          string                  `xml:"SP-ContactXML"`
	XMLResponse         string                  `xml:"SP-XMLResponse"`
	HostRecordType      string                  `xml:"SP-HostRecordType"`
	HostName            string                  `xml:"SP-HostName"`
	HostAddress         string                  `xml:"SP-HostAddress"`
	EmailResultsOnly    string                  `xml:"SP-EmailResultsOnly"`
	XML                 AdvancedDomainSearchXML `xml:"SP-XML"`
}

type AdvancedDomainSearchXML struct {
	Raw string `xml:",innerxml"`
}

type AdvancedDomainSearchDomain struct {
	DomainNameID               int64  `xml:"DomainNameID"`
	SLD                        string `xml:"SLD"`
	TLD                        string `xml:"TLD"`
	AutoRenew                  string `xml:"AutoRenew"`
	ExpDate                    string `xml:"ExpDate"`
	DomainRegistrationStatus   string `xml:"DomainRegistrationStatus"`
	DeleteType                 string `xml:"DeleteType"`
	LoginID                    string `xml:"LoginID"`
	AccountID                  string `xml:"AccountID"`
	NSStatus                   string `xml:"NSStatus"`
	FolderStatus               string `xml:"FolderStatus"`
	RRProcessor                string `xml:"RRProcessor"`
	RRCompanyName              string `xml:"RRCompanyName"`
	HasIDProtect               string `xml:"HasIDProtect"`
	IDProtectExpires           string `xml:"IDProtectExpires"`
	DomainFolderStatus         string `xml:"DomainFolderStatus"`
	AbleToReactivate           string `xml:"AbleToReactivate"`
	IsPremiumName              string `xml:"IsPremiumName"`
	PremiumPrice               string `xml:"PremiumPrice"`
	PremiumAboveThresholdPrice string `xml:"PremiumAboveThresholdPrice"`
	PremiumCategory            string `xml:"PremiumCategory"`
	ReactivatePrice            string `xml:"ReactivatePrice"`
	WBLStatusID                string `xml:"WBLStatusID"`
	WBLStatus                  string `xml:"WBLStatus"`
	WBLExpDate                 string `xml:"WBLExpDate"`
	WBLAutoRenew               string `xml:"WBLAutoRenew"`
	NameServers                string `xml:"NameServers"`
	Vas                        string `xml:"Vas"`
}

type GetAllDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	GetAllDomains GetAllDomains `xml:"GetAllDomains"`
}

type GetAllDomains struct {
	DomainDetails     []GetAllDomainsDetail `xml:"DomainDetail"`
	DomainCount       int                   `xml:"domaincount"`
	UserRequestStatus string                `xml:"UserRequestStatus"`
}

type GetAllDomainsDetail struct {
	DomainName     string `xml:"DomainName"`
	DomainNameID   int64  `xml:"DomainNameID"`
	ExpirationDate string `xml:"expiration-date"`
	LockStatus     string `xml:"lockstatus"`
	AutoRenew      string `xml:"AutoRenew"`
}

type GetDomainCountResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RegisteredCount     int            `xml:"RegisteredCount"`
	HostCount           int            `xml:"HostCount"`
	ExpiringCount       int            `xml:"ExpiringCount"`
	ExpiredDomainsCount int            `xml:"ExpiredDomainsCount"`
	RGP                 int            `xml:"RGP"`
	ExtendedRGP         int            `xml:"ExtendedRGP"`
	KeywordCount        int            `xml:"KeywordCount"`
	ProcessCount        int            `xml:"ProcessCount"`
	WatchlistCount      int            `xml:"WatchlistCount"`
	CartItemCount       int            `xml:"CartItemCount"`
	TrafficMsg          *DomainTraffic `xml:"trafficmsg"`
}

type GetDomainExpResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ExpirationDate string `xml:"ExpirationDate"`
}

type DomainTraffic struct {
	VistaCustomer  bool   `xml:"VistaCustomer"`
	RedirectorData bool   `xml:"RedirectorData"`
	Month          string `xml:"Month"`
	PageViews      string `xml:"PageViews"`
	Visitors       string `xml:"Visitors"`
	FreeTrial      bool   `xml:"FreeTrial"`
	PDQVista       string `xml:"PDQVista"`
}

type GetDomainNameIDResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DomainRRP    string `xml:"DomainRRP"`
	SLD          string `xml:"SLD"`
	TLD          string `xml:"TLD"`
	DomainNameID int64  `xml:"DomainNameID"`
}

type GetDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	GetDomains GetDomains `xml:"GetDomains"`
}

type GetDomains struct {
	Tab              string         `xml:"tab"`
	MultiRRP         bool           `xml:"multirrp"`
	DomainList       GetDomainsList `xml:"domain-list"`
	EndPosition      int            `xml:"EndPosition"`
	PreviousRecords  int            `xml:"PreviousRecords"`
	NextRecords      int            `xml:"NextRecords"`
	OrderBy          string         `xml:"OrderBy"`
	Result           bool           `xml:"Result"`
	StartPosition    int            `xml:"StartPosition"`
	DomainCount      int            `xml:"DomainCount"`
	TotalDomainCount int            `xml:"TotalDomainCount"`
	StartLetter      string         `xml:"StartLetter"`
}

type GetDomainsList struct {
	Type    string           `xml:"type,attr"`
	Domains []GetDomainsItem `xml:"domain"`
}

type GetDomainsItem struct {
	DomainNameID   int64  `xml:"DomainNameID"`
	SLD            string `xml:"sld"`
	TLD            string `xml:"tld"`
	NSStatus       string `xml:"ns-status"`
	ExpirationDate string `xml:"expiration-date"`
	AutoRenew      string `xml:"auto-renew"`
	WPPSStatus     string `xml:"wppsstatus"`
	WPPSExpDate    string `xml:"wppsexpdate"`
	RRProcessor    string `xml:"RRProcessor"`
}

type GetDomainStatusResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DomainStatus DomainStatus `xml:"DomainStatus"`
}

type DomainStatus struct {
	DomainName string `xml:"DomainName"`
	Registrar  string `xml:"Registrar"`
	InAccount  string `xml:"InAccount"`
	StatusDesc string `xml:"StatusDesc"`
	ExpDate    string `xml:"ExpDate"`
	OrderID    string `xml:"OrderID"`
}

type GetDomainSldTldResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DomainRRP    string `xml:"DomainRRP"`
	SLD          string `xml:"SLD"`
	TLD          string `xml:"TLD"`
	DomainNameID int64  `xml:"DomainNameID"`
}

type GetExpiredDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DomainDetails []ExpiredDomainDetail `xml:"DomainDetail"`
	DomainCount   int                   `xml:"domaincount"`
}

type ExpiredDomainDetail struct {
	DomainName     string `xml:"DomainName"`
	DomainNameID   int64  `xml:"DomainNameID"`
	Status         string `xml:"status"`
	ExpirationDate string `xml:"expiration-date"`
	LockStatus     string `xml:"lockstatus"`
}

type GetExtendInfoResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RegistrarHold       bool                          `xml:"RegistrarHold"`
	Expiration          string                        `xml:"Expiration"`
	MaxExtension        int                           `xml:"MaxExtension"`
	MinAllowed          int                           `xml:"MinAllowed"`
	CCAuthorized        bool                          `xml:"CCAuthorized"`
	Price               string                        `xml:"Price"`
	Balance             string                        `xml:"Balance"`
	AvailableBalance    string                        `xml:"AvailableBalance"`
	CustomerPrefs       ExtendInfoCustomerPrefs       `xml:"CustomerPrefs"`
	CustomerInformation ExtendInfoCustomerInformation `xml:"CustomerInformation"`
}

type ExtendInfoCustomerPrefs struct {
	DefPeriod            int                          `xml:"DefPeriod"`
	AllowDNS             bool                         `xml:"AllowDNS"`
	ShowPopups           bool                         `xml:"ShowPopups"`
	AutoRenew            bool                         `xml:"AutoRenew"`
	RegLock              bool                         `xml:"RegLock"`
	AutoPakRenew         bool                         `xml:"AutoPakRenew"`
	UseDNS               bool                         `xml:"UseDNS"`
	ResellerStatus       string                       `xml:"ResellerStatus"`
	RenewalSetting       int                          `xml:"RenewalSetting"`
	RenewalBCC           int                          `xml:"RenewalBCC"`
	RenewalURLForward    bool                         `xml:"RenewalURLForward"`
	RenewalEmailForward  bool                         `xml:"RenewalEmailForward"`
	MailNumLimit         int                          `xml:"MailNumLimit"`
	IDProtect            bool                         `xml:"IDProtect"`
	DefIDProtectRenew    bool                         `xml:"DefIDProtectRenew"`
	DefWBLRenew          bool                         `xml:"DefWBLRenew"`
	NameJetSales         bool                         `xml:"NameJetSales"`
	DefaultHostRecords   ExtendInfoDefaultHostRecords `xml:"defaulthostrecords"`
	DefaultHostRecordOwn bool                         `xml:"defaulthostrecordown"`
	UseOurDNS            bool                         `xml:"UseOurDNS"`
	NameServers          ExtendInfoNameServers        `xml:"NameServers"`
}

type ExtendInfoDefaultHostRecords struct {
	HostRecords []ExtendInfoHostRecord `xml:"hostrecord"`
}

type ExtendInfoHostRecord struct {
	HostName   string `xml:"hostname,attr"`
	Address    string `xml:"address,attr"`
	RecordType string `xml:"recordtype,attr"`
}

type ExtendInfoNameServers struct {
	DNS1 string `xml:"DNS1"`
	DNS2 string `xml:"DNS2"`
	DNS3 string `xml:"DNS3"`
	DNS4 string `xml:"DNS4"`
	DNS5 string `xml:"DNS5"`
}

type ExtendInfoCustomerInformation struct {
	AcceptTerms   bool   `xml:"AcceptTerms"`
	URL           string `xml:"URL"`
	ParentAccount string `xml:"ParentAccount"`
	ParentLogin   string `xml:"ParentLogin"`
	NoService     bool   `xml:"NoService"`
	BulkRegLimit  int    `xml:"BulkRegLimit"`
	Account       string `xml:"Account"`
}

type GetNewsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Alerts Alerts `xml:"Alerts"`
}

type GetProductNewsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Alerts Alerts `xml:"Alerts"`
}

type Alerts struct {
	AlertEntries []Alert `xml:"Alert"`
	Total        int     `xml:"Total"`
}

type Alert struct {
	Product string `xml:"Product"`
	Message string `xml:"Message"`
}

type GetPasswordBitResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DomainPassword string `xml:"DomainPassword"`
	PasswordSet    int    `xml:"password-set"`
}

type GetRegistrationStatusResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RegistrationStatus string `xml:"RegistrationStatus"`
	PurchaseStatus     string `xml:"PurchaseStatus"`
}

type GetRegLockResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RegLock   string `xml:"reg-lock"`
	Registrar string `xml:"registrar"`
}

type GetRenewResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RenewName         bool `xml:"auto-renew"`
	PakExist          bool `xml:"PakExist"`
	AutoPakRenew      bool `xml:"AutoPakRenew"`
	EmailFwdExists    bool `xml:"EmailFwdExists"`
	EmailForwardRenew bool `xml:"EmailForwardRenew"`
	URLFwdExists      bool `xml:"URLFwdExists"`
	URLForwardRenew   bool `xml:"URLForwardRenew"`
	IDProtectRenew    bool `xml:"IDProtectRenew"`
	IDProtectExists   bool `xml:"IDProtectExists"`
	MobilizerRenew    bool `xml:"MobilizerRenew"`
}

type GetSubAccountPasswordResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
}

type ParseDomainResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ParseDomain ParsedDomain `xml:"ParseDomain"`
}

type ParsedDomain struct {
	Host string `xml:"Host"`
	SLD  string `xml:"SLD"`
	TLD  string `xml:"TLD"`
}

type PortalGetAwardedDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Domains PortalAwardedDomains `xml:"Domains"`
	Success string               `xml:"Success"`
}

type PortalAwardedDomains struct {
	Domains           []PortalAwardedDomain `xml:"Domain"`
	DomainCount       int                   `xml:"DomainCount"`
	TotalDomainCount  int                   `xml:"TotalDomainCount"`
	NextStartPosition int                   `xml:"NextStartPosition"`
}

type PortalAwardedDomain struct {
	DomainName              string `xml:"DomainName"`
	EmailAddress            string `xml:"EmailAddress"`
	ExpirationDate          string `xml:"ExpirationDate"`
	RegisterDate            string `xml:"RegisterDate"`
	ForeignLoginID          string `xml:"ForeignLoginId"`
	PortalDomainID          string `xml:"PortalDomainId"`
	RegisterPrice           string `xml:"RegisterPrice"`
	RenewPrice              string `xml:"RenewPrice"`
	RegistrationPeriod      string `xml:"RegistrationPeriod"`
	ResellerProvisioned     string `xml:"ResellerProvisioned"`
	ResellerProvisionedDate string `xml:"ResellerProvisionedDate"`
	IsPremium               string `xml:"IsPremium"`
	RegistrationStatus      string `xml:"RegistrationStatus"`
}

type PortalGetDomainInfoResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Domain       PortalDomainInfo `xml:"Domain"`
	Transactions RawXML           `xml:"Transactions"`
	Services     RawXML           `xml:"services"`
	RAASettings  RawXML           `xml:"RAASettings"`
	Payment      RawXML           `xml:"Payment"`
	Contacts     RawXML           `xml:"Contacts"`
}

type PortalDomainInfo struct {
	DomainName              string `xml:"DomainName"`
	EmailAddress            string `xml:"EmailAddress"`
	ExpirationDate          string `xml:"ExpirationDate"`
	RegisterDate            string `xml:"RegisterDate"`
	ForeignLoginID          string `xml:"ForeignLoginId"`
	PortalDomainID          string `xml:"PortalDomainId"`
	RegisterPrice           string `xml:"RegisterPrice"`
	RenewPrice              string `xml:"RenewPrice"`
	RegistrationPeriod      string `xml:"RegistrationPeriod"`
	ResellerProvisioned     string `xml:"ResellerProvisioned"`
	ResellerProvisionedDate string `xml:"ResellerProvisionedDate"`
	IsPremium               string `xml:"IsPremium"`
	RegistrationStatus      string `xml:"RegistrationStatus"`
}

type RawXML struct {
	Raw string `xml:",innerxml"`
}

type PortalGetTokenResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Token string `xml:"token"`
}

type PortalUpdateAwardedDomainsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Success string `xml:"Success"`
}

type RPTGetReportResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Report Report `xml:"rpt"`
}

type Report struct {
	RptString             string        `xml:"rptString"`
	ThreeMoPast           string        `xml:"threemopast"`
	SixMoPast             string        `xml:"sixmopast"`
	BeginDate             string        `xml:"begindate"`
	EndDate               string        `xml:"enddate"`
	Options               ReportOptions `xml:"option"`
	ReportType            string        `xml:"reporttype"`
	ReportName            string        `xml:"reportname"`
	QueueReportResults    string        `xml:"QueueReportResults"`
	Results               ReportResults `xml:"results"`
	NewStartPosition      string        `xml:"NewStartPosition"`
	ThisStartPosition     string        `xml:"ThisStartPosition"`
	ThisEndPosition       string        `xml:"ThisEndPosition"`
	PreviousStartPosition string        `xml:"PreviousStartPosition"`
	TotalReturned         int           `xml:"TotalReturned"`
	TotalRows             int           `xml:"TotalRows"`
	ShowPrevious          string        `xml:"ShowPrevious"`
	ShowNext              string        `xml:"ShowNext"`
	ShowPaging            string        `xml:"ShowPaging"`
	Version               string        `xml:"Version"`
}

type ReportOptions struct {
	ReportTypes []ReportTypeOption `xml:"reporttype"`
}

type ReportTypeOption struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type ReportResults struct {
	Raw string `xml:",innerxml"`
}

type SetPasswordResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
}

type SetRegLockResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RegLock   string `xml:"reg-lock"`
	RRPCodeSR string `xml:"RRPCodeSR"`
	RRPText   string `xml:"RRPText"`
}

type SetRenewResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RenewName         bool `xml:"RenewName"`
	AutoPakRenew      bool `xml:"AutoPakRenew"`
	EmailForwardRenew bool `xml:"EmailForwardRenew"`
	URLForwardRenew   bool `xml:"URLForwardRenew"`
	WPPSRenew         bool `xml:"WPPSRenew"`
}

type StatusDomainResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DomainStatus DomainStatus `xml:"DomainStatus"`
}

type ValidatePasswordResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
}

func decodeResponseMeta(resp Response, meta ResponseMeta) response.ResponseMeta {
	return response.ResponseMeta{
		Command:        resp.Command,
		Language:       resp.Language,
		ErrCount:       resp.ErrCount,
		RRPCode:        resp.ResponseCode,
		RRPText:        resp.ResponseText,
		APIType:        meta.APIType,
		ResponseCount:  meta.ResponseCount,
		MinPeriod:      meta.MinPeriod,
		MaxPeriod:      meta.MaxPeriod,
		Server:         meta.Server,
		Site:           meta.Site,
		IsLockable:     meta.IsLockable,
		IsRealTimeTLD:  meta.IsRealTimeTLD,
		TimeDifference: meta.TimeDifference,
		ExecTime:       meta.ExecTime,
		Done:           meta.Done,
		TrackingKey:    meta.TrackingKey,
		RequestDate:    meta.RequestDate,
	}
}

func (r *AdvancedDomainSearchResponse) Decode() *response.AdvancedDomainSearch {
	return &response.AdvancedDomainSearch{
		SearchParams:  decodeAdvancedDomainSearchParams(r.DomainSearch.SearchParams),
		TotalResults:  r.DomainSearch.TotalResults,
		StartPosition: r.DomainSearch.StartPosition,
		NextPosition:  r.DomainSearch.NextPosition,
		MultiRRP:      r.DomainSearch.MultiRRP,
		TLDOverride:   r.DomainSearch.TLDOverride,
		Domains:       decodeAdvancedDomainSearchDomains(r.DomainSearch.Domains),
		ResponseMeta:  decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetAllDomainsResponse) Decode() *response.GetAllDomains {
	return &response.GetAllDomains{
		Domains:           decodeGetAllDomainsDetails(r.GetAllDomains.DomainDetails),
		DomainCount:       r.GetAllDomains.DomainCount,
		UserRequestStatus: r.GetAllDomains.UserRequestStatus,
		ResponseMeta:      decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDomainCountResponse) Decode() *response.GetDomainCount {
	result := response.GetDomainCount{
		RegisteredCount:     r.RegisteredCount,
		HostCount:           r.HostCount,
		ExpiringCount:       r.ExpiringCount,
		ExpiredDomainsCount: r.ExpiredDomainsCount,
		RGP:                 r.RGP,
		ExtendedRGP:         r.ExtendedRGP,
		KeywordCount:        r.KeywordCount,
		ProcessCount:        r.ProcessCount,
		WatchlistCount:      r.WatchlistCount,
		CartItemCount:       r.CartItemCount,
		ResponseMeta:        decodeResponseMeta(r.Response, r.ResponseMeta),
	}
	if r.TrafficMsg != nil {
		result.TrafficMsg = &response.DomainTraffic{
			VistaCustomer:  r.TrafficMsg.VistaCustomer,
			RedirectorData: r.TrafficMsg.RedirectorData,
			Month:          r.TrafficMsg.Month,
			PageViews:      r.TrafficMsg.PageViews,
			Visitors:       r.TrafficMsg.Visitors,
			FreeTrial:      r.TrafficMsg.FreeTrial,
			PDQVista:       r.TrafficMsg.PDQVista,
		}
	}
	return &result
}

func (r *GetDomainExpResponse) Decode() *response.GetDomainExp {
	return &response.GetDomainExp{
		ExpirationDate: r.ExpirationDate,
		ResponseMeta:   decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDomainNameIDResponse) Decode() *response.GetDomainNameID {
	return &response.GetDomainNameID{
		DomainRRP:    r.DomainRRP,
		SLD:          r.SLD,
		TLD:          r.TLD,
		DomainNameID: r.DomainNameID,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDomainsResponse) Decode() *response.GetDomains {
	return &response.GetDomains{
		Tab:              r.GetDomains.Tab,
		MultiRRP:         r.GetDomains.MultiRRP,
		DomainListType:   r.GetDomains.DomainList.Type,
		Domains:          decodeGetDomainsItems(r.GetDomains.DomainList.Domains),
		EndPosition:      r.GetDomains.EndPosition,
		PreviousRecords:  r.GetDomains.PreviousRecords,
		NextRecords:      r.GetDomains.NextRecords,
		OrderBy:          r.GetDomains.OrderBy,
		Result:           r.GetDomains.Result,
		StartPosition:    r.GetDomains.StartPosition,
		DomainCount:      r.GetDomains.DomainCount,
		TotalDomainCount: r.GetDomains.TotalDomainCount,
		StartLetter:      r.GetDomains.StartLetter,
		ResponseMeta:     decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDomainStatusResponse) Decode() *response.GetDomainStatus {
	return &response.GetDomainStatus{
		DomainName:   r.DomainStatus.DomainName,
		Registrar:    r.DomainStatus.Registrar,
		InAccount:    r.DomainStatus.InAccount,
		StatusDesc:   r.DomainStatus.StatusDesc,
		ExpDate:      r.DomainStatus.ExpDate,
		OrderID:      r.DomainStatus.OrderID,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDomainSldTldResponse) Decode() *response.GetDomainSldTld {
	return &response.GetDomainSldTld{
		DomainRRP:    r.DomainRRP,
		SLD:          r.SLD,
		TLD:          r.TLD,
		DomainNameID: r.DomainNameID,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetExpiredDomainsResponse) Decode() *response.GetExpiredDomains {
	return &response.GetExpiredDomains{
		Domains:      decodeExpiredDomainDetails(r.DomainDetails),
		DomainCount:  r.DomainCount,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetExtendInfoResponse) Decode() *response.GetExtendInfo {
	return &response.GetExtendInfo{
		RegistrarHold:       r.RegistrarHold,
		Expiration:          r.Expiration,
		MaxExtension:        r.MaxExtension,
		MinAllowed:          r.MinAllowed,
		CCAuthorized:        r.CCAuthorized,
		Price:               r.Price,
		Balance:             r.Balance,
		AvailableBalance:    r.AvailableBalance,
		CustomerPrefs:       decodeExtendInfoCustomerPrefs(r.CustomerPrefs),
		CustomerInformation: decodeExtendInfoCustomerInformation(r.CustomerInformation),
		ResponseMeta:        decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetNewsResponse) Decode() *response.GetNews {
	return &response.GetNews{
		Alerts:       decodeAlerts(r.Alerts),
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetProductNewsResponse) Decode() *response.GetProductNews {
	return &response.GetProductNews{
		Alerts:       decodeAlerts(r.Alerts),
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetPasswordBitResponse) Decode() *response.GetPasswordBit {
	return &response.GetPasswordBit{
		DomainPassword: r.DomainPassword,
		PasswordSet:    r.PasswordSet,
		ResponseMeta:   decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetRegistrationStatusResponse) Decode() *response.GetRegistrationStatus {
	return &response.GetRegistrationStatus{
		RegistrationStatus: r.RegistrationStatus,
		PurchaseStatus:     r.PurchaseStatus,
		ResponseMeta:       decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetRegLockResponse) Decode() *response.GetRegLock {
	return &response.GetRegLock{
		RegLock:      r.RegLock,
		Registrar:    r.Registrar,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetRenewResponse) Decode() *response.GetRenew {
	return &response.GetRenew{
		RenewName:         r.RenewName,
		PakExist:          r.PakExist,
		AutoPakRenew:      r.AutoPakRenew,
		EmailFwdExists:    r.EmailFwdExists,
		EmailForwardRenew: r.EmailForwardRenew,
		URLFwdExists:      r.URLFwdExists,
		URLForwardRenew:   r.URLForwardRenew,
		IDProtectRenew:    r.IDProtectRenew,
		IDProtectExists:   r.IDProtectExists,
		MobilizerRenew:    r.MobilizerRenew,
		ResponseMeta:      decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetSubAccountPasswordResponse) Decode() *response.GetSubAccountPassword {
	return &response.GetSubAccountPassword{
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *ParseDomainResponse) Decode() *response.ParseDomain {
	return &response.ParseDomain{
		Host:         r.ParseDomain.Host,
		SLD:          r.ParseDomain.SLD,
		TLD:          r.ParseDomain.TLD,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *PortalGetAwardedDomainsResponse) Decode() *response.PortalGetAwardedDomains {
	return &response.PortalGetAwardedDomains{
		Domains:           decodePortalAwardedDomains(r.Domains.Domains),
		DomainCount:       r.Domains.DomainCount,
		TotalDomainCount:  r.Domains.TotalDomainCount,
		NextStartPosition: r.Domains.NextStartPosition,
		Success:           r.Success,
		ResponseMeta:      decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *PortalGetDomainInfoResponse) Decode() *response.PortalGetDomainInfo {
	return &response.PortalGetDomainInfo{
		Domain: response.PortalDomainInfo{
			DomainName:              r.Domain.DomainName,
			EmailAddress:            r.Domain.EmailAddress,
			ExpirationDate:          r.Domain.ExpirationDate,
			RegisterDate:            r.Domain.RegisterDate,
			ForeignLoginID:          r.Domain.ForeignLoginID,
			PortalDomainID:          r.Domain.PortalDomainID,
			RegisterPrice:           r.Domain.RegisterPrice,
			RenewPrice:              r.Domain.RenewPrice,
			RegistrationPeriod:      r.Domain.RegistrationPeriod,
			ResellerProvisioned:     r.Domain.ResellerProvisioned,
			ResellerProvisionedDate: r.Domain.ResellerProvisionedDate,
			IsPremium:               r.Domain.IsPremium,
			RegistrationStatus:      r.Domain.RegistrationStatus,
		},
		TransactionsXML: r.Transactions.Raw,
		ServicesXML:     r.Services.Raw,
		RAASettingsXML:  r.RAASettings.Raw,
		PaymentXML:      r.Payment.Raw,
		ContactsXML:     r.Contacts.Raw,
		ResponseMeta:    decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *PortalGetTokenResponse) Decode() *response.PortalGetToken {
	return &response.PortalGetToken{
		Token:        r.Token,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *PortalUpdateAwardedDomainsResponse) Decode() *response.PortalUpdateAwardedDomains {
	return &response.PortalUpdateAwardedDomains{
		Success:      r.Success,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *RPTGetReportResponse) Decode() *response.RPTGetReport {
	return &response.RPTGetReport{
		Report: response.Report{
			RptString:             r.Report.RptString,
			ThreeMoPast:           r.Report.ThreeMoPast,
			SixMoPast:             r.Report.SixMoPast,
			BeginDate:             r.Report.BeginDate,
			EndDate:               r.Report.EndDate,
			Options:               decodeReportOptions(r.Report.Options),
			ReportType:            r.Report.ReportType,
			ReportName:            r.Report.ReportName,
			QueueReportResults:    r.Report.QueueReportResults,
			ResultsXML:            r.Report.Results.Raw,
			NewStartPosition:      r.Report.NewStartPosition,
			ThisStartPosition:     r.Report.ThisStartPosition,
			ThisEndPosition:       r.Report.ThisEndPosition,
			PreviousStartPosition: r.Report.PreviousStartPosition,
			TotalReturned:         r.Report.TotalReturned,
			TotalRows:             r.Report.TotalRows,
			ShowPrevious:          r.Report.ShowPrevious,
			ShowNext:              r.Report.ShowNext,
			ShowPaging:            r.Report.ShowPaging,
			Version:               r.Report.Version,
		},
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *SetPasswordResponse) Decode() *response.SetPassword {
	return &response.SetPassword{
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *SetRegLockResponse) Decode() *response.SetRegLock {
	return &response.SetRegLock{
		RegLock:      r.RegLock,
		RRPCodeSR:    r.RRPCodeSR,
		RRPText:      r.RRPText,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *SetRenewResponse) Decode() *response.SetRenew {
	return &response.SetRenew{
		RenewName:         r.RenewName,
		AutoPakRenew:      r.AutoPakRenew,
		EmailForwardRenew: r.EmailForwardRenew,
		URLForwardRenew:   r.URLForwardRenew,
		WPPSRenew:         r.WPPSRenew,
		ResponseMeta:      decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *StatusDomainResponse) Decode() *response.StatusDomain {
	return &response.StatusDomain{
		DomainName:   r.DomainStatus.DomainName,
		Registrar:    r.DomainStatus.Registrar,
		InAccount:    r.DomainStatus.InAccount,
		ExpDate:      r.DomainStatus.ExpDate,
		OrderID:      r.DomainStatus.OrderID,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *ValidatePasswordResponse) Decode() *response.ValidatePassword {
	return &response.ValidatePassword{
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func decodeAdvancedDomainSearchParams(params AdvancedDomainSearchParams) response.AdvancedDomainSearchParams {
	return response.AdvancedDomainSearchParams{
		TLDList:             params.TLDList,
		SLD:                 params.SLD,
		SearchCriteria:      params.SearchCriteria,
		ParkingStatus:       params.ParkingStatus,
		RegistrationStatus:  params.RegistrationStatus,
		AutoRenew:           params.AutoRenew,
		Locked:              params.Locked,
		DaysTillExpires:     params.DaysTillExpires,
		DaysExpired:         params.DaysExpired,
		NSStatus:            params.NSStatus,
		NameServer:          params.NameServer,
		HasIDProtect:        params.HasIDProtect,
		HasPOPMail:          params.HasPOPMail,
		EmailForwarding:     params.EmailForwarding,
		HasWebHosting:       params.HasWebHosting,
		ExcludeNumbers:      params.ExcludeNumbers,
		ExcludeDashes:       params.ExcludeDashes,
		IncludeSubAccounts:  params.IncludeSubAccounts,
		SubAccountLogin:     params.SubAccountLogin,
		RecordsToReturn:     params.RecordsToReturn,
		StartPosition:       params.StartPosition,
		OrderBy:             params.OrderBy,
		FolderOption:        params.FolderOption,
		FolderSyncStatus:    params.FolderSyncStatus,
		FolderName:          params.FolderName,
		CreationDate:        params.CreationDate,
		DaysUntilIDPExpires: params.DaysUntilIDPExpires,
		WBLStatusID:         params.WBLStatusID,
		WBLEnabled:          params.WBLEnabled,
		WBLAutoRenew:        params.WBLAutoRenew,
		ContactXML:          params.ContactXML,
		XMLResponse:         params.XMLResponse,
		HostRecordType:      params.HostRecordType,
		HostName:            params.HostName,
		HostAddress:         params.HostAddress,
		EmailResultsOnly:    params.EmailResultsOnly,
		XML:                 params.XML.Raw,
	}
}

func decodeAdvancedDomainSearchDomains(domains []AdvancedDomainSearchDomain) []response.AdvancedDomainSearchDomain {
	if len(domains) == 0 {
		return nil
	}
	result := make([]response.AdvancedDomainSearchDomain, 0, len(domains))
	for _, domain := range domains {
		result = append(result, response.AdvancedDomainSearchDomain{
			DomainNameID:               domain.DomainNameID,
			SLD:                        domain.SLD,
			TLD:                        domain.TLD,
			AutoRenew:                  domain.AutoRenew,
			ExpDate:                    domain.ExpDate,
			DomainRegistrationStatus:   domain.DomainRegistrationStatus,
			DeleteType:                 domain.DeleteType,
			LoginID:                    domain.LoginID,
			AccountID:                  domain.AccountID,
			NSStatus:                   domain.NSStatus,
			FolderStatus:               domain.FolderStatus,
			RRProcessor:                domain.RRProcessor,
			RRCompanyName:              domain.RRCompanyName,
			HasIDProtect:               domain.HasIDProtect,
			IDProtectExpires:           domain.IDProtectExpires,
			DomainFolderStatus:         domain.DomainFolderStatus,
			AbleToReactivate:           domain.AbleToReactivate,
			IsPremiumName:              domain.IsPremiumName,
			PremiumPrice:               domain.PremiumPrice,
			PremiumAboveThresholdPrice: domain.PremiumAboveThresholdPrice,
			PremiumCategory:            domain.PremiumCategory,
			ReactivatePrice:            domain.ReactivatePrice,
			WBLStatusID:                domain.WBLStatusID,
			WBLStatus:                  domain.WBLStatus,
			WBLExpDate:                 domain.WBLExpDate,
			WBLAutoRenew:               domain.WBLAutoRenew,
			NameServers:                domain.NameServers,
			Vas:                        domain.Vas,
		})
	}
	return result
}

func decodeGetAllDomainsDetails(details []GetAllDomainsDetail) []response.GetAllDomainsDetail {
	if len(details) == 0 {
		return nil
	}
	result := make([]response.GetAllDomainsDetail, 0, len(details))
	for _, detail := range details {
		result = append(result, response.GetAllDomainsDetail{
			DomainName:     detail.DomainName,
			DomainNameID:   detail.DomainNameID,
			ExpirationDate: detail.ExpirationDate,
			LockStatus:     detail.LockStatus,
			AutoRenew:      detail.AutoRenew,
		})
	}
	return result
}

func decodeGetDomainsItems(domains []GetDomainsItem) []response.GetDomainsItem {
	if len(domains) == 0 {
		return nil
	}
	result := make([]response.GetDomainsItem, 0, len(domains))
	for _, domain := range domains {
		result = append(result, response.GetDomainsItem{
			DomainNameID:   domain.DomainNameID,
			SLD:            domain.SLD,
			TLD:            domain.TLD,
			NSStatus:       domain.NSStatus,
			ExpirationDate: domain.ExpirationDate,
			AutoRenew:      domain.AutoRenew,
			WPPSStatus:     domain.WPPSStatus,
			WPPSExpDate:    domain.WPPSExpDate,
			RRProcessor:    domain.RRProcessor,
		})
	}
	return result
}

func decodeExpiredDomainDetails(domains []ExpiredDomainDetail) []response.ExpiredDomainDetail {
	if len(domains) == 0 {
		return nil
	}
	result := make([]response.ExpiredDomainDetail, 0, len(domains))
	for _, domain := range domains {
		result = append(result, response.ExpiredDomainDetail{
			DomainName:     domain.DomainName,
			DomainNameID:   domain.DomainNameID,
			Status:         domain.Status,
			ExpirationDate: domain.ExpirationDate,
			LockStatus:     domain.LockStatus,
		})
	}
	return result
}

func decodeExtendInfoCustomerPrefs(prefs ExtendInfoCustomerPrefs) response.ExtendInfoCustomerPrefs {
	return response.ExtendInfoCustomerPrefs{
		DefPeriod:            prefs.DefPeriod,
		AllowDNS:             prefs.AllowDNS,
		ShowPopups:           prefs.ShowPopups,
		AutoRenew:            prefs.AutoRenew,
		RegLock:              prefs.RegLock,
		AutoPakRenew:         prefs.AutoPakRenew,
		UseDNS:               prefs.UseDNS,
		ResellerStatus:       prefs.ResellerStatus,
		RenewalSetting:       prefs.RenewalSetting,
		RenewalBCC:           prefs.RenewalBCC,
		RenewalURLForward:    prefs.RenewalURLForward,
		RenewalEmailForward:  prefs.RenewalEmailForward,
		MailNumLimit:         prefs.MailNumLimit,
		IDProtect:            prefs.IDProtect,
		DefIDProtectRenew:    prefs.DefIDProtectRenew,
		DefWBLRenew:          prefs.DefWBLRenew,
		NameJetSales:         prefs.NameJetSales,
		DefaultHostRecords:   decodeExtendInfoDefaultHostRecords(prefs.DefaultHostRecords),
		DefaultHostRecordOwn: prefs.DefaultHostRecordOwn,
		UseOurDNS:            prefs.UseOurDNS,
		NameServers:          decodeExtendInfoNameServers(prefs.NameServers),
	}
}

func decodeExtendInfoDefaultHostRecords(records ExtendInfoDefaultHostRecords) response.ExtendInfoDefaultHostRecords {
	if len(records.HostRecords) == 0 {
		return response.ExtendInfoDefaultHostRecords{}
	}
	result := response.ExtendInfoDefaultHostRecords{
		HostRecords: make([]response.ExtendInfoHostRecord, 0, len(records.HostRecords)),
	}
	for _, record := range records.HostRecords {
		result.HostRecords = append(result.HostRecords, response.ExtendInfoHostRecord{
			HostName:   record.HostName,
			Address:    record.Address,
			RecordType: record.RecordType,
		})
	}
	return result
}

func decodeExtendInfoNameServers(servers ExtendInfoNameServers) response.ExtendInfoNameServers {
	return response.ExtendInfoNameServers{
		DNS1: servers.DNS1,
		DNS2: servers.DNS2,
		DNS3: servers.DNS3,
		DNS4: servers.DNS4,
		DNS5: servers.DNS5,
	}
}

func decodeExtendInfoCustomerInformation(info ExtendInfoCustomerInformation) response.ExtendInfoCustomerInformation {
	return response.ExtendInfoCustomerInformation{
		AcceptTerms:   info.AcceptTerms,
		URL:           info.URL,
		ParentAccount: info.ParentAccount,
		ParentLogin:   info.ParentLogin,
		NoService:     info.NoService,
		BulkRegLimit:  info.BulkRegLimit,
		Account:       info.Account,
	}
}

func decodeAlerts(alerts Alerts) response.Alerts {
	result := response.Alerts{
		Total: alerts.Total,
	}
	if len(alerts.AlertEntries) == 0 {
		return result
	}
	result.Entries = make([]response.Alert, 0, len(alerts.AlertEntries))
	for _, entry := range alerts.AlertEntries {
		result.Entries = append(result.Entries, response.Alert{
			Product: entry.Product,
			Message: entry.Message,
		})
	}
	return result
}

func decodePortalAwardedDomains(domains []PortalAwardedDomain) []response.PortalAwardedDomain {
	if len(domains) == 0 {
		return nil
	}
	result := make([]response.PortalAwardedDomain, 0, len(domains))
	for _, domain := range domains {
		result = append(result, response.PortalAwardedDomain{
			DomainName:              domain.DomainName,
			EmailAddress:            domain.EmailAddress,
			ExpirationDate:          domain.ExpirationDate,
			RegisterDate:            domain.RegisterDate,
			ForeignLoginID:          domain.ForeignLoginID,
			PortalDomainID:          domain.PortalDomainID,
			RegisterPrice:           domain.RegisterPrice,
			RenewPrice:              domain.RenewPrice,
			RegistrationPeriod:      domain.RegistrationPeriod,
			ResellerProvisioned:     domain.ResellerProvisioned,
			ResellerProvisionedDate: domain.ResellerProvisionedDate,
			IsPremium:               domain.IsPremium,
			RegistrationStatus:      domain.RegistrationStatus,
		})
	}
	return result
}

func decodeReportOptions(options ReportOptions) response.ReportOptions {
	result := response.ReportOptions{}
	if len(options.ReportTypes) == 0 {
		return result
	}
	result.ReportTypes = make([]response.ReportTypeOption, 0, len(options.ReportTypes))
	for _, option := range options.ReportTypes {
		result.ReportTypes = append(result.ReportTypes, response.ReportTypeOption{
			ID:   option.ID,
			Name: option.Name,
		})
	}
	return result
}
