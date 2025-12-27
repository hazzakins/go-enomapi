package domainmanagement

import (
	"strconv"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

// GetDomainSRVHosts retrieves SRV host records for a domain name.
func (c Client) GetDomainSRVHosts(domain enomapi.Domain) (*response.GetDomainSRVHosts, error) {
	resp := internal.GetDomainSRVHostsResponse{}

	cmd := c.NewCommand("GetDomainSRVHosts")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// UpdateMetaTagRequest defines input for the UpdateMetaTag command.
type UpdateMetaTagRequest struct {
	Domain          enomapi.Domain
	MetaTagHostID   int64
	TitleBar        string
	SiteDescription string
	Keywords        string
}

// UpdateMetaTag updates HTML metatags for a domain host record.
func (c Client) UpdateMetaTag(req UpdateMetaTagRequest) (*response.UpdateMetaTag, error) {
	resp := internal.UpdateMetaTagResponse{}

	cmd := c.NewCommand("UpdateMetaTag")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	cmd.AddParam("MetaTagHostID", strconv.FormatInt(req.MetaTagHostID, 10))
	if req.TitleBar != "" {
		cmd.AddParam("TitleBar", req.TitleBar)
	}
	if req.SiteDescription != "" {
		cmd.AddParam("SiteDescription", req.SiteDescription)
	}
	if req.Keywords != "" {
		cmd.AddParam("Keywords", req.Keywords)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// SetSPFHostsRequest defines input for the SetSPFHosts command.
type SetSPFHostsRequest struct {
	Domain     enomapi.Domain
	HostID     *int64
	HostName   string
	RecordType string
	Address    string
	MXPref     *int
}

// SetSPFHosts creates or updates an SPF host record for a domain.
func (c Client) SetSPFHosts(req SetSPFHostsRequest) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetSPFHosts")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	if req.HostID != nil {
		cmd.AddParam("HostID", strconv.FormatInt(*req.HostID, 10))
	}
	if req.HostName != "" {
		cmd.AddParam("HostName", req.HostName)
	}
	if req.RecordType != "" {
		cmd.AddParam("RecordType", req.RecordType)
	}
	if req.Address != "" {
		cmd.AddParam("Address", req.Address)
	}
	if req.MXPref != nil {
		cmd.AddParam("MXPref", strconv.Itoa(*req.MXPref))
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// SetHostsRequest defines input for the SetHosts command.
type SetHostsRequest struct {
	Domain      enomapi.Domain
	HostRecords []SetHostRecord
}

// SetHostRecord defines a host record used in SetHosts.
type SetHostRecord struct {
	HostName   string
	RecordType string
	Address    string
	MXPref     *int
}

// SetHosts replaces a domain's host records with the provided set.
func (c Client) SetHosts(req SetHostsRequest) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetHosts")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	for i, record := range req.HostRecords {
		index := strconv.Itoa(i + 1)
		if record.HostName != "" {
			cmd.AddParam("HostName"+index, record.HostName)
		}
		if record.RecordType != "" {
			cmd.AddParam("RecordType"+index, record.RecordType)
		}
		if record.Address != "" {
			cmd.AddParam("Address"+index, record.Address)
		}
		if record.MXPref != nil {
			cmd.AddParam("MXPref"+index, strconv.Itoa(*record.MXPref))
		}
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// SetDomainSRVHostsRequest defines input for the SetDomainSRVHosts command.
type SetDomainSRVHostsRequest struct {
	Domain      enomapi.Domain
	HostRecords []SetDomainSRVHostRecord
}

// SetDomainSRVHostRecord defines an SRV host record used in SetDomainSRVHosts.
type SetDomainSRVHostRecord struct {
	HostID   *int64
	Service  string
	Protocol string
	Priority int
	Weight   int
	Port     int
	Target   string
}

// SetDomainSRVHosts replaces a domain's SRV host records with the provided set.
func (c Client) SetDomainSRVHosts(req SetDomainSRVHostsRequest) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetDomainSRVHosts")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	for i, record := range req.HostRecords {
		index := strconv.Itoa(i + 1)
		if record.HostID != nil {
			cmd.AddParam("HostID"+index, strconv.FormatInt(*record.HostID, 10))
		}
		cmd.AddParam("Service"+index, record.Service)
		cmd.AddParam("Protocol"+index, record.Protocol)
		cmd.AddParam("Priority"+index, strconv.Itoa(record.Priority))
		cmd.AddParam("Weight"+index, strconv.Itoa(record.Weight))
		cmd.AddParam("Port"+index, strconv.Itoa(record.Port))
		cmd.AddParam("Target"+index, record.Target)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetHosts retrieves host records for a domain name.
func (c Client) GetHosts(domain enomapi.Domain) (*response.GetHosts, error) {
	resp := internal.GetHostsResponse{}

	cmd := c.NewCommand("GetHosts")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetMetaTagRequest defines input for the GetMetaTag command.
type GetMetaTagRequest struct {
	Domain        enomapi.Domain
	MetaTagHostID int64
}

// GetMetaTag retrieves HTML metatags for a domain host record.
func (c Client) GetMetaTag(req GetMetaTagRequest) (*response.UpdateMetaTag, error) {
	resp := internal.UpdateMetaTagResponse{}

	cmd := c.NewCommand("GetMetaTag")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	cmd.AddParam("MetaTagHostID", strconv.FormatInt(req.MetaTagHostID, 10))

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetRegHosts retrieves non-mail host records for a domain name.
func (c Client) GetRegHosts(domain enomapi.Domain) (*response.GetRegHosts, error) {
	resp := internal.GetRegHostsResponse{}

	cmd := c.NewCommand("GetRegHosts")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("ExtFormat", "1")

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetSPFHosts retrieves SPF host records for a domain name.
func (c Client) GetSPFHosts(domain enomapi.Domain) (*response.GetSPFHosts, error) {
	resp := internal.GetSPFHostsResponse{}

	cmd := c.NewCommand("GetSPFHosts")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
