package response

// CheckNSStatus reports registry status details for a name server.
type CheckNSStatus struct {
	CheckSuccess string
	NameServer   CheckNSStatusNameServer
	ResponseMeta ResponseMeta
}

// CheckNSStatusNameServer describes a name server returned from CheckNSStatus.
type CheckNSStatusNameServer struct {
	Name         string
	AttribID     string
	IPAddress    string
	AttribUPID   string
	AttribCLID   string
	AttribCRID   string
	Status       []string
	AttribUpdate string
	AttribCRDate string
}

// GetDNS returns the name servers configured for a domain.
type GetDNS struct {
	DNS                      []string
	UseDNS                   string
	NSStatus                 string
	HostsNumLimit            int
	DNSRegistrySynced        string
	NameserverRegistrySynced string
	RRPCodeGDNS              int
	ResponseMeta             ResponseMeta
}

// GetDNSStatus returns DNS status information for a domain.
type GetDNSStatus struct {
	UseDNS        string
	NSStatus      string
	HostsNumLimit int
	ResponseMeta  ResponseMeta
}

// GetHomeDomainList returns domains in the account using eNom DNS.
type GetHomeDomainList struct {
	StartPosition    int
	Display          int
	OrderBy          string
	NewStartPosition int
	DomainCount      int
	TotalDomains     int
	Domains          []HomeDomain
	ResponseMeta     ResponseMeta
}

// HomeDomain describes a domain returned by GetHomeDomainList.
type HomeDomain struct {
	ID        int64
	Name      string
	RegStatus string
}

// RegisterNameServer represents a RegisterNameServer response.
type RegisterNameServer struct {
	NameServer     string
	IPAddress      string
	RegLock        string
	RegistrarLocks []string
	NSSuccess      string
	ResponseMeta   ResponseMeta
}

// UpdateNameServer represents an UpdateNameServer response.
type UpdateNameServer struct {
	NSSuccess    string
	ResponseMeta ResponseMeta
}

// DeleteNameServer represents a DeleteNameServer response.
type DeleteNameServer struct {
	NSSuccess      string
	RegistrarLocks []string
	ResponseMeta   ResponseMeta
}

// SetDNSHost represents a SetDNSHost response.
type SetDNSHost struct {
	IP           string
	ResponseMeta ResponseMeta
}

// DnsSecUpdate represents the response for DNSSEC update commands.
type DnsSecUpdate struct {
	ResponseCode    int
	ResponseMessage string
	ResponseMeta    ResponseMeta
}

// GetDnsSec represents the response for GetDnsSec.
type GetDnsSec struct {
	ResponseCode    int
	ResponseMessage string
	Keys            []DnsSecKey
	ResponseMeta    ResponseMeta
}

// DnsSecKey describes a DNSSEC key entry.
type DnsSecKey struct {
	Algorithm  int
	Digest     string
	DigestType int
	KeyTag     int
}
