package response

// GetDomainSRVHosts returns SRV host records for a domain.
type GetDomainSRVHosts struct {
	SRVRecords   []SRVHostRecord
	ResponseMeta ResponseMeta
}

// SRVHostRecord describes an SRV host record.
type SRVHostRecord struct {
	HostID     int64
	HostName   string
	Protocol   string
	Address    string
	RecordType string
	MXPref     int
	Weight     int
	Priority   int
	Port       int
}

// GetHosts returns host records for a domain.
type GetHosts struct {
	HostRecords  []HostRecord
	HostCount    int
	ResponseMeta ResponseMeta
}

// HostRecord describes a host record returned by GetHosts.
type HostRecord struct {
	HostID     int64
	HostName   string
	Address    string
	RecordType string
	MXPref     int
}

// GetRegHosts returns non-mail host records for a domain.
type GetRegHosts struct {
	HostRecords  []RegHostRecord
	HostCount    int
	ResponseMeta ResponseMeta
}

// RegHostRecord describes a host record returned by GetRegHosts.
type RegHostRecord struct {
	HostID     int64
	HostName   string
	Address    string
	RecordType string
}

// UpdateMetaTag returns metatag settings for a host record.
type UpdateMetaTag struct {
	TitleBar        string
	SiteDescription string
	Keywords        string
	DomainNameID    int64
	ResponseMeta    ResponseMeta
}

// GetSPFHosts returns SPF host records for a domain.
type GetSPFHosts struct {
	Hosts        []SPFHostRecord
	HostCount    int
	ResponseMeta ResponseMeta
}

// SPFHostRecord describes an SPF host record.
type SPFHostRecord struct {
	Name       string
	Type       string
	ARadio     string
	MXRadio    string
	PTRRadio   string
	ALLRadio   string
	ARecords   string
	MXRecords  string
	IPRecords  string
	INCRecords string
	HostID     int64
}
