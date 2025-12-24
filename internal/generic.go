package internal

import (
	"encoding/xml"

	"github.com/hazzakins/go-enomapi/response"
)

// GenericResponse is a minimal response wrapper for commands that do not expose
// additional structured fields beyond the standard metadata.
type GenericResponse struct {
	XMLName xml.Name `xml:"interface-response"`
	Response
	ResponseMeta
}

// Decode converts the generic response into its public counterpart.
func (r *GenericResponse) Decode() *response.Generic {
	return &response.Generic{
		ResponseMeta: decodeResponseMeta(r.Response, r.ResponseMeta),
	}
}
