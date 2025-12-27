package response

import "time"

type DomainInfo struct {
	Domain         string
	Name           string
	Extension      string
	IsRegistered   bool
	ExpirationDate *time.Time
}
