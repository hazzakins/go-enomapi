## Project Roadmap and Baseline Tasks

This file contains a baseline set of tasks to take `go-enomapi` from an idea to a usable, production-ready Go client for the ENOM API.

- **1. Project init**: repository scaffolding, `go.mod`, CI basics, code formatting/lint rules.
- **2. Client core**: implement `Client` type, configuration struct, constructors, HTTP transport abstraction.
- **3. Authentication**: implement ENOM authentication (API key/username), sandbox vs production modes.
- **4. Domains API**: domain search/availability, registration, renewals, transfers.
- **5. Contacts API**: create/read/update contact records.
- **6. Orders API**: create orders, check status, list orders.
- **7. DNS API**: manage DNS records (CRUD).
- **8. TLD metadata**: list supported TLDs, pricing and capabilities.
- **9. Error handling**: typed errors, retryable vs fatal errors, backoff strategy.
- **10. Rate limiting & retries**: global rate limiter, configurable retry/backoff.
- **11. Logging & observability**: pluggable logger interface, request tracing.
- **12. Testing**: unit tests with mocks, integration tests (sandbox), CI test matrix.
- **13. Examples**: small example programs demonstrating common flows.
- **14. Documentation**: README (this file), GoDoc comments, usage examples.
- **15. Release process**: semantic versioning, changelog, tagging, and release notes.

Notes
- Prioritize a small, usable subset (Client core, Authentication, Domains API, and Tests) for the first release.
- Keep the API surface small and idiomatic; prefer explicit structs and context-aware methods.

## Outstanding endpoints (from docs index pages)

These are endpoints documented in `docs/` index pages that do not yet have client implementations. Reference docs (for example `api-error-codes.md`) are excluded.

### Account management
- AddContact
- AuthorizeTLD
- CreateSubAccount
- DeleteSubaccount
- GetAccountInfo
- GetAccountPassword
- GetAccountValidation
- GetCusPreferences
- GetCustomerPaymentInfo
- GetGlobalChangeStatus
- GetGlobalChangeStatusDetail
- GetOrderDetail
- GetOrderList
- GetReport
- GetResellerInfo
- GetServiceContact
- GetSubAccountDetails
- GetSubAccounts
- GetTransHistory
- GetWebHostingAll
- RAA_GetInfo
- RAA_ResendNotifications
- RemoveTLD
- SendAccountEmail
- SubAccountDomains
- UpdateAccountInfo
- UpdateCusPreferences
- UpdateRenewalSettings
- CheckLogin
- CreateAccount
- GetAllAccountInfo
- DeleteCustomerDefinedData
- GetCustomerDefinedData
- SetCustomerDefinedData

### Accounting and reports
- CommissionAccount
- GetBalance
- PE_GetCustomerPricing
- PE_GetDomainPricing
- PE_GetEapPricing
- PE_GetPremiumPricing
- PE_GetPOPPrice
- PE_GetProductPrice
- PE_GetResellerPrice
- PE_GetRetailPrice
- PE_GetRetailPricing
- PE_GetRocketPrice
- RefillAccount
- UpdateNotificationAmount
- DeleteCustomerDefinedData
- GetCustomerDefinedData
- GetOrderDetail
- GetOrderList
- GetReport
- GetTransHistory
- SetCustomerDefinedData

