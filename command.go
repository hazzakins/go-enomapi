package enomapi

import (
	"net/url"
)

type Command struct {
	Name   string
	Params url.Values
}

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

func (c *Command) AddParam(key string, value string) {
	c.Params.Set(key, value)
}
