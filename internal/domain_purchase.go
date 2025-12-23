package internal

import (
	"fmt"

	"github.com/hazzakins/go-enomapi/response"
)

type DomainPurchase struct {
	Response
	OrderID           string             `xml:"OrderID"`
	OrderDelayed      bool               `xml:"OrderDelayed"`
	OrderStatus       string             `xml:"OrderStatus"`
	OrderDescription  string             `xml:"OrderDescription"`
	DomainInfo        DomainPurchaseInfo `xml:"Info"`
	TotalCharged      float64            `xml:"TotalCharged"`
	RegistrantPartyID string             `xml:"RegistrantPartyID"`
	IsRealTimeTLD     bool               `xml:"IsRealTimeTLD"`
}

type DomainPurchaseInfo struct {
	RegistryCreateDate string `xml:"RegistryCreateDate"`
	RegistryExpDate    string `xml:"RegistryExpDate"`
}

func (d *DomainPurchase) Decode() (*response.DomainPurchase, error) {
	if d.OrderID == "" {
		return nil, fmt.Errorf("no OrderID received - RRPCode: %d", d.ResponseCode)
	}
	result := response.DomainPurchase{
		OrderID:          d.OrderID,
		OrderCompleted:   d.DomainInfo.RegistryCreateDate != "",
		OrderStatus:      d.OrderStatus,
		OrderDescription: d.OrderDescription,
		Price:            d.TotalCharged,
	}
	var err error
	if d.DomainInfo.RegistryCreateDate != "" {
		result.RegistrationDate, err = ParseDate(d.DomainInfo.RegistryCreateDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RegistryCreateDate: %w", err)
		}
	}
	if d.DomainInfo.RegistryExpDate != "" {
		result.ExpirationDate, err = ParseDate(d.DomainInfo.RegistryExpDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RegistryExpDate: %w", err)
		}
	}
	return &result, nil
}
