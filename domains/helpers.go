package domains

import (
	"strconv"

	"github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/internal"
	"github.com/hazzakins/go-enomapi/response"
)

func addBoolParam(cmd *enomapi.Command, key string, value *bool) {
	if value == nil {
		return
	}
	cmd.AddParam(key, boolToTitle(*value))
}

func addBoolIntParam(cmd *enomapi.Command, key string, value *bool) {
	if value == nil {
		return
	}
	cmd.AddParam(key, boolToInt(*value))
}

func addIntParam(cmd *enomapi.Command, key string, value *int) {
	if value == nil {
		return
	}
	cmd.AddParam(key, strconv.Itoa(*value))
}

func boolToTitle(value bool) string {
	if value {
		return "True"
	}
	return "False"
}

func boolToInt(value bool) string {
	if value {
		return "1"
	}
	return "0"
}

func addParams(cmd *enomapi.Command, params map[string]string) {
	for key, value := range params {
		if value == "" {
			continue
		}
		cmd.AddParam(key, value)
	}
}

func toTransferResponse(resp *internal.GenericResponse) *response.TransferResponse {
	decoded := resp.Decode()
	return &response.TransferResponse{ResponseMeta: decoded.ResponseMeta}
}
