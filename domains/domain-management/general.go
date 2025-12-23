package domainmanagement

import (
	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type Client struct {
	*enomapi.Client
}

// TODO: AdvancedDomainSearch
// TODO: GetAllDomains
// TODO: GetDomainCount
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

// TODO: GetDomainNameID
// TODO: GetDomains
// TODO: GetDomainStatus
// TODO: GetDomainSldTld
// TODO: GetExpiredDomains
// TODO: GetExtendInfo
// TODO: GetHomeDomainList
// TODO: GetNews
// TODO: GetPasswordBit
// TODO: GetProductNews
// TODO: GetRegistrationStatus
// TODO: GetRegLock
// TODO: GetRenew
// TODO: GetSubAccountPassword
// TODO: ParseDomain
// TODO: Portal_GetAwardedDomains
// TODO: Portal_GetDomainInfo
// TODO: Portal_GetToken
// TODO: Portal_UpdateAwardedDomains
// TODO: RPT_GetReport
// TODO: SetPassword
// TODO: SetRegLock
// TODO: SetRenew
// TODO: StatusDomain
// TODO: ValidatePassword
