package domains

import (
	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

type Client struct {
	*enomapi.Client
}

// ParameterSet represents a collection of optional parameters to include with a
// command. Only entries with non-empty values are transmitted.
type ParameterSet map[string]string

func addParameters(cmd *enomapi.Command, params ParameterSet) {
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}
}

func executeGenericCommand(c *Client, command string, params ParameterSet) (*response.Generic, error) {
	resp := internal.GenericResponse{}

	cmd := c.NewCommand(command)
	addParameters(cmd, params)
	return resp.Decode(), c.Execute(cmd, &resp)
}
