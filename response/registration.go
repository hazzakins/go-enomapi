package response

import "time"

type AddBulkDomains struct {
	Items              []AddBulkDomainsItem
	ListCount          int
	CartErrors         int
	AllItemsSuccessful string
	CartItems          int
	Success            string
	UseCart            string
	ResponseMeta       ResponseMeta
}

type AddBulkDomainsItem struct {
	WscAccountOverride string
	ItemName           string
	ItemID             string
	Price              string
	ICANNFees          string
	CartItemID         string
	NewDomainNameID    string
	ItemAdded          string
	ItemError          string
	DomainName         string
}

type Preconfigure struct {
	PreConfigSuccess string
	Count            int
	ResponseMeta     ResponseMeta
}

type DomainPurchase struct {
	// If this is not "" then the order was placed successfully
	OrderID string
	// If this is false the order was queued. Check OrderStatus and OrderDescription.
	OrderCompleted bool
	OrderStatus    string
	// Only set if OrderCompleted == true
	OrderDescription string
	// Only set if OrderCompleted == true
	RegistrationDate time.Time
	// Only set if OrderCompleted == true
	ExpirationDate time.Time
	// The total amount charged for this purchase
	Price float64
}

type DeleteRegistration struct {
	DomainDeleted string
	ErrString     string
	ErrSource     string
	ErrSection    string
	ResponseMeta  ResponseMeta
}
