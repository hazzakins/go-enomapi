package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

type CheckNSStatusResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	NsCheckSuccess string        `xml:"NsCheckSuccess"`
	CheckNsStatus  CheckNsStatus `xml:"CheckNsStatus"`
}

type CheckNsStatus struct {
	Name         string   `xml:"name"`
	AttribID     string   `xml:"attrib-id"`
	IPAddress    string   `xml:"ipaddress"`
	AttribUPID   string   `xml:"attrib-upid"`
	AttribCLID   string   `xml:"attrib-clid"`
	AttribCRID   string   `xml:"attrib-crid"`
	Status       []string `xml:"status>status"`
	AttribUpdate string   `xml:"attrib-update"`
	AttribCRDate string   `xml:"attrib-crdate"`
}

type GetDNSResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DNS                      []string `xml:"dns"`
	UseDNS                   string   `xml:"UseDNS"`
	NSStatus                 string   `xml:"NSStatus"`
	HostsNumLimit            int      `xml:"HostsNumLimit"`
	DNSRegistrySynced        string   `xml:"DNSRegistrySynced"`
	NameserverRegistrySynced string   `xml:"NameserverRegistrySynced"`
	RRPCodeGDNS              int      `xml:"RRPCodeGDNS"`
}

type GetDNSStatusResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	UseDNS        string `xml:"UseDNS"`
	NSStatus      string `xml:"NSStatus"`
	HostsNumLimit int    `xml:"HostsNumLimit"`
}

type GetHomeDomainListResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	GetHomeDomains GetHomeDomains `xml:"GetHomeDomains"`
}

type GetHomeDomains struct {
	StartPosition    int          `xml:"StartPosition"`
	Display          int          `xml:"Display"`
	OrderBy          string       `xml:"OrderBy"`
	NewStartPosition int          `xml:"NewStartPosition"`
	DomainCount      int          `xml:"DomainCount"`
	TotalDomains     int          `xml:"TotalDomains"`
	Domains          []HomeDomain `xml:"Domains>Domain"`
}

type HomeDomain struct {
	ID        int64  `xml:"ID"`
	Name      string `xml:"Name"`
	RegStatus string `xml:"RegStatus"`
}

type RegisterNameServerResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	RegisterNameServer NameServerResult `xml:"RegisterNameserver"`
}

type UpdateNameServerResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	UpdateNameServer NameServerResult `xml:"RegisterNameserver"`
}

type DeleteNameServerResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	DeleteNameServer NameServerResult `xml:"RegisterNameserver"`
}

type NameServerResult struct {
	NS             string   `xml:"NS"`
	IP             string   `xml:"IP"`
	RegLock        string   `xml:"reg-lock"`
	RegistrarLocks []string `xml:"RegistrarLock"`
	NsSuccess      string   `xml:"NsSuccess"`
}

type SetDNSHostResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	IP string `xml:"IP"`
}

type AddDnsSecResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ResponseCode    int    `xml:"ResponseCode"`
	ResponseMessage string `xml:"ResponseMessage"`
}

type DeleteDnsSecResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ResponseCode    int    `xml:"ResponseCode"`
	ResponseMessage string `xml:"ResponseMessage"`
}

type GetDnsSecResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ResponseCode    int         `xml:"ResponseCode"`
	ResponseMessage string      `xml:"ResponseMessage"`
	DnsSecKeys      []DnsSecKey `xml:"DNSSEC"`
	DnsSecKeysAlt   []DnsSecKey `xml:"DnsSec"`
	DnsSecKeysWrap  []DnsSecKey `xml:"DNSSECs>DNSSEC"`
	Algorithms      []int       `xml:"Algorithm"`
	Digests         []string    `xml:"Digest"`
	DigestTypes     []int       `xml:"DigestType"`
	KeyTags         []int       `xml:"KeyTag"`
}

type DnsSecKey struct {
	Algorithm  int    `xml:"Algorithm"`
	Digest     string `xml:"Digest"`
	DigestType int    `xml:"DigestType"`
	KeyTag     int    `xml:"KeyTag"`
}

