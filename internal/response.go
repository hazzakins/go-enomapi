package internal

import (
	"fmt"
	"strings"
)

// Response represents the base response fields returned by the eNom API.
type Response struct {
	ResponseCode int     `xml:"RRPCode"`
	ResponseText string  `xml:"RRPText"`
	Command      string  `xml:"Command"`
	Language     string  `xml:"Language"`
	ErrCount     int     `xml:"ErrCount"`
	Errors       []Error `xml:"errors"`
}

// Error represents an individual error entry in the response.
type Error struct {
	Error string `xml:",any"`
}

// ResponseInfo describes the meaning of a response code per documentation.
type ResponseInfo struct {
	Code      int
	Message   string
	Category  string
	Retryable bool
}

// responseCodeTable maps RRPCode values to documentation info.
var responseCodeTable = map[int]ResponseInfo{
	200: {Code: 200, Message: "Command completed successfully", Category: "Success", Retryable: false},
	210: {Code: 210, Message: "Domain name available", Category: "Success", Retryable: false},
	211: {Code: 211, Message: "Domain name not available", Category: "Success", Retryable: false},
	212: {Code: 212, Message: "Nameserver name available", Category: "Success", Retryable: false},
	213: {Code: 213, Message: "Nameserver name not available", Category: "Success", Retryable: false},
	220: {Code: 220, Message: "Command completed successfully. Server closing connection", Category: "Success", Retryable: false},
	420: {Code: 420, Message: "Command failed due to server error. Server closing connection.", Category: "Failure", Retryable: false},
	421: {Code: 421, Message: "Command failed due to server error. Client should try again.", Category: "Failure", Retryable: true},
	500: {Code: 500, Message: "Invalid command name", Category: "Failure", Retryable: false},
	501: {Code: 501, Message: "Invalid command option", Category: "Failure", Retryable: false},
	502: {Code: 502, Message: "Invalid entity value", Category: "Failure", Retryable: false},
	503: {Code: 503, Message: "Invalid attribute name", Category: "Failure", Retryable: false},
	504: {Code: 504, Message: "Missing required attribute", Category: "Failure", Retryable: false},
	505: {Code: 505, Message: "Invalid attribute value syntax", Category: "Failure", Retryable: false},
	506: {Code: 506, Message: "Invalid option value", Category: "Failure", Retryable: false},
	507: {Code: 507, Message: "Invalid command format", Category: "Failure", Retryable: false},
	508: {Code: 508, Message: "Missing required entity", Category: "Failure", Retryable: false},
	509: {Code: 509, Message: "Missing command option", Category: "Failure", Retryable: false},
	520: {Code: 520, Message: "Server closing connection. Client should try opening new connection;", Category: "Failure", Retryable: false},
	521: {Code: 521, Message: "Too many sessions open. Server closing connection.", Category: "Failure", Retryable: false},
	530: {Code: 530, Message: "Authentication failed", Category: "Failure", Retryable: false},
	531: {Code: 531, Message: "Authorization failed", Category: "Failure", Retryable: false},
	532: {Code: 532, Message: "Domain names linked with name server", Category: "Failure", Retryable: false},
	533: {Code: 533, Message: "Domain name has active name servers", Category: "Failure", Retryable: false},
	534: {Code: 534, Message: "Domain name has not been flagged for transfer", Category: "Failure", Retryable: false},
	535: {Code: 535, Message: "Restricted IP address", Category: "Failure", Retryable: false},
	536: {Code: 536, Message: "Domain already flagged for transfer", Category: "Failure", Retryable: false},
	540: {Code: 540, Message: "Attribute value is not unique", Category: "Failure", Retryable: false},
	541: {Code: 541, Message: "Invalid attribute value", Category: "Failure", Retryable: false},
	542: {Code: 542, Message: "Invalid old value for an attribute", Category: "Failure", Retryable: false},
	543: {Code: 543, Message: "Final or implicit attribute cannot be updated", Category: "Failure", Retryable: false},
	544: {Code: 544, Message: "Entity on hold", Category: "Failure", Retryable: false},
	545: {Code: 545, Message: "Entity reference not found", Category: "Failure", Retryable: false},
	546: {Code: 546, Message: "Credit limit exceeded", Category: "Failure", Retryable: false},
	547: {Code: 547, Message: "Invalid command sequence", Category: "Failure", Retryable: false},
	548: {Code: 548, Message: "Domain is not up for renewal", Category: "Failure", Retryable: false},
	549: {Code: 549, Message: "Command failed", Category: "Failure", Retryable: false},
	550: {Code: 550, Message: "Parent domain not registered", Category: "Failure", Retryable: false},
	551: {Code: 551, Message: "Parent domain status does not allow for operation", Category: "Failure", Retryable: false},
	552: {Code: 552, Message: "Domain status does not allow for operation", Category: "Failure", Retryable: false},
	553: {Code: 553, Message: "Operation not allowed. Domain pending transfer", Category: "Failure", Retryable: false},
	554: {Code: 554, Message: "Domain already registered", Category: "Failure", Retryable: false},
	555: {Code: 555, Message: "Domain already renewed", Category: "Failure", Retryable: false},
	556: {Code: 556, Message: "Maximum registration period exceeded", Category: "Failure", Retryable: false},
	880: {Code: 880, Message: "Failed to send request to registry", Category: "Failure", Retryable: false},
}

// LookupResponseInfo returns documentation info for a response code.
func LookupResponseInfo(code int) (ResponseInfo, bool) {
	info, ok := responseCodeTable[code]
	return info, ok
}

// Info returns the documentation info for this response, if known.
func (r Response) Info() (ResponseInfo, bool) {
	return LookupResponseInfo(r.ResponseCode)
}

type BaseResponse interface {
	Err() error
}

type Decodable interface {
	XMLProxy() BaseResponse
}

// Err returns an error if one occured at API level. This means a response was
// parsed successfully, but it contains an error message to be presented to the
// user.
func (r *Response) Err() error {
	if r.ResponseCode > 300 {
		// see also https://api.enom.com/docs/rrp-return-codes
		return fmt.Errorf("RRP error %d: %s", r.ResponseCode, r.ResponseText)
	}
	switch len(r.Errors) {
	case 0:
		return nil
	case 1:
		return fmt.Errorf("error from API: %s", r.Errors[0].Error)
	default:
		var err strings.Builder
		fmt.Fprintf(&err, "%d errors from API: ", len(r.Errors))
		err.WriteString(r.Errors[0].Error)
		for _, e := range r.Errors[1:] {
			fmt.Fprintf(&err, "; %s", e.Error)
		}
		return fmt.Errorf(err.String())
	}
}
