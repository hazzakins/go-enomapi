package enomapi

import "strings"

func NewDomain(fqdn string) Domain {
	result := Domain{}
	result.Name, result.Extension, _ = strings.Cut(fqdn, ".")
	return result
}
