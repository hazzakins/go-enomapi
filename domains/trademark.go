package domains

import (
	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type TMGetNoticeRequest struct {
	Domain    enomapi.Domain
	LookupKey string
}

type TMUpdateCartRequest struct {
	Domain        enomapi.Domain
	TcnID         string
	TcnExpDate    string
	TcnAcceptDate string
}

// Upstream documentation: https://api.enom.com/docs/tm-check
func (c *Client) TMCheck(domain enomapi.Domain) (*response.TMCheck, error) {
	resp := internal.TMCheckResponse{}

	cmd := c.NewCommand("TM_Check")
	cmd.AddParam("SLD", domain.Name)
	cmd.AddParam("TLD", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tm-getnotice
func (c *Client) TMGetNotice(req TMGetNoticeRequest) (*response.TMGetNotice, error) {
	resp := internal.TMGetNoticeResponse{}

	cmd := c.NewCommand("TM_GetNotice")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)
	cmd.AddParam("LookupKey", req.LookupKey)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tm-updatecart
func (c *Client) TMUpdateCart(req TMUpdateCartRequest) (*response.TMUpdateCart, error) {
	resp := internal.TMUpdateCartResponse{}

	cmd := c.NewCommand("TM_UpdateCart")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("tcnID", req.TcnID)
	cmd.AddParam("tcnExpDate", req.TcnExpDate)
	cmd.AddParam("tcnAcceptDate", req.TcnAcceptDate)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
