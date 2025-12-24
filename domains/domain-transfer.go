package domains

import (
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

// Upstream documentation: https://api.enom.com/docs/pe-get-tld-id
func (c *Client) PEGetTLDID(tld string) (*response.PEGetTLDID, error) {
	resp := internal.PEGetTLDIDResponse{}

	cmd := c.NewCommand("PE_GetTLDID")
	cmd.AddParam("TLD", tld)

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/pe-set-pricing
func (c *Client) PESetPricing(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("PE_SetPricing")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/pushdomain
func (c *Client) PushDomain(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("PushDomain")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/refill-account
func (c *Client) RefillAccount(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("RefillAccount")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/setresellerservicespricing
func (c *Client) SetResellerServicesPricing(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetResellerServicesPricing")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/setresellertldpricing
func (c *Client) SetResellerTLDPricing(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SetResellerTLDPricing")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/synchauthinfo
func (c *Client) SynchAuthInfo(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("SynchAuthInfo")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-cancelorder
func (c *Client) TPCancelOrder(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_CancelOrder")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-createorder
func (c *Client) TPCreateOrder(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_CreateOrder")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getdetailsbydomain
func (c *Client) TPGetDetailsByDomain(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetDetailsByDomain")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorder
func (c *Client) TPGetOrder(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrder")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderdetail
func (c *Client) TPGetOrderDetail(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderDetail")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getordersbydomain
func (c *Client) TPGetOrdersByDomain(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrdersByDomain")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderreview
func (c *Client) TPGetOrderReview(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderReview")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-getorderstatuses
func (c *Client) TPGetOrderStatuses(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetOrderStatuses")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-get-tld-info
func (c *Client) TPGetTLDInfo(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_GetTLDInfo")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-resendemail
func (c *Client) TPResendEmail(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_ResendEmail")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-resubmit-locked
func (c *Client) TPResubmitLocked(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_ResubmitLocked")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-submitorder
func (c *Client) TPSubmitOrder(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_SubmitOrder")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/tp-updateorderdetail
func (c *Client) TPUpdateOrderDetail(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("TP_UpdateOrderDetail")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/updateaccountpricing
func (c *Client) UpdateAccountPricing(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("UpdateAccountPricing")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}

// Upstream documentation: https://api.enom.com/docs/updatepushlist
func (c *Client) UpdatePushList(params map[string]string) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand("UpdatePushList")
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}

	err := c.Execute(cmd, &resp)
	return resp.Decode(), err
}
