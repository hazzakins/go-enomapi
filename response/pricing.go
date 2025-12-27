package response

type PEGetTLDID struct {
	TLDID        string
	ResponseMeta ResponseMeta
}

type PEGetDomainPricing struct {
	DisplayPriceIncrease int
	Products             []PEDomainPricingProduct
	Count                int
	ResponseMeta         ResponseMeta
}

type PEDomainPricingProduct struct {
	TLD                string
	TLDID              int
	RegisterPrice      float64
	ResellerPriceReg   float64
	RegisterEnabled    bool
	RenewPrice         float64
	ResellerPriceRenew float64
	RenewEnabled       bool
	TransferPrice      float64
	ResellerPriceTran  float64
	TransferEnabled    bool
	RGPPrice           float64
	ResellerPriceRGP   float64
	RGPEnabled         bool
}
