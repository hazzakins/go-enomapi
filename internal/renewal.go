package internal

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/hazzakins/go-enomapi/response"
)

type ExtendResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Extension     string           `xml:"Extension"`
	DomainName    string           `xml:"DomainName"`
	OrderID       string           `xml:"OrderID"`
	UseCreditCard bool             `xml:"UseCreditCard"`
	TotalCharged  float64          `xml:"TotalCharged"`
	DomainInfo    ExtendDomainInfo `xml:"DomainInfo"`
}

type ExtendDomainInfo struct {
	RegistryExpDate string `xml:"RegistryExpDate"`
}

func (r *ExtendResponse) Decode() *response.Extend {
	return &response.Extend{
		Extension:       r.Extension,
		DomainName:      r.DomainName,
		OrderID:         r.OrderID,
		RegistryExpDate: r.DomainInfo.RegistryExpDate,
		UseCreditCard:   r.UseCreditCard,
		TotalCharged:    r.TotalCharged,
		ResponseMeta:    decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

type ExtendRGPResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	Extension  string `xml:"Extension"`
	DomainName string `xml:"DomainName"`
	OrderID    string `xml:"OrderID"`
}

func (r *ExtendRGPResponse) Decode() *response.ExtendRGP {
	return &response.ExtendRGP{
		Extension:    r.Extension,
		DomainName:   r.DomainName,
		OrderID:      r.OrderID,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

type InsertNewOrderResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	OrderID              string                `xml:"OrderID"`
	PremiumDomainOrderID string                `xml:"PremiumDomainOrderID"`
	Fields               []InsertNewOrderField `xml:",any"`
}

type InsertNewOrderField struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (r *InsertNewOrderResponse) Decode() *response.InsertNewOrder {
	return &response.InsertNewOrder{
		OrderID:              r.OrderID,
		PremiumDomainOrderID: r.PremiumDomainOrderID,
		ProductTypes:         decodeInsertNewOrderProductTypes(r.Fields),
		ResponseMeta:         decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

func decodeInsertNewOrderProductTypes(fields []InsertNewOrderField) []response.InsertNewOrderProductType {
	if len(fields) == 0 {
		return nil
	}
	items := make([]response.InsertNewOrderProductType, 0, len(fields))
	for _, field := range fields {
		name := field.XMLName.Local
		if !strings.HasPrefix(name, "ProdType") {
			continue
		}
		suffix := strings.TrimPrefix(name, "ProdType")
		index := 0
		if suffix != "" {
			if parsed, err := strconv.Atoi(suffix); err == nil {
				index = parsed
			}
		}
		items = append(items, response.InsertNewOrderProductType{
			Index:       index,
			ProductType: strings.TrimSpace(field.Value),
		})
	}
	if len(items) == 0 {
		return nil
	}
	return items
}
