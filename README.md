
# go-enomapi

A lightweight Go client for the ENOM reseller API.

Inspiration taken from https://git.sr.ht/~bitfehler/go-enom/

## Overview

`go-enomapi` is an early-stage client focused on a small, usable core. Current coverage includes basic domain checks, TLD metadata, and a handful of domain-management helpers. The API surface is still evolving.

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
- Point `NewClient` at the sandbox or production reseller URL.

Example sandbox base URL:

```
https://resellertest.enom.com/
```

## Available APIs (current)

- Domain availability check (`Check`)
- Domain purchase (`Purchase`) (partial)
- TLD list/details (`GetTLDList`, `GetTLDDetails`)
- Name spinner (`NameSpinner`)
- Domain info (`GetDomainInfo`)

Most endpoints are listed in `docs/` but are not yet implemented.

## CLI (experimental)

An interactive CLI lives at `cmd/enomcli`. It currently supports:

- Checking availability and price details
- Fetching domain info

It expects `RESELLERID` and `APIKEY` in a `.env` file.

## Contributing

Please see `PROJECTS.md` for the project roadmap and baseline tasks. Contributions are welcome — open issues or pull requests and follow standard Go project practices.

## License

This project will use an OSI-approved license; add a `LICENSE` file when ready.
