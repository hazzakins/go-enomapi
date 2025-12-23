package domainmanagement

import (
	"strconv"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

// ContactType defines the contact type for the Contacts command.
type ContactType string

const (
	ContactTypeRegistrant ContactType = "REGISTRANT"
	ContactTypeAuxBilling ContactType = "AUXBILLING"
	ContactTypeTech       ContactType = "TECH"
	ContactTypeAdmin      ContactType = "ADMIN"
)

// ContactDetails represents contact data for Contacts updates.
type ContactDetails struct {
	FirstName           string
	LastName            string
	OrganizationName    string
	JobTitle            string
	Address1            string
	Address2            string
	City                string
	StateProvinceChoice string
	StateProvince       string
	PostalCode          string
	Country             string
	EmailAddress        string
	Phone               string
	PhoneExt            string
	Fax                 string
}

// ContactsRequest defines input for the Contacts command.
type ContactsRequest struct {
	Domain                enomapi.Domain
	ContactType           ContactType
	Registrant            *ContactDetails
	AuxBilling            *ContactDetails
	Tech                  *ContactDetails
	Admin                 *ContactDetails
	ExtendedAttributes    string
	IRTPOptOut            *bool
	IRTPOptOutReason      string
	IRTPEmailLanguageCode string
}

// Contacts updates contact information for a domain name.
func (c Client) Contacts(req ContactsRequest) (*response.Contacts, error) {
	resp := internal.ContactsResponse{}

	cmd := c.NewCommand("Contacts")
	cmd.AddParam("sld", req.Domain.Name)
	cmd.AddParam("tld", req.Domain.Extension)

	if req.ContactType != "" {
		cmd.AddParam("ContactType", string(req.ContactType))
	}

	if req.Registrant != nil {
		addContactParams(cmd, "Registrant", *req.Registrant)
	}
	if req.AuxBilling != nil {
		addContactParams(cmd, "AuxBilling", *req.AuxBilling)
	}
	if req.Tech != nil {
		addContactParams(cmd, "Tech", *req.Tech)
	}
	if req.Admin != nil {
		addContactParams(cmd, "Admin", *req.Admin)
	}

	if req.ExtendedAttributes != "" {
		cmd.AddParam("ExtendedAttributes", req.ExtendedAttributes)
	}
	if req.IRTPOptOut != nil {
		cmd.AddParam("IRTPOptOut", strconv.FormatBool(*req.IRTPOptOut))
	}
	if req.IRTPOptOutReason != "" {
		cmd.AddParam("IRTPOptOutReason", req.IRTPOptOutReason)
	}
	if req.IRTPEmailLanguageCode != "" {
		cmd.AddParam("IRTPEmailLanguageCode", req.IRTPEmailLanguageCode)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetContacts returns contact information for a domain in the account.
func (c Client) GetContacts(domain enomapi.Domain) (*response.GetContacts, error) {
	resp := internal.GetContactsResponse{}

	cmd := c.NewCommand("GetContacts")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// TODO: GetExtAttributes

// GetWhoisContact returns public Whois contact information for a domain name.
func (c Client) GetWhoisContact(domain enomapi.Domain) (*response.GetWhoisContact, error) {
	resp := internal.GetWhoisContactResponse{}

	cmd := c.NewCommand("GetWhoisContact")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// GetWPPSInfo retrieves ID Protect (Whois privacy protection) status.
func (c Client) GetWPPSInfo(domain enomapi.Domain) (*response.GetWPPSInfo, error) {
	resp := internal.GetWPPSInfoResponse{}

	cmd := c.NewCommand("GetWPPSInfo")
	cmd.AddParam("sld", domain.Name)
	cmd.AddParam("tld", domain.Extension)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

func addContactParams(cmd *enomapi.Command, prefix string, contact ContactDetails) {
	if contact.FirstName != "" {
		cmd.AddParam(prefix+"FirstName", contact.FirstName)
	}
	if contact.LastName != "" {
		cmd.AddParam(prefix+"LastName", contact.LastName)
	}
	if contact.OrganizationName != "" {
		cmd.AddParam(prefix+"OrganizationName", contact.OrganizationName)
	}
	if contact.JobTitle != "" {
		cmd.AddParam(prefix+"JobTitle", contact.JobTitle)
	}
	if contact.Address1 != "" {
		cmd.AddParam(prefix+"Address1", contact.Address1)
	}
	if contact.Address2 != "" {
		cmd.AddParam(prefix+"Address2", contact.Address2)
	}
	if contact.City != "" {
		cmd.AddParam(prefix+"City", contact.City)
	}
	if contact.StateProvinceChoice != "" {
		cmd.AddParam(prefix+"StateProvinceChoice", contact.StateProvinceChoice)
	}
	if contact.StateProvince != "" {
		cmd.AddParam(prefix+"StateProvince", contact.StateProvince)
	}
	if contact.PostalCode != "" {
		cmd.AddParam(prefix+"PostalCode", contact.PostalCode)
	}
	if contact.Country != "" {
		cmd.AddParam(prefix+"Country", contact.Country)
	}
	if contact.EmailAddress != "" {
		cmd.AddParam(prefix+"EmailAddress", contact.EmailAddress)
	}
	if contact.Phone != "" {
		cmd.AddParam(prefix+"Phone", contact.Phone)
	}
	if contact.PhoneExt != "" {
		cmd.AddParam(prefix+"PhoneExt", contact.PhoneExt)
	}
	if contact.Fax != "" {
		cmd.AddParam(prefix+"Fax", contact.Fax)
	}
}
