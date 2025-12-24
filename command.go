package enomapi

import (
	"net/url"
)

// Command represents an ENOM API command with the parameters to execute it.
type Command struct {
	Name   string
	Params url.Values
}

// NewCommand creates a Command pre-populated with the client's default
// parameters for authentication and response handling.
func (c *Client) NewCommand(commandName string) *Command {
	cmd := &Command{
		Name: commandName}
	cmd.setDefaultParams(c)
	return cmd
}

func (c *Command) toParams() string {
	return c.Params.Encode()
}

func (c *Command) setDefaultParams(client *Client) {
	c.Params = url.Values{}
	c.Params.Set("Command", c.Name)
	c.Params.Set("UID", client.resellerID)
	c.Params.Set("PW", client.apikey)
	c.Params.Set("ResponseType", "XML")
}

// AddParam sets or replaces a parameter on the Command.
func (c *Command) AddParam(key string, value string) {
	c.Params.Set(key, value)
}
