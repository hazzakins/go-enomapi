package domains

import (
	"fmt"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type Client struct {
	*enomapi.Client
}

// TODO: AddBulkDomains
// Upstream documentation: https://api.enom.com/docs/addbulkdomains
// TODO: CancelOrder
// Upstream documentation: https://api.enom.com/docs/cancel-order

func (c *Client) Check(domain enomapi.Domain) (*response.DomainCheck, error) {
	resp := internal.DomainCheckResponse{}

	cmd := c.NewCommand("Check")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)
	cmd.AddParam("Version", "2")
	cmd.AddParam("IncludePrice", "true")
	cmd.AddParam("IncludeProperties", "true")

	err := c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}

	if resp.RRPCode > 300 {
		return nil, fmt.Errorf("RRP error %d: %s", resp.RRPCode, resp.RRPText)
	}

	return resp.Decode()
}

// TODO: GetConfirmationSettings
// Upstream documentation: https://api.enom.com/docs/get-confirmation-settings
// TODO: GetExtAttributes
// Upstream documentation: https://api.enom.com/docs/get-ext-attributes
// TODO: GetIDNCodes
// Upstream documentation: https://api.enom.com/docs/get-idn-codes
// TODO: GetNameSuggestions
// Upstream documentation: https://api.enom.com/docs/getnamesuggestions

func (c *Client) GetTLDDetails(tld string) (*response.TLDDetails, error) {
	resp := internal.TLDDetailsResponse{}

	cmd := c.NewCommand("GetTLDDetails")
	cmd.AddParam("tld", tld)

	err := c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Decode()
}

func (c *Client) GetTLDList() (response.TLDList, error) {
	resp := internal.TLDListResponse{}

	cmd := c.NewCommand("GetTLDList")

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

// TODO: Preconfigure
// Upstream documentation: https://api.enom.com/docs/preconfigure

// Upstream documentation: https://api.enom.com/docs/purchase
func (c Client) Purchase(domain enomapi.Domain) (*response.DomainPurchase, error) {
	resp := internal.DomainPurchase{}

	cmd := c.NewCommand("purchase")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Decode()
}

// TODO: Queue_DomainPurchase
// Upstream documentation: https://api.enom.com/docs/queue-domain-purchase
// TODO: Queue_GetDomains
// Upstream documentation: https://api.enom.com/docs/Queue_GetDomains
// TODO: Queue_GetExtAttributes
// Upstream documentation: https://api.enom.com/docs/queue-get-ext-attributes
// TODO: Queue_GetOrderDetail
// Upstream documentation: https://api.enom.com/docs/queue-get-order-detail
// TODO: Queue_GetOrders
// Upstream documentation: https://api.enom.com/docs/queue-get-orders
// TODO: TM_Check
// Upstream documentation: https://api.enom.com/docs/tm-check
// TODO: TM_GetNotice
// Upstream documentation: https://api.enom.com/docs/tm-getnotice
// TODO: TM_UpdateCart
// Upstream documentation: https://api.enom.com/docs/tm-updatecart
// TODO: GetAgreementPage
// Upstream documentation: https://api.enom.com/docs/GetAgreementPage
// TODO: Queue_GetInfo
// Upstream documentation: https://api.enom.com/docs/Queue_GetInfo
// TODO: DeleteRegistration
// Upstream documentation: https://api.enom.com/docs/DeleteRegistration
