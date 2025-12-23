package domainmanagement

import (
	"strconv"
	"strings"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type Client struct {
	*enomapi.Client
}

type AdvancedDomainSearchRequest struct {
	TLDList                   string
	SearchCriteria            string
	SLD                       string
	ExcludeNumbers            *bool
	ExcludeDashes             *bool
	ParkingStatus             *bool
	RegistrationStatus        string
	AutoRenew                 *bool
	Locked                    *bool
	CreationDate              string
	DaysTillExpires           *int
	DaysExpired               *int
	NSStatus                  string
	NameServer                string
	HostRecordType            string
	HostName                  string
	HostAddress               string
	HasIDProtect              *bool
	DaysUntilIDProtectExpires *int
	HasPOPMail                *bool
	HasWebHosting             *bool
	IncludeSubaccounts        *bool
	SubaccountLogin           string
	OrderBy                   string
	StartPosition             *int
	RecordsToReturn           *int
	FolderOption              *int
	FolderName                string
	FolderSyncStatus          *int
	MultiLang                 string
}

type GetAllDomainsRequest struct {
	UseDNS         string
	GetDefaultOnly string
	Letter         string
	RegistrarLock  string
	AutoRenew      string
	NameServer     string
	UseEnomNS      string
}

type GetDomainsRequest struct {
	Tab           string
	DaysToExpired *int
	RegStatus     string
	Display       *int
	Start         *int
	OrderBy       string
	StartLetter   string
	MultiLang     string
	Domain        string
	ExtFormat     *int
}

type PortalGetAwardedDomainsRequest struct {
	RecordCount   *int
	RecordOption  *int
	StartPosition *int
	TLD           string
	DomainNameID  string
}

type PortalGetTokenRequest struct {
	PortalUserID string
	Email        string
}

type PortalUpdateAwardedDomainsRequest struct {
	DomainList string
	DomainIDs  []int64
}

type RPTGetReportRequest struct {
	Version          *int
	ReportType       int
	BeginDate        string
	EndDate          string
	Start            *int
	RecordsToReturn  *int
	Download         bool
	ReportOutputType string
}

// AdvancedDomainSearch searches for domains in the account.
func (c Client) AdvancedDomainSearch(req AdvancedDomainSearchRequest) (*response.AdvancedDomainSearch, error) {
	resp := internal.AdvancedDomainSearchResponse{}

	cmd := c.NewCommand("AdvancedDomainSearch")
	if req.TLDList != "" {
		cmd.AddParam("TLDList", req.TLDList)
	}
	if req.SearchCriteria != "" {
		cmd.AddParam("SearchCriteria", req.SearchCriteria)
	}
	if req.SLD != "" {
		cmd.AddParam("SLD", req.SLD)
	}
	addBoolIntParam(cmd, "ExcludeNumbers", req.ExcludeNumbers)
	addBoolIntParam(cmd, "ExcludeDashes", req.ExcludeDashes)
	addBoolIntParam(cmd, "ParkingStatus", req.ParkingStatus)
	if req.RegistrationStatus != "" {
		cmd.AddParam("RegistrationStatus", req.RegistrationStatus)
	}
	addBoolIntParam(cmd, "AutoRenew", req.AutoRenew)
	addBoolIntParam(cmd, "Locked", req.Locked)
	if req.CreationDate != "" {
		cmd.AddParam("CreationDate", req.CreationDate)
	}
	addIntParam(cmd, "DaysTillExpires", req.DaysTillExpires)
	addIntParam(cmd, "DaysExpired", req.DaysExpired)
	if req.NSStatus != "" {
		cmd.AddParam("NSStatus", req.NSStatus)
	}
	if req.NameServer != "" {
		cmd.AddParam("NameServer", req.NameServer)
	}
	if req.HostRecordType != "" {
		cmd.AddParam("HostRecordType", req.HostRecordType)
	}
	if req.HostName != "" {
		cmd.AddParam("HostName", req.HostName)
	}
	if req.HostAddress != "" {
		cmd.AddParam("HostAddress", req.HostAddress)
	}
	addBoolIntParam(cmd, "HasIDProtect", req.HasIDProtect)
	addIntParam(cmd, "DaysUntilExpires", req.DaysUntilIDProtectExpires)
	addBoolIntParam(cmd, "HasPOPMail", req.HasPOPMail)
	addBoolIntParam(cmd, "HasWebHosting", req.HasWebHosting)
	addBoolIntParam(cmd, "IncludeSubaccounts", req.IncludeSubaccounts)
	if req.SubaccountLogin != "" {
		cmd.AddParam("SubaccountLogin", req.SubaccountLogin)
	}
	if req.OrderBy != "" {
		cmd.AddParam("OrderBy", req.OrderBy)
	}
	addIntParam(cmd, "StartPosition", req.StartPosition)
	addIntParam(cmd, "RecordsToReturn", req.RecordsToReturn)
	addIntParam(cmd, "FolderOption", req.FolderOption)
	if req.FolderName != "" {
		cmd.AddParam("FolderName", req.FolderName)
	}
	addIntParam(cmd, "FolderSyncStatus", req.FolderSyncStatus)
	if req.MultiLang != "" {
		cmd.AddParam("MultiLang", req.MultiLang)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetAllDomains returns all domains in the account.
func (c Client) GetAllDomains(req GetAllDomainsRequest) (*response.GetAllDomains, error) {
	resp := internal.GetAllDomainsResponse{}

	cmd := c.NewCommand("GetAllDomains")
	if req.UseDNS != "" {
		cmd.AddParam("UseDNS", req.UseDNS)
	}
	if req.GetDefaultOnly != "" {
		cmd.AddParam("GetDefaultOnly", req.GetDefaultOnly)
	}
	if req.Letter != "" {
		cmd.AddParam("Letter", req.Letter)
	}
	if req.RegistrarLock != "" {
		cmd.AddParam("RegistrarLock", req.RegistrarLock)
	}
	if req.AutoRenew != "" {
		cmd.AddParam("AutoRenew", req.AutoRenew)
	}
	if req.NameServer != "" {
		cmd.AddParam("NameServer", req.NameServer)
	}
	if req.UseEnomNS != "" {
		cmd.AddParam("UseEnomNS", req.UseEnomNS)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDomainCount returns counts for domain categories in the account.
func (c Client) GetDomainCount() (*response.GetDomainCount, error) {
	resp := internal.GetDomainCountResponse{}

	cmd := c.NewCommand("GetDomainCount")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// TODO: GetDomainExp
// TODO: GetDomainInfo
func (c Client) GetDomainInfo(domain enomapi.Domain) (*response.GetDomainInfo, error) {
	resp := internal.GetDomainInfoResponse{}

	cmd := c.NewCommand("GetDomainInfo")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDomainNameID returns the ID number for a domain.
func (c Client) GetDomainNameID(domain enomapi.Domain) (*response.GetDomainNameID, error) {
	resp := internal.GetDomainNameIDResponse{}

	cmd := c.NewCommand("GetDomainNameID")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDomains returns a page of domains in the account.
func (c Client) GetDomains(req GetDomainsRequest) (*response.GetDomains, error) {
	resp := internal.GetDomainsResponse{}

	cmd := c.NewCommand("GetDomains")
	if req.Tab != "" {
		cmd.AddParam("Tab", req.Tab)
	}
	addIntParam(cmd, "DaysToExpired", req.DaysToExpired)
	if req.RegStatus != "" {
		cmd.AddParam("RegStatus", req.RegStatus)
	}
	addIntParam(cmd, "Display", req.Display)
	addIntParam(cmd, "Start", req.Start)
	if req.OrderBy != "" {
		cmd.AddParam("OrderBy", req.OrderBy)
	}
	if req.StartLetter != "" {
		cmd.AddParam("StartLetter", req.StartLetter)
	}
	if req.MultiLang != "" {
		cmd.AddParam("MultiLang", req.MultiLang)
	}
	if req.Domain != "" {
		cmd.AddParam("Domain", req.Domain)
	}
	addIntParam(cmd, "ExtFormat", req.ExtFormat)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDomainStatus returns the status for a non-real-time TLD order.
func (c Client) GetDomainStatus(domain enomapi.Domain, orderID string, orderType string) (*response.GetDomainStatus, error) {
	resp := internal.GetDomainStatusResponse{}

	cmd := c.NewCommand("GetDomainStatus")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	if orderID != "" {
		cmd.AddParam("OrderID", orderID)
	}
	if orderType != "" {
		cmd.AddParam("OrderType", orderType)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDomainSldTld returns the SLD and TLD for a domain ID.
func (c Client) GetDomainSldTld(domainNameID int64) (*response.GetDomainSldTld, error) {
	resp := internal.GetDomainSldTldResponse{}

	cmd := c.NewCommand("GetDomainSLDTLD")
	cmd.AddParam("DomainNameID", strconv.FormatInt(domainNameID, 10))

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetExpiredDomains returns domains in expired or redemption status.
func (c Client) GetExpiredDomains() (*response.GetExpiredDomains, error) {
	resp := internal.GetExpiredDomainsResponse{}

	cmd := c.NewCommand("GetExpiredDomains")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// TODO: GetExtendInfo
// TODO: GetHomeDomainList
// GetNews returns registry maintenance alerts.
func (c Client) GetNews() (*response.GetNews, error) {
	resp := internal.GetNewsResponse{}

	cmd := c.NewCommand("GetNews")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetPasswordBit returns domain password details.
func (c Client) GetPasswordBit(domain enomapi.Domain) (*response.GetPasswordBit, error) {
	resp := internal.GetPasswordBitResponse{}

	cmd := c.NewCommand("GetPasswordBit")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetProductNews returns product maintenance alerts.
func (c Client) GetProductNews() (*response.GetProductNews, error) {
	resp := internal.GetProductNewsResponse{}

	cmd := c.NewCommand("GetProductNews")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetRegistrationStatus returns registration and purchase status for a domain.
func (c Client) GetRegistrationStatus(domain enomapi.Domain) (*response.GetRegistrationStatus, error) {
	resp := internal.GetRegistrationStatusResponse{}

	cmd := c.NewCommand("GetRegistrationStatus")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetRegLock returns registrar lock status for a domain.
func (c Client) GetRegLock(domain enomapi.Domain) (*response.GetRegLock, error) {
	resp := internal.GetRegLockResponse{}

	cmd := c.NewCommand("GetRegLock")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// TODO: GetRenew
// GetSubAccountPassword emails the domain password to the registrant.
func (c Client) GetSubAccountPassword(domain enomapi.Domain) (*response.GetSubAccountPassword, error) {
	resp := internal.GetSubAccountPasswordResponse{}

	cmd := c.NewCommand("GetSubAccountPassword")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// ParseDomain splits a domain into host, SLD, and TLD.
func (c Client) ParseDomain(passedDomain string) (*response.ParseDomain, error) {
	resp := internal.ParseDomainResponse{}

	cmd := c.NewCommand("ParseDomain")
	cmd.AddParam("PassedDomain", passedDomain)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// PortalGetAwardedDomains returns portal-awarded domains.
func (c Client) PortalGetAwardedDomains(req PortalGetAwardedDomainsRequest) (*response.PortalGetAwardedDomains, error) {
	resp := internal.PortalGetAwardedDomainsResponse{}

	cmd := c.NewCommand("Portal_GetAwardedDomains")
	addIntParam(cmd, "RecordCount", req.RecordCount)
	addIntParam(cmd, "RecordOption", req.RecordOption)
	addIntParam(cmd, "StartPosition", req.StartPosition)
	if req.TLD != "" {
		cmd.AddParam("Tld", req.TLD)
	}
	if req.DomainNameID != "" {
		cmd.AddParam("DomainNameID", req.DomainNameID)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// PortalGetDomainInfo returns portal-specific domain info.
func (c Client) PortalGetDomainInfo(domain enomapi.Domain) (*response.PortalGetDomainInfo, error) {
	resp := internal.PortalGetDomainInfoResponse{}

	cmd := c.NewCommand("Portal_GetDomainInfo")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// PortalGetToken requests an access token for the portal.
func (c Client) PortalGetToken(req PortalGetTokenRequest) (*response.PortalGetToken, error) {
	resp := internal.PortalGetTokenResponse{}

	cmd := c.NewCommand("Portal_GetToken")
	cmd.AddParam("PortalUserID", req.PortalUserID)
	cmd.AddParam("Email", req.Email)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// PortalUpdateAwardedDomains marks portal domains as provisioned.
func (c Client) PortalUpdateAwardedDomains(req PortalUpdateAwardedDomainsRequest) (*response.PortalUpdateAwardedDomains, error) {
	resp := internal.PortalUpdateAwardedDomainsResponse{}

	cmd := c.NewCommand("Portal_UpdateAwardedDomains")
	domainList := req.DomainList
	if domainList == "" && len(req.DomainIDs) > 0 {
		ids := make([]string, 0, len(req.DomainIDs))
		for _, id := range req.DomainIDs {
			ids = append(ids, strconv.FormatInt(id, 10))
		}
		domainList = strings.Join(ids, ",")
	}
	if domainList != "" {
		cmd.AddParam("domainlist", domainList)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// RPTGetReport requests an itemized activity report.
func (c Client) RPTGetReport(req RPTGetReportRequest) (*response.RPTGetReport, error) {
	resp := internal.RPTGetReportResponse{}

	cmd := c.NewCommand("RPT_GetReport")
	addIntParam(cmd, "Version", req.Version)
	cmd.AddParam("ReportType", strconv.Itoa(req.ReportType))
	if req.BeginDate != "" {
		cmd.AddParam("BeginDate", req.BeginDate)
	}
	if req.EndDate != "" {
		cmd.AddParam("EndDate", req.EndDate)
	}
	addIntParam(cmd, "Start", req.Start)
	addIntParam(cmd, "RecordsToReturn", req.RecordsToReturn)
	cmd.AddParam("Download", boolToTitle(req.Download))
	if req.ReportOutputType != "" {
		cmd.AddParam("ReportOutputType", req.ReportOutputType)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// SetPassword sets the domain password.
func (c Client) SetPassword(domain enomapi.Domain, password string) (*response.SetPassword, error) {
	resp := internal.SetPasswordResponse{}

	cmd := c.NewCommand("SetPassword")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("DomainPassword", password)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// SetRegLock sets the registrar lock for a domain.
func (c Client) SetRegLock(domain enomapi.Domain, unlockRegistrar bool) (*response.SetRegLock, error) {
	resp := internal.SetRegLockResponse{}

	cmd := c.NewCommand("SetRegLock")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("UnlockRegistrar", boolToInt(unlockRegistrar))

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// TODO: SetRenew
// StatusDomain returns basic status for a domain.
func (c Client) StatusDomain(domain enomapi.Domain, orderType string) (*response.StatusDomain, error) {
	resp := internal.StatusDomainResponse{}

	cmd := c.NewCommand("StatusDomain")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	if orderType != "" {
		cmd.AddParam("OrderType", orderType)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// ValidatePassword validates a domain password.
func (c Client) ValidatePassword(domain enomapi.Domain, password string) (*response.ValidatePassword, error) {
	resp := internal.ValidatePasswordResponse{}

	cmd := c.NewCommand("ValidatePassword")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("DomainPassword", password)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

func addBoolIntParam(cmd *enomapi.Command, key string, value *bool) {
	if value == nil {
		return
	}
	cmd.AddParam(key, boolToInt(*value))
}

func addIntParam(cmd *enomapi.Command, key string, value *int) {
	if value == nil {
		return
	}
	cmd.AddParam(key, strconv.Itoa(*value))
}

func boolToInt(value bool) string {
	if value {
		return "1"
	}
	return "0"
}

func boolToTitle(value bool) string {
	if value {
		return "True"
	}
	return "False"
}
