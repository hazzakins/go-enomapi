package internal

import (
	"github.com/hazzakins/go-enomapi/response"
)

type SpinnerDomain struct {
	Name     string `xml:"name,attr"`
	Com      string `xml:"com,attr"`
	ComScore string `xml:"comscore,attr"`
	Net      string `xml:"net,attr"`
	NetScore string `xml:"netscore,attr"`
	Tv       string `xml:"tv,attr"`
	TvScore  string `xml:"tvscore,attr"`
	Cc       string `xml:"cc,attr"`
	CcScore  string `xml:"ccscore,attr"`
}

type NameSpinner struct {
	SpinCount int `xml:"spincount"`
	TLDList   string
	Domains   []SpinnerDomain `xml:"domains>domain"`
}

type SpinnerResponse struct {
	Response
	NameSpinner NameSpinner `xml:"namespin"`
}

func (s *SpinnerResponse) Decode() *response.NameSpinner {
	return &response.NameSpinner{}
}
