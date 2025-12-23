package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

type TLDListResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	TLDList  []string `xml:"tldlist>tld>tld"`
	TLDCount int32    `xml:"tldlist>tldcount"`
}

func (t *TLDListResponse) Decode() response.TLDList {
	result := make(response.TLDList, t.TLDCount)
	copy(result, t.TLDList)
	return result
}
