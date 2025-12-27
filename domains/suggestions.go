package domains

import (
	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type GetNameSuggestionsRequest struct {
	SearchTerm     string
	TldList        string
	OnlyTldList    string
	ExcludeTldList string
	MaxResult      *int
	SpinType       *int
	Adult          *bool
	Premium        *bool
	AllGA          *bool
}

// Upstream documentation: https://api.enom.com/docs/getnamesuggestions
func (c *Client) GetNameSuggestions(req GetNameSuggestionsRequest) (*response.GetNameSuggestions, error) {
	resp := internal.GetNameSuggestionsResponse{}

	cmd := c.NewCommand("GetNameSuggestions")
	cmd.AddParam("SearchTerm", req.SearchTerm)
	if req.TldList != "" {
		cmd.AddParam("TldList", req.TldList)
	}
	if req.OnlyTldList != "" {
		cmd.AddParam("OnlyTldList", req.OnlyTldList)
	}
	if req.ExcludeTldList != "" {
		cmd.AddParam("ExcludeTldList", req.ExcludeTldList)
	}
	addIntParam(cmd, "MaxResult", req.MaxResult)
	addIntParam(cmd, "SpinType", req.SpinType)
	addBoolParam(cmd, "Adult", req.Adult)
	addBoolParam(cmd, "Premium", req.Premium)
	addBoolParam(cmd, "AllGA", req.AllGA)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/namespinner
func (c Client) NameSpinner(domain enomapi.Domain, tldList string) (*response.NameSpinner, error) {
	resp := internal.SpinnerResponse{}

	cmd := c.NewCommand("NameSpinner")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("tldlist", tldList)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
