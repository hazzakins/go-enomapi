package response

type GetExtAttributes struct {
	Attributes   []ExtAttribute
	ResponseMeta ResponseMeta
}

type ExtAttribute struct {
	ID          int
	Name        string
	Application int
	UserDefined bool
	Required    int
	Description string
	IsChild     int
	Options     []ExtAttributeOption
}

type ExtAttributeOption struct {
	ID          int
	Value       string
	Title       string
	Description string
}

type GetIDNCodes struct {
	TLDs         []IDNCodesTLD
	ResponseMeta ResponseMeta
}

type IDNCodesTLD struct {
	TLD       string
	Languages []IDNLanguage
}

type IDNLanguage struct {
	Code string
	Name string
}

type TLDList []string

type TLDRegistrationDetails struct {
	IsRealTime    bool
	DNSRequired   bool
	DNSMinServers int
	DNSMaxServers int
}

type TLDDetails struct {
	TLD          string
	Lockable     bool
	SupportsIDN  bool
	Registration TLDRegistrationDetails
}
