package response

// Extend reports the outcome of a domain renewal request.
type Extend struct {
	Extension       string
	DomainName      string
	OrderID         string
	RegistryExpDate string
	UseCreditCard   bool
	TotalCharged    float64
	ResponseMeta    ResponseMeta
}

// ExtendRGP reports the outcome of a redemption grace period renewal request.
type ExtendRGP struct {
	Extension    string
	DomainName   string
	OrderID      string
	ResponseMeta ResponseMeta
}

// InsertNewOrder reports the outcome of a shopping cart checkout.
type InsertNewOrder struct {
	OrderID              string
	PremiumDomainOrderID string
	ProductTypes         []InsertNewOrderProductType
	ResponseMeta         ResponseMeta
}

// InsertNewOrderProductType captures a product type entry returned by InsertNewOrder.
type InsertNewOrderProductType struct {
	Index       int
	ProductType string
}

// UpdateExpiredDomains reports the outcome of reactivating an expired domain.
type UpdateExpiredDomains struct {
	Status       bool
	OrderID      string
	ResponseMeta ResponseMeta
}

// UpdateRenewalSettings reports the outcome of updating renewal notifications.
type UpdateRenewalSettings struct {
	AcceptTermsStatus string
	RenewalSetting    string
	ResponseMeta      ResponseMeta
}
