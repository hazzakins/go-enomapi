package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

type GetDomainSRVHostsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	SRVRecords []SRVRecord `xml:"srv-records>srv"`
}

type GetHostsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	HostCount   int          `xml:"HostCount"`
	HostRecords []HostRecord `xml:"host"`
}

type GetRegHostsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	HostCount   int             `xml:"HostCount"`
	HostRecords []RegHostRecord `xml:"hostrecords>host"`
}

type UpdateMetaTagResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	MetaTags MetaTag `xml:"metatags"`
}

type GetSPFHostsResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	HostCount int             `xml:"HostCount"`
	SPFHosts  []SPFHostRecord `xml:"host"`
}

type SRVRecord struct {
	HostID     int64  `xml:"HostID"`
	HostName   string `xml:"HostName"`
	Protocol   string `xml:"Protocol"`
	Address    string `xml:"Address"`
	RecordType string `xml:"RecordType"`
	MXPref     int    `xml:"mxPref"`
	Weight     int    `xml:"Weight"`
	Priority   int    `xml:"priority"`
	Port       int    `xml:"Port"`
}

type HostRecord struct {
	HostID     int64  `xml:"hostid"`
	HostName   string `xml:"name"`
	Address    string `xml:"address"`
	RecordType string `xml:"type"`
	MXPref     int    `xml:"mxpref"`
}

type RegHostRecord struct {
	HostID     int64  `xml:"hostID"`
	HostName   string `xml:"HostName"`
	Address    string `xml:"Address"`
	RecordType string `xml:"RecordType"`
}

type MetaTag struct {
	TitleBar        string `xml:"titlebar"`
	SiteDescription string `xml:"sitedescription"`
	Keywords        string `xml:"keywords"`
	DomainNameID    int64  `xml:"DomainNameID"`
}

type SPFHostRecord struct {
	Name       string `xml:"name"`
	Type       string `xml:"type"`
	ARadio     string `xml:"a-radio"`
	MXRadio    string `xml:"mx-radio"`
	PTRRadio   string `xml:"ptr-radio"`
	ALLRadio   string `xml:"all-radio"`
	ARecords   string `xml:"a-records"`
	MXRecords  string `xml:"mx-records"`
	IPRecords  string `xml:"ip-records"`
	INCRecords string `xml:"inc-records"`
	HostID     int64  `xml:"hostid"`
}

func (r *GetDomainSRVHostsResponse) Decode() *response.GetDomainSRVHosts {
	records := make([]response.SRVHostRecord, 0, len(r.SRVRecords))
	for _, record := range r.SRVRecords {
		records = append(records, response.SRVHostRecord{
			HostID:     record.HostID,
			HostName:   record.HostName,
			Protocol:   record.Protocol,
			Address:    record.Address,
			RecordType: record.RecordType,
			MXPref:     record.MXPref,
			Weight:     record.Weight,
			Priority:   record.Priority,
			Port:       record.Port,
		})
	}

	return &response.GetDomainSRVHosts{
		SRVRecords:   records,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetHostsResponse) Decode() *response.GetHosts {
	records := make([]response.HostRecord, 0, len(r.HostRecords))
	for _, record := range r.HostRecords {
		records = append(records, response.HostRecord{
			HostID:     record.HostID,
			HostName:   record.HostName,
			Address:    record.Address,
			RecordType: record.RecordType,
			MXPref:     record.MXPref,
		})
	}

	return &response.GetHosts{
		HostRecords:  records,
		HostCount:    r.HostCount,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetRegHostsResponse) Decode() *response.GetRegHosts {
	records := make([]response.RegHostRecord, 0, len(r.HostRecords))
	for _, record := range r.HostRecords {
		records = append(records, response.RegHostRecord{
			HostID:     record.HostID,
			HostName:   record.HostName,
			Address:    record.Address,
			RecordType: record.RecordType,
		})
	}

	return &response.GetRegHosts{
		HostRecords:  records,
		HostCount:    r.HostCount,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *UpdateMetaTagResponse) Decode() *response.UpdateMetaTag {
	return &response.UpdateMetaTag{
		TitleBar:        r.MetaTags.TitleBar,
		SiteDescription: r.MetaTags.SiteDescription,
		Keywords:        r.MetaTags.Keywords,
		DomainNameID:    r.MetaTags.DomainNameID,
		ResponseMeta:    decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func (r *GetSPFHostsResponse) Decode() *response.GetSPFHosts {
	hosts := make([]response.SPFHostRecord, 0, len(r.SPFHosts))
	for _, host := range r.SPFHosts {
		hosts = append(hosts, response.SPFHostRecord{
			Name:       host.Name,
			Type:       host.Type,
			ARadio:     host.ARadio,
			MXRadio:    host.MXRadio,
			PTRRadio:   host.PTRRadio,
			ALLRadio:   host.ALLRadio,
			ARecords:   host.ARecords,
			MXRecords:  host.MXRecords,
			IPRecords:  host.IPRecords,
			INCRecords: host.INCRecords,
			HostID:     host.HostID,
		})
	}

	return &response.GetSPFHosts{
		Hosts:        hosts,
		HostCount:    r.HostCount,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}
