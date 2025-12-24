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
