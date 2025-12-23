
# go-enomapi

A lightweight Go client for the ENOM API to be used by other Go projects.

Inspiration Taken from https://git.sr.ht/~bitfehler/go-enom/

## Overview

`go-enomapi` provides a small, well-documented client for interacting with the ENOM reseller API. The goal is to offer a simple, idiomatic Go interface for common ENOM operations (domains, contacts, orders, DNS, and TLD information) with sensible defaults for retries, logging, and error handling.

## Installation

Add the module to your project:

```bash
go get github.com/hazzakins/go-enomapi
```

## Quick Start

Example usage showing the core flow (authentication + basic request):

```go
package main

import (
	"context"
	"fmt"
	"log"

	enom "github.com/hazzakins/go-enomapi"
)

func main() {
	ctx := context.Background()

	// Create client with API credentials (example constructor)
	cfg := enom.Config{
		Username: "YOUR_ENOM_USERNAME",
		ApiKey:   "YOUR_API_KEY",
		Sandbox:  true,
	}

	client, err := enom.NewClient(cfg)
	if err != nil {
		log.Fatalf("failed to create enom client: %v", err)
	}

	// Example: Check domain availability
	resp, err := client.Domains().CheckAvailability(ctx, "example.com")
	if err != nil {
		log.Fatalf("request error: %v", err)
	}

	fmt.Printf("Availability: %+v\n", resp)
}
```

Replace the constructor and method names above with the actual API once implemented; these are representative of the intended surface.

## Configuration & Authentication

- Use API key and username as provided by ENOM.
- Support for sandbox/test mode will be added (see `projects.md`).

## Features (planned)

- Domain search and registration
- Contact management
- Order creation and status
- DNS management
- TLD metadata and pricing
- Retry/backoff and rate-limit handling
- Context-aware requests using `context.Context`

## Contributing

Please see `projects.md` for the project roadmap and baseline tasks. Contributions are welcome — open issues or pull requests and follow standard Go project practices.

## License

This project will use an OSI-approved license; add a `LICENSE` file when ready.
