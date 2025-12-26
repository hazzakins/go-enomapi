package domainmanagement

import (
	"strconv"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

// AddDnsSecRequest defines input for the AddDnsSec command.
type AddDnsSecRequest struct {
	Domain         enomapi.Domain
	SetMaxLifeOnly *bool
	MaxSigLife     *int
	Algorithm      int
	Digest         string
	DigestType     int
	KeyTag         int
}

// DeleteDnsSecRequest defines input for the DeleteDnsSec command.
type DeleteDnsSecRequest struct {
	Domain     enomapi.Domain
	Algorithm  int
	Digest     string
	DigestType int
	KeyTag     int
}

// GetHomeDomainListRequest defines paging options for GetHomeDomainList.
type GetHomeDomainListRequest struct {
	StartPosition *int
	Display       *int
	OrderBy       string
}

// ModifyNSRequest defines input for the ModifyNS command.
type ModifyNSRequest struct {
	Domain      enomapi.Domain
	UseDNS      string
	NameServers []string
}

// ModifyNSHostingRequest defines input for the ModifyNSHosting command.
type ModifyNSHostingRequest struct {
	Domain      enomapi.Domain
	NameServers []string
}

// SetDNSHostRequest defines input for the SetDNSHost command.
type SetDNSHostRequest struct {
	Zone           string
	DomainPassword string
	Address        string
}

// UpdateNameServerRequest defines input for the UpdateNameServer command.
type UpdateNameServerRequest struct {
	NameServer string
	OldIP      string
	NewIP      string
}

// AddDnsSec creates a DNSSEC record for a domain.
func (c Client) AddDnsSec(req AddDnsSecRequest) (*response.DnsSecUpdate, error) {
	resp := internal.AddDnsSecResponse{}

	cmd := c.NewCommand("AddDnsSec")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	if req.SetMaxLifeOnly != nil {
		cmd.AddParam("SetMaxLifeOnly", strconv.FormatBool(*req.SetMaxLifeOnly))
	}
	addIntParam(cmd, "MaxSigLife", req.MaxSigLife)
	cmd.AddParam("Alg", strconv.Itoa(req.Algorithm))
	cmd.AddParam("Digest", req.Digest)
	cmd.AddParam("DigestType", strconv.Itoa(req.DigestType))
	cmd.AddParam("KeyTag", strconv.Itoa(req.KeyTag))

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// CheckNSStatus retrieves status information for a name server.
func (c Client) CheckNSStatus(nameServer string) (*response.CheckNSStatus, error) {
	resp := internal.CheckNSStatusResponse{}

	cmd := c.NewCommand("CheckNSStatus")
	cmd.AddParam("CheckNSName", nameServer)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// DeleteDnsSec removes a DNSSEC record from a domain.
func (c Client) DeleteDnsSec(req DeleteDnsSecRequest) (*response.DnsSecUpdate, error) {
	resp := internal.DeleteDnsSecResponse{}

	cmd := c.NewCommand("DeleteDnsSec")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	cmd.AddParam("Alg", strconv.Itoa(req.Algorithm))
	cmd.AddParam("Digest", req.Digest)
	cmd.AddParam("DigestType", strconv.Itoa(req.DigestType))
	cmd.AddParam("KeyTag", strconv.Itoa(req.KeyTag))

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// DeleteNameServer deletes a registered name server.
func (c Client) DeleteNameServer(nameServer string) (*response.DeleteNameServer, error) {
	resp := internal.DeleteNameServerResponse{}

	cmd := c.NewCommand("DeleteNameServer")
	cmd.AddParam("NS", nameServer)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDNS returns name server settings for a domain name.
func (c Client) GetDNS(domain enomapi.Domain) (*response.GetDNS, error) {
	resp := internal.GetDNSResponse{}

	cmd := c.NewCommand("GetDNS")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDnsSec returns DNSSEC keys for a domain name.
func (c Client) GetDnsSec(domain enomapi.Domain) (*response.GetDnsSec, error) {
	resp := internal.GetDnsSecResponse{}

	cmd := c.NewCommand("GetDnsSec")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetDNSStatus returns DNS status for a domain name.
func (c Client) GetDNSStatus(domain enomapi.Domain) (*response.GetDNSStatus, error) {
	resp := internal.GetDNSStatusResponse{}

	cmd := c.NewCommand("GetDNSStatus")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// ModifyNS updates the name servers for a domain.
func (c Client) ModifyNS(req ModifyNSRequest) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("ModifyNS")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	if req.UseDNS != "" {
		cmd.AddParam("UseDNS", req.UseDNS)
	}
	addNameServers(cmd, req.NameServers)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// ModifyNSHosting updates the registrar's name server settings without registry changes.
func (c Client) ModifyNSHosting(req ModifyNSHostingRequest) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("ModifyNSHosting")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	addNameServers(cmd, req.NameServers)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// RegisterNameServer registers a new name server.
func (c Client) RegisterNameServer(nameServer string, ipAddress string) (*response.RegisterNameServer, error) {
	resp := internal.RegisterNameServerResponse{}

	cmd := c.NewCommand("RegisterNameServer")
	cmd.AddParam("Add", "true")
	cmd.AddParam("NSName", nameServer)
	cmd.AddParam("IP", ipAddress)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// SetDNSHost updates a dynamic DNS host record.
func (c Client) SetDNSHost(req SetDNSHostRequest) (*response.SetDNSHost, error) {
	resp := internal.SetDNSHostResponse{}

	cmd := c.NewCommand("SetDNSHost")
	cmd.AddParam("zone", req.Zone)
	cmd.AddParam("domainpassword", req.DomainPassword)
	if req.Address != "" {
		cmd.AddParam("address", req.Address)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// UpdateNameServer updates a registered name server's IP address.
func (c Client) UpdateNameServer(req UpdateNameServerRequest) (*response.UpdateNameServer, error) {
	resp := internal.UpdateNameServerResponse{}

	cmd := c.NewCommand("UpdateNameServer")
	cmd.AddParam("OldIP", req.OldIP)
	cmd.AddParam("NewIP", req.NewIP)
	cmd.AddParam("NS", req.NameServer)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

func addNameServers(cmd *enomapi.Command, nameServers []string) {
	for i, nameServer := range nameServers {
		cmd.AddParam("NS"+strconv.Itoa(i+1), nameServer)
	}
}
