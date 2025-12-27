package response

type GetNameSuggestions struct {
	SuggestionsXML string
	ResponseMeta   ResponseMeta
}

type NameSpinner struct {
	SpinCount    int
	TLDList      string
	OriginalSLD  string
	Domains      []SpinnerDomain
	ResponseMeta ResponseMeta
}

type SpinnerDomain struct {
	Name     string
	Com      string
	ComScore string
	Net      string
	NetScore string
	Tv       string
	TvScore  string
	Cc       string
	CcScore  string
}
