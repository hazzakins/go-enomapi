package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

type PEGetTLDIDResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
	ProductID PEProductID `xml:"productid"`
}

type PEProductID struct {
	TLDID string `xml:"tldid"`
}

func (r *PEGetTLDIDResponse) Decode() *response.PEGetTLDID {
	return &response.PEGetTLDID{
		TLDID:        r.ProductID.TLDID,
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}