### Domains
- AddBulkDomains
- CancelOrder
- GetConfirmationSettings
- GetExtAttributes
- GetIDNCodes
- GetNameSuggestions
- Preconfigure
- Queue_DomainPurchase
- Queue_GetDomains
- Queue_GetExtAttributes
- Queue_GetOrderDetail
- Queue_GetOrders
- TM_Check
- TM_GetNotice
- TM_UpdateCart
- GetAgreementPage
- Queue_GetInfo
- DeleteRegistration
- Extend
- Extend_RGP
- GetDomainExp
- GetExtendInfo
- GetRenew
- InsertNewOrder
- SetRenew
- UpdateExpiredDomains
- UpdateRenewalSettings
- PE_GetTLDID
- PE_SetPricing
- PushDomain
- RefillAccount
- SetResellerServicesPricing
- SetResellerTLDPricing
- SynchAuthInfo
- TP_CancelOrder
- TP_CreateOrder
- TP_GetDetailsByDomain
- TP_GetOrder
- TP_GetOrderDetail
- TP_GetOrdersByDomain
- TP_GetOrderReview
- TP_GetOrderStatuses
- TP_GetTLDInfo
- TP_ResendEmail
- TP_ResubmitLocked
- TP_SubmitOrder
- TP_UpdateOrderDetail
- UpdateAccountPricing
- UpdatePushList
- GetDomainSRVHosts
- GetHosts
- GetMetaTag
- GetRegHosts
- GetSPFHosts
- SetDomainSRVHosts
- SetHosts
- SetSPFHosts
- UpdateMetaTag
- AddDnsSec
- CheckNSStatus
- DeleteDnsSec
- DeleteNameServer
- GetDNS
- GetDnsSec
- GetDNSStatus
- GetHomeDomainList
- ModifyNS
- ModifyNSHosting
- RegisterNameServer
- SetDNSHost
- UpdateNameServer
- AddDomainFolder
- DeleteDomainFolder
- GetDomainFolderDetail
- GetDomainFolderList
- RemoveUnsyncedDomains
- UpdateDomainFolder
- AssignToDomainFolder
- TLD_AddWatchlist
- TLD_DeleteWatchlist
- TLD_GetWatchlist
- TLD_GetWatchlistTlds
- TLD_GetTLD
- TLD_Overview
- NM_CancelOrder
- NM_ExtendOrder
- NM_GetPremiumDomainSettings
- NM_GetSearchCategories
- NM_ProcessOrder
- NM_Search
- NM_SetPremiumDomainSettings
- XXX_GetMemberId
- XXX_RemoveMemberId
- XXX_SetMemberId
- TEL_AddCTHUser
- TEL_GetCTHUserInfo
- TEL_GetCTHUserList
- TEL_GetPrivacy
- TEL_IsCTHUser
- TEL_UpdateCTHUser
- TEL_UpdatePrivacy

### DNS hosting
- DeleteHostedDomain
- ExtendDomainDNS
- PurchaseServices
- SetDNSHost

### Email hosting
- DeleteAllPOPPaks
- DeletePOP3
- DeletePOPPak
- DisableServices
- EnableServices
- Forwarding
- GetCatchAll
- GetDotNameForwarding
- GetForwarding
- GetMailHosts
- GetPOP3
- GetPOPExpirations
- GetPOPForwarding
- ModifyPOP3
- PurchasePOPBundle
- RenewPOPBundle
- SetCatchAll
- SetDotNameForwarding
- SetPakRenew
- SetPOPForwarding
- SetUpPOP3User

### Shopping cart
- AddBulkDomains
- AddToCart
- DeleteFromCart
- GetCartContent
- InsertNewOrder
- PurchasePreview
- UpdateCart

### SSL certificates
- CertChangeApproverEmail
- CertConfigureCert
- CertGetApproverEmail
- CertGetCertDetail
- CertModifyOrder
- CertParseCSR
- CertPurchaseCert
- CertReissueCert
- CertResendApproverEmail
- CertResendFulfillmentEmail
- GetCerts
- AM_AutoRenew
- AM_Configure
- AM_GetAccounts
- AM_GetAccountDetail

### Value-added services
- DisableServices
- EnableServices
- GetDomainServices
- GetDomainSubServices
- GetHomeDomainList
- GetIPResolver
- PurchaseServices
- RenewServices
- ServiceSelect
- SetDomainSubServices
- SetIPResolver

### Whois publicity service
- PurchaseServices
- AddBulkDomains
- AddToCart
- VAS_Update
- VAS_GetList
- VAS_GetDetail
