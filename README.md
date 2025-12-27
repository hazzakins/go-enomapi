
# go-enomapi

A lightweight Go client for the ENOM reseller API.

Inspiration taken from https://git.sr.ht/~bitfehler/go-enom/

## Overview

`go-enomapi` is an early-stage client focused on a small, usable core. Current coverage includes domain availability and registration flows, renewals, transfers, pricing/TLD metadata, and a broad set of domain-management helpers. The API surface is still evolving.

## Installation

Add the module to your project:

```bash
go get github.com/hazzakins/go-enomapi
```

## Quick Start

Example usage showing authentication + a domain availability check:

```go
package main

import (
	"fmt"
	"log"

	enom "github.com/hazzakins/go-enomapi"
	"github.com/hazzakins/go-enomapi/domains"
)

func main() {
	client, err := enom.NewClient(
		"https://resellertest.enom.com/",
		"YOUR_RESELLER_ID",
		"YOUR_API_KEY",
	)
	if err != nil {
		log.Fatalf("failed to create enom client: %v", err)
	}

	domainsClient := &domains.Client{Client: client}
	resp, err := domainsClient.Check(enom.NewDomain("example.com"))
	if err != nil {
		log.Fatalf("request error: %v", err)
	}

	fmt.Printf("Available: %t\n", resp.IsAvailable)
}
```

## Configuration & Authentication

- Use API key and reseller ID as provided by ENOM.
- Point `NewClient` at the sandbox or production reseller URL; this selects sandbox vs production.
- The client uses `http.Get` directly today, so there is not yet a custom transport or context-aware timeout control.

Example sandbox base URL:

```
https://resellertest.enom.com/
```

## Available APIs (current)

- Domains: availability (`Check`), registration (`Purchase`, `PurchaseWithOptions`, `AddBulkDomains`, `Preconfigure`, `DeleteRegistration`)
- Orders/queues: `CancelOrder`, `GetConfirmationSettings`, `GetAgreementPage`, `QueueDomainPurchase`, `QueueGet*`, `QueueGetInfo`
- Renewals: `Extend`, `ExtendRGP`, `GetRenew`, `GetExtendInfo`, `InsertNewOrder`, `SetRenew`, `UpdateExpiredDomains`, `UpdateRenewalSettings`
- Transfers: `PushDomain`, `SynchAuthInfo`, `TP*` transfer endpoints, `UpdatePushList`
- Domain management: contacts and whois, domain info/status/search, locks/passwords, portal helpers, reports
- DNS & name servers: `GetDNS`, `SetDNSHost`, `ModifyNS`, `RegisterNameServer`, DNSSEC, host records (SPF/SRV/meta/hosts)
- TLD metadata & pricing: `GetTLDList`, `GetTLDDetails`, `GetIDNCodes`, `GetExtAttributes`, `PEGetDomainPricing`, `PEGetTLDID`, `PESetPricing`, reseller pricing helpers
- Suggestions & trademark: `GetNameSuggestions`, `NameSpinner`, `TMCheck`, `TMGetNotice`, `TMUpdateCart`

## Not yet implemented

- Account management, accounting/reporting, DNS hosting, email hosting, shopping cart, SSL certificates, value-added services, and whois publicity service APIs
- Domain watchlists/overview, aftermarket (NM_*), TEL_*, and XXX_* endpoints

## CLI (experimental)

An interactive CLI lives at `cmd/enomcli`. It currently supports:

- Checking availability and price details
- Fetching domain info

It expects `RESELLERID` and `APIKEY` in a `.env` file.

## Contributing

Please see `PROJECTS.md` for the project roadmap and baseline tasks. Contributions are welcome — open issues or pull requests and follow standard Go project practices.

