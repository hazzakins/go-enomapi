package domains

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type AddBulkDomainsItem struct {
	Domain   enomapi.Domain
	NumYears *int
}

type AddBulkDomainsRequest struct {
	ProductType string
	Items       []AddBulkDomainsItem
	AutoRenew   *bool
	RegLock     *bool
	UseCart     *bool
}

type PreconfigureRequest struct {
	Domain             enomapi.Domain
	Load               int
	ExtendedAttributes map[string]string
	AutoRenew          int
	RegLock            int
	IDNCodes           []string
	OptContactReg      string
	OptTechnical       string
	OptAdministrative  string
	OptContactAux      string
	PreConfigDNS       string
	NameServers        []string
	UseHostRecords     *int
	HostNames          []string
	RecordTypes        []string
	Addresses          []string
	ContactParams      map[string]string
	AccessPassword1    string
	AccessPassword2    string
}

type PurchaseRequest struct {
	Domain             enomapi.Domain
	UseDNS             string
	NameServers        []string
	ExtendedAttributes map[string]string
}

// Upstream documentation: https://api.enom.com/docs/addbulkdomains
func (c *Client) AddBulkDomains(req AddBulkDomainsRequest) (*response.AddBulkDomains, error) {
	resp := internal.AddBulkDomainsResponse{}

	cmd := c.NewCommand("AddBulkDomains")
	cmd.AddParam("ProductType", req.ProductType)
	cmd.AddParam("ListCount", strconv.Itoa(len(req.Items)))
	for i, item := range req.Items {
		index := strconv.Itoa(i + 1)
		cmd.AddParam("SLD"+index, item.Domain.Name)
		cmd.AddParam("TLD"+index, item.Domain.Extension)
		if item.NumYears != nil {
			cmd.AddParam("NumYears"+index, strconv.Itoa(*item.NumYears))
		}
	}
	addBoolIntParam(cmd, "AutoRenew", req.AutoRenew)
	addBoolIntParam(cmd, "RegLock", req.RegLock)
	addBoolIntParam(cmd, "UseCart", req.UseCart)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/preconfigure
func (c *Client) Preconfigure(req PreconfigureRequest) (*response.Preconfigure, error) {
	resp := internal.PreconfigureResponse{}

	cmd := c.NewCommand("Preconfigure")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)
	cmd.AddParam("Load", strconv.Itoa(req.Load))
	cmd.AddParam("AutoRenew", strconv.Itoa(req.AutoRenew))
	cmd.AddParam("RegLock", strconv.Itoa(req.RegLock))
	if len(req.IDNCodes) > 0 {
		for i, code := range req.IDNCodes {
			cmd.AddParam("IDN"+strconv.Itoa(i+1), code)
		}
	}
	if req.OptContactReg != "" {
		cmd.AddParam("OptContactReg", req.OptContactReg)
	}
	if req.OptTechnical != "" {
		cmd.AddParam("OptTechnical", req.OptTechnical)
	}
	if req.OptAdministrative != "" {
		cmd.AddParam("OptAdministrative", req.OptAdministrative)
	}
	if req.OptContactAux != "" {
		cmd.AddParam("OptContactAux", req.OptContactAux)
	}
	if req.PreConfigDNS != "" {
		cmd.AddParam("PreConfigDNS", req.PreConfigDNS)
	}
	if len(req.NameServers) > 0 {
		for i, ns := range req.NameServers {
			cmd.AddParam("NS"+strconv.Itoa(i+1), ns)
		}
	}
	addIntParam(cmd, "UseHostRecords", req.UseHostRecords)
	if len(req.HostNames) > 0 {
		cmd.AddParam("HostName", strings.Join(req.HostNames, ","))
	}
	if len(req.RecordTypes) > 0 {
		cmd.AddParam("RecordType", strings.Join(req.RecordTypes, ","))
	}
	if len(req.Addresses) > 0 {
		cmd.AddParam("Address", strings.Join(req.Addresses, ","))
	}
	if req.AccessPassword1 != "" {
		cmd.AddParam("AccessPassword1", req.AccessPassword1)
	}
	if req.AccessPassword2 != "" {
		cmd.AddParam("AccessPassword2", req.AccessPassword2)
	}
	for key, value := range req.ExtendedAttributes {
		cmd.AddParam(key, value)
	}
	for key, value := range req.ContactParams {
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/purchase
func (c Client) Purchase(domain enomapi.Domain) (*response.DomainPurchase, error) {
	resp := internal.DomainPurchase{}

	tldID, err := c.PEGetTLDID(domain.Extension)
	if err != nil {
		return nil, fmt.Errorf("pe_gettldid: %w", err)
	}
	if tldID.TLDID == "" {
		return nil, fmt.Errorf("pe_gettldid: empty TLDID")
	}

	cmd := c.NewCommand("Purchase")
	cmd.AddParam("SLD", domain.Name)
	cmd.AddParam("TLD", domain.Extension)

	err = c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Decode()
}

// PurchaseWithOptions purchases a domain with optional name server settings.
func (c Client) PurchaseWithOptions(req PurchaseRequest) (*response.DomainPurchase, error) {
	resp := internal.DomainPurchase{}

	tldID, err := c.PEGetTLDID(req.Domain.Extension)
	if err != nil {
		return nil, fmt.Errorf("pe_gettldid: %w", err)
	}
	if tldID.TLDID == "" {
		return nil, fmt.Errorf("pe_gettldid: empty TLDID")
	}

	cmd := c.NewCommand("Purchase")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)
	if req.UseDNS != "" {
		cmd.AddParam("UseDNS", req.UseDNS)
	}
	if len(req.NameServers) > 0 {
		for i, ns := range req.NameServers {
			cmd.AddParam("NS"+strconv.Itoa(i+1), ns)
		}
	}
	for key, value := range req.ExtendedAttributes {
		cmd.AddParam(key, value)
	}

	err = c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Decode()
}

// Upstream documentation: https://api.enom.com/docs/DeleteRegistration
func (c *Client) DeleteRegistration(domain enomapi.Domain, endUserIP string) (*response.DeleteRegistration, error) {
	resp := internal.DeleteRegistrationResponse{}

	cmd := c.NewCommand("DeleteRegistration")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("EndUserIP", endUserIP)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