func (r *CheckNSStatusResponse) Decode() *response.CheckNSStatus {
	return &response.CheckNSStatus{
		CheckSuccess: r.NsCheckSuccess,
		NameServer: response.CheckNSStatusNameServer{
			Name:         r.CheckNsStatus.Name,
			AttribID:     r.CheckNsStatus.AttribID,
			IPAddress:    r.CheckNsStatus.IPAddress,
			AttribUPID:   r.CheckNsStatus.AttribUPID,
			AttribCLID:   r.CheckNsStatus.AttribCLID,
			AttribCRID:   r.CheckNsStatus.AttribCRID,
			Status:       append([]string(nil), r.CheckNsStatus.Status...),
			AttribUpdate: r.CheckNsStatus.AttribUpdate,
			AttribCRDate: r.CheckNsStatus.AttribCRDate,
		},
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDNSResponse) Decode() *response.GetDNS {
	registrySynced := r.DNSRegistrySynced
	if registrySynced == "" {
		registrySynced = r.NameserverRegistrySynced
	}
	return &response.GetDNS{
		DNS:                      append([]string(nil), r.DNS...),
		UseDNS:                   r.UseDNS,
		NSStatus:                 r.NSStatus,
		HostsNumLimit:            r.HostsNumLimit,
		DNSRegistrySynced:        registrySynced,
		NameserverRegistrySynced: r.NameserverRegistrySynced,
		RRPCodeGDNS:              r.RRPCodeGDNS,
		ResponseMeta:             decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDNSStatusResponse) Decode() *response.GetDNSStatus {
	return &response.GetDNSStatus{
		UseDNS:        r.UseDNS,
		NSStatus:      r.NSStatus,
		HostsNumLimit: r.HostsNumLimit,
		ResponseMeta:  decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetHomeDomainListResponse) Decode() *response.GetHomeDomainList {
	return &response.GetHomeDomainList{
		StartPosition:    r.GetHomeDomains.StartPosition,
		Display:          r.GetHomeDomains.Display,
		OrderBy:          r.GetHomeDomains.OrderBy,
		NewStartPosition: r.GetHomeDomains.NewStartPosition,
		DomainCount:      r.GetHomeDomains.DomainCount,
		TotalDomains:     r.GetHomeDomains.TotalDomains,
		Domains:          decodeHomeDomains(r.GetHomeDomains.Domains),
		ResponseMeta:     decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *RegisterNameServerResponse) Decode() *response.RegisterNameServer {
	return decodeNameServerResult(r.RegisterNameServer, decodeResponseMeta(r.Response, r.ResponseMeta))
}

func (r *UpdateNameServerResponse) Decode() *response.UpdateNameServer {
	return &response.UpdateNameServer{
		NSSuccess:    r.UpdateNameServer.NsSuccess,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *DeleteNameServerResponse) Decode() *response.DeleteNameServer {
	return &response.DeleteNameServer{
		NSSuccess:      r.DeleteNameServer.NsSuccess,
		RegistrarLocks: append([]string(nil), r.DeleteNameServer.RegistrarLocks...),
		ResponseMeta:   decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *SetDNSHostResponse) Decode() *response.SetDNSHost {
	return &response.SetDNSHost{
		IP:           r.IP,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *AddDnsSecResponse) Decode() *response.DnsSecUpdate {
	return &response.DnsSecUpdate{
		ResponseCode:    r.ResponseCode,
		ResponseMessage: r.ResponseMessage,
		ResponseMeta:    decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *DeleteDnsSecResponse) Decode() *response.DnsSecUpdate {
	return &response.DnsSecUpdate{
		ResponseCode:    r.ResponseCode,
		ResponseMessage: r.ResponseMessage,
		ResponseMeta:    decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetDnsSecResponse) Decode() *response.GetDnsSec {
	keys := decodeDnsSecKeys(r)
	return &response.GetDnsSec{
		ResponseCode:    r.ResponseCode,
		ResponseMessage: r.ResponseMessage,
		Keys:            keys,
		ResponseMeta:    decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func decodeNameServerResult(result NameServerResult, meta response.ResponseMeta) *response.RegisterNameServer {
	return &response.RegisterNameServer{
		NameServer:     result.NS,
		IPAddress:      result.IP,
		RegLock:        result.RegLock,
		RegistrarLocks: append([]string(nil), result.RegistrarLocks...),
		NSSuccess:      result.NsSuccess,
		ResponseMeta:   meta,
	}
}

func decodeHomeDomains(domains []HomeDomain) []response.HomeDomain {
	if len(domains) == 0 {
		return nil
	}
	result := make([]response.HomeDomain, 0, len(domains))
	for _, domain := range domains {
		result = append(result, response.HomeDomain{
			ID:        domain.ID,
			Name:      domain.Name,
			RegStatus: domain.RegStatus,
		})
	}
	return result
}

func decodeDnsSecKeys(r *GetDnsSecResponse) []response.DnsSecKey {
	keys := r.DnsSecKeys
	if len(keys) == 0 {
		keys = r.DnsSecKeysAlt
	}
	if len(keys) == 0 {
		keys = r.DnsSecKeysWrap
	}
	if len(keys) == 0 {
		return decodeDnsSecKeysFromFields(r)
	}
	result := make([]response.DnsSecKey, 0, len(keys))
	for _, key := range keys {
		result = append(result, response.DnsSecKey{
			Algorithm:  key.Algorithm,
			Digest:     key.Digest,
			DigestType: key.DigestType,
			KeyTag:     key.KeyTag,
		})
	}
	return result
}

func decodeDnsSecKeysFromFields(r *GetDnsSecResponse) []response.DnsSecKey {
	count := len(r.Digests)
	if len(r.Algorithms) < count {
		count = len(r.Algorithms)
	}
	if len(r.DigestTypes) < count {
		count = len(r.DigestTypes)
	}
	if len(r.KeyTags) < count {
		count = len(r.KeyTags)
	}
	if count == 0 {
		return nil
	}
	keys := make([]response.DnsSecKey, 0, count)
	for i := 0; i < count; i++ {
		keys = append(keys, response.DnsSecKey{
			Algorithm:  r.Algorithms[i],
			Digest:     r.Digests[i],
			DigestType: r.DigestTypes[i],
			KeyTag:     r.KeyTags[i],
		})
	}
	return keys
}
