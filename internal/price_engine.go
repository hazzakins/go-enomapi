package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

// PEGetTLDIDResponse is the raw XML shape for the price engine TLD ID lookup.
type PEGetTLDIDResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ProductID PEProductID `xml:"productid"`
}

// PEProductID contains the TLD identifier returned by the price engine.
type PEProductID struct {
	TLDID string `xml:"tldid"`
}

// Decode converts the raw PEGetTLDIDResponse into the public response type.
func (r *PEGetTLDIDResponse) Decode() *response.PEGetTLDID {
	return &response.PEGetTLDID{
		TLDID:        r.ProductID.TLDID,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}

// PEGetDomainPricingResponse is the raw XML shape for domain pricing lookups.
type PEGetDomainPricingResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	PriceStructure PEDomainPricingStructure `xml:"pricestructure"`
}

// PEDomainPricingStructure holds the pricing structure list.
type PEDomainPricingStructure struct {
	DisplayPriceIncrease int                      `xml:"DisplayPriceIncrease"`
	Products             []PEDomainPricingProduct `xml:"product"`
	Count                int                      `xml:"count"`
}

// PEDomainPricingProduct captures pricing details for a single TLD.
type PEDomainPricingProduct struct {
	TLD                string  `xml:"tld"`
	TLDID              int     `xml:"tldid"`
	RegisterPrice      float64 `xml:"registerprice"`
	ResellerPriceReg   float64 `xml:"resellerpricereg"`
	RegisterEnabled    bool    `xml:"registerenabled"`
	RenewPrice         float64 `xml:"renewprice"`
	ResellerPriceRenew float64 `xml:"resellerpricerenew"`
	RenewEnabled       bool    `xml:"renewenabled"`
	TransferPrice      float64 `xml:"transferprice"`
	ResellerPriceTran  float64 `xml:"resellerpricetran"`
	TransferEnabled    bool    `xml:"transferenabled"`
	RGPPrice           float64 `xml:"rgpprice"`
	ResellerPriceRGP   float64 `xml:"resellerpricergp"`
	RGPEnabled         bool    `xml:"rgpenabled"`
}

// Decode converts the raw PEGetDomainPricingResponse into the public response type.
func (r *PEGetDomainPricingResponse) Decode() *response.PEGetDomainPricing {
	products := make([]response.PEDomainPricingProduct, 0, len(r.PriceStructure.Products))
	for _, product := range r.PriceStructure.Products {
		products = append(products, response.PEDomainPricingProduct{
			TLD:                product.TLD,
			TLDID:              product.TLDID,
			RegisterPrice:      product.RegisterPrice,
			ResellerPriceReg:   product.ResellerPriceReg,
			RegisterEnabled:    product.RegisterEnabled,
			RenewPrice:         product.RenewPrice,
			ResellerPriceRenew: product.ResellerPriceRenew,
			RenewEnabled:       product.RenewEnabled,
			TransferPrice:      product.TransferPrice,
			ResellerPriceTran:  product.ResellerPriceTran,
			TransferEnabled:    product.TransferEnabled,
			RGPPrice:           product.RGPPrice,
			ResellerPriceRGP:   product.ResellerPriceRGP,
			RGPEnabled:         product.RGPEnabled,
		})
	}
	return &response.PEGetDomainPricing{
		DisplayPriceIncrease: r.PriceStructure.DisplayPriceIncrease,
		Products:             products,
		Count:                r.PriceStructure.Count,
		ResponseMeta:         decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}
