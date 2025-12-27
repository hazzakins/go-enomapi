package response

// TransferResponse is a minimal response wrapper for transfer and pricing
// commands that do not currently expose additional structured fields beyond the
// standard metadata.
type TransferResponse struct {
	ResponseMeta ResponseMeta
}

type PESetPricing = TransferResponse
type PushDomain = TransferResponse
type RefillAccount = TransferResponse
type SetResellerServicesPricing = TransferResponse
type SetResellerTLDPricing = TransferResponse
type SynchAuthInfo = TransferResponse
type TPCancelOrder = TransferResponse
type TPCreateOrder = TransferResponse
type TPGetDetailsByDomain = TransferResponse
type TPGetOrder = TransferResponse
type TPGetOrderDetail = TransferResponse
type TPGetOrdersByDomain = TransferResponse
type TPGetOrderReview = TransferResponse
type TPGetOrderStatuses = TransferResponse
type TPGetTLDInfo = TransferResponse
type TPResendEmail = TransferResponse
type TPResubmitLocked = TransferResponse
type TPSubmitOrder = TransferResponse
type TPUpdateOrderDetail = TransferResponse
type UpdateAccountPricing = TransferResponse
type UpdatePushList = TransferResponse
