package response

type TMCheck struct {
	LookupKey    string
	ResponseMeta ResponseMeta
}

type TMGetNotice struct {
	TcnID        string
	TcnStartDate string
	TcnExpDate   string
	SLD          string
	ClaimsXML    string
	ResponseMeta ResponseMeta
}

type TMUpdateCart struct {
	Success      string
	ResponseMeta ResponseMeta
}
