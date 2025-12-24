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
	OriginalSLD string      `xml:"originalsld"`
	ResponseMeta
}

func (s *SpinnerResponse) Decode() *response.NameSpinner {
	return &response.NameSpinner{
		SpinCount:    s.NameSpinner.SpinCount,
		TLDList:      s.NameSpinner.TLDList,
		OriginalSLD:  s.OriginalSLD,
		Domains:      decodeSpinnerDomains(s.NameSpinner.Domains),
		ResponseMeta: decodeResponseMeta(s.Response, s.ResponseMeta),
	}
}

func decodeSpinnerDomains(domains []SpinnerDomain) []response.SpinnerDomain {
	if len(domains) == 0 {
		return nil
	}
	result := make([]response.SpinnerDomain, 0, len(domains))
	for _, domain := range domains {
		result = append(result, response.SpinnerDomain{
			Name:     domain.Name,
			Com:      domain.Com,
			ComScore: domain.ComScore,
			Net:      domain.Net,
			NetScore: domain.NetScore,
			Tv:       domain.Tv,
			TvScore:  domain.TvScore,
			Cc:       domain.Cc,
			CcScore:  domain.CcScore,
		})
	}
	return result
}
