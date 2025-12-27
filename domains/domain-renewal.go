package domains

import (
	"strconv"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type ExtendRequest struct {
	Domain                        enomapi.Domain
	NumYears                      int
	OverrideOrder                 *bool
	UseCreditCard                 *bool
	EndUserIP                     string
	ChargeAmount                  string
	CustomerSuppliedPrice         string
	CardType                      string
	CCName                        string
	CreditCardNumber              string
	CreditCardExpMonth            string
	CreditCardExpYear             string
	CVV2                          string
	CCAddress                     string
	CCZip                         string
	CCCountry                     string
	RegistrantFirstName           string
	RegistrantLastName            string
	RegistrantAddress1            string
	RegistrantAddress2            string
	RegistrantCity                string
	RegistrantStateProvinceChoice string
	RegistrantStateProvince       string
	RegistrantPostalCode          string
	RegistrantCountry             string
	RegistrantEmailAddress        string
	RegistrantOrganizationName    string
	RegistrantJobTitle            string
	RegistrantPhone               string
	RegistrantFax                 string
}

// Upstream documentation: https://api.enom.com/docs/extend
func (c *Client) Extend(req ExtendRequest) (*response.Extend, error) {
	resp := internal.ExtendResponse{}

	cmd := c.NewCommand("Extend")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)
	cmd.AddParam("NumYears", strconv.Itoa(req.NumYears))
	addBoolIntParam(cmd, "OverrideOrder", req.OverrideOrder)
	addBoolParam(cmd, "UseCreditCard", req.UseCreditCard)
	if req.EndUserIP != "" {
		cmd.AddParam("EndUserIP", req.EndUserIP)
	}
	if req.ChargeAmount != "" {
		cmd.AddParam("ChargeAmount", req.ChargeAmount)
	}
	if req.CustomerSuppliedPrice != "" {
		cmd.AddParam("CustomerSuppliedPrice", req.CustomerSuppliedPrice)
	}
	if req.CardType != "" {
		cmd.AddParam("CardType", req.CardType)
	}
	if req.CCName != "" {
		cmd.AddParam("CCName", req.CCName)
	}
	if req.CreditCardNumber != "" {
		cmd.AddParam("CreditCardNumber", req.CreditCardNumber)
	}
	if req.CreditCardExpMonth != "" {
		cmd.AddParam("CreditCardExpMonth", req.CreditCardExpMonth)
	}
	if req.CreditCardExpYear != "" {
		cmd.AddParam("CreditCardExpYear", req.CreditCardExpYear)
	}
	if req.CVV2 != "" {
		cmd.AddParam("CVV2", req.CVV2)
	}
	if req.CCAddress != "" {
		cmd.AddParam("CCAddress", req.CCAddress)
	}
	if req.CCZip != "" {
		cmd.AddParam("CCZip", req.CCZip)
	}
	if req.CCCountry != "" {
		cmd.AddParam("CCCountry", req.CCCountry)
	}
	if req.RegistrantFirstName != "" {
		cmd.AddParam("RegistrantFirstName", req.RegistrantFirstName)
	}
	if req.RegistrantLastName != "" {
		cmd.AddParam("RegistrantLastName", req.RegistrantLastName)
	}
	if req.RegistrantAddress1 != "" {
		cmd.AddParam("RegistrantAddress1", req.RegistrantAddress1)
	}
	if req.RegistrantAddress2 != "" {
		cmd.AddParam("RegistrantAddress2", req.RegistrantAddress2)
	}
	if req.RegistrantCity != "" {
		cmd.AddParam("RegistrantCity", req.RegistrantCity)
	}
	if req.RegistrantStateProvinceChoice != "" {
		cmd.AddParam("RegistrantStateProvinceChoice", req.RegistrantStateProvinceChoice)
	}
	if req.RegistrantStateProvince != "" {
		cmd.AddParam("RegistrantStateProvince", req.RegistrantStateProvince)
	}
	if req.RegistrantPostalCode != "" {
		cmd.AddParam("RegistrantPostalCode", req.RegistrantPostalCode)
	}
	if req.RegistrantCountry != "" {
		cmd.AddParam("RegistrantCountry", req.RegistrantCountry)
	}
	if req.RegistrantEmailAddress != "" {
		cmd.AddParam("RegistrantEmailAddress", req.RegistrantEmailAddress)
	}
	if req.RegistrantOrganizationName != "" {
		cmd.AddParam("RegistrantOrganizationName", req.RegistrantOrganizationName)
	}
	if req.RegistrantJobTitle != "" {
		cmd.AddParam("RegistrantJobTitle", req.RegistrantJobTitle)
	}
	if req.RegistrantPhone != "" {
		cmd.AddParam("RegistrantPhone", req.RegistrantPhone)
	}
	if req.RegistrantFax != "" {
		cmd.AddParam("RegistrantFax", req.RegistrantFax)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

type ExtendRGPRequest struct {
	Domain enomapi.Domain
}

type InsertNewOrderRequest struct {
	EndUserIP       string
	ItemCount       *int
	UseWireTransfer *bool
}

type SetRenewRequest struct {
	Domain            enomapi.Domain
	RenewFlag         bool
	AutoPakRenew      *bool
	EmailForwardRenew *bool
	URLForwardRenew   *bool
	WPPSRenew         *bool
}

// Upstream documentation: https://api.enom.com/docs/extend-rgp
func (c *Client) ExtendRGP(req ExtendRGPRequest) (*response.ExtendRGP, error) {
	resp := internal.ExtendRGPResponse{}

	cmd := c.NewCommand("Extend_RGP")
	cmd.AddParam("SLD", req.Domain.Name)
	cmd.AddParam("TLD", req.Domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/get-domain-exp
func (c *Client) GetDomainExp(domain enomapi.Domain) (*response.GetDomainExp, error) {
	resp := internal.GetDomainExpResponse{}

	cmd := c.NewCommand("GetDomainExp")
	cmd.AddParam("SLD", domain.Name)
	cmd.AddParam("TLD", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetExtendInfo returns renewal information for a domain.
func (c *Client) GetExtendInfo(domain enomapi.Domain) (*response.GetExtendInfo, error) {
	resp := internal.GetExtendInfoResponse{}

	cmd := c.NewCommand("GetExtendInfo")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetRenew returns auto-renew settings for a domain.
func (c *Client) GetRenew(domain enomapi.Domain) (*response.GetRenew, error) {
	resp := internal.GetRenewResponse{}

	cmd := c.NewCommand("GetRenew")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/insertneworder
func (c *Client) InsertNewOrder(req InsertNewOrderRequest) (*response.InsertNewOrder, error) {
	resp := internal.InsertNewOrderResponse{}

	cmd := c.NewCommand("InsertNewOrder")
	cmd.AddParam("EndUserIP", req.EndUserIP)
	addIntParam(cmd, "ItemCount", req.ItemCount)
	if req.UseWireTransfer != nil {
		if *req.UseWireTransfer {
			cmd.AddParam("UseWireTransfer", "Yes")
		} else {
			cmd.AddParam("UseWireTransfer", "No")
		}
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// SetRenew updates the auto-renew flag and related settings for a domain.
func (c *Client) SetRenew(req SetRenewRequest) (*response.SetRenew, error) {
	resp := internal.SetRenewResponse{}

	cmd := c.NewCommand("SetRenew")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)
	cmd.AddParam("RenewFlag", boolToInt(req.RenewFlag))
	addBoolIntParam(cmd, "AutoPakRenew", req.AutoPakRenew)
	addBoolIntParam(cmd, "EmailForwardRenew", req.EmailForwardRenew)
	addBoolIntParam(cmd, "URLForwardRenew", req.URLForwardRenew)
	addBoolIntParam(cmd, "WPPSRenew", req.WPPSRenew)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
// TODO: UpdateExpiredDomains
// TODO: UpdateRenewalSettings
