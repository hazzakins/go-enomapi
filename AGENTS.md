## Agent Guide for go-enomapi

This document explains how to work in this repository and how to extend the
library in an idiomatic, low-friction way.

## Project Intent

`go-enomapi` is a lightweight, idiomatic Go client for the ENOM reseller API.
The initial focus is a small, usable core: client config/auth, domains API,
and tests. Additional APIs (contacts, orders, DNS, TLD metadata) follow after
the core is stable.

## Roadmap Snapshot (priority order)

1. Client core (Client type, config, constructors, transport abstraction)
2. Authentication (API key/username, sandbox vs production)
3. Domains API (availability, registration, renewals, transfers)
4. Testing (unit tests with mocks, integration tests in sandbox)
5. Error handling and retries (typed errors, backoff, rate limiting)
6. Remaining APIs (contacts, orders, DNS, TLD metadata)
7. Logging/observability, examples, and documentation polish

## Development Workflow

- Prefer small, focused changes; keep the public surface minimal and stable.
- Expose context-aware methods (`ctx context.Context`) for all network calls.
- Keep configuration explicit via a `Config` struct and a `NewClient` constructor.
- Favor typed responses and errors; avoid `map[string]any` in exported APIs.

## Code Organization (planned)

- `client.go`: Client, Config, constructor, base request logic
- `transport/` or `internal/transport/`: HTTP transport abstraction
- `domains/`: domain operations and request/response types
- `errors.go`: typed errors, retry classification
- `internal/`: internal helpers (auth signing, query building, parsing)

Adjust as needed, but keep packages cohesive and avoid deep nesting.

## Go Style and Best Practices

- Follow standard Go formatting and naming; run `gofmt` on changed files.
- Use small interfaces; avoid interface-heavy design in public APIs.
- Return sentinel or typed errors and wrap with `fmt.Errorf("...: %w", err)`.
- Validate inputs at boundaries; do not panic on user input.
- Avoid global state; inject timeouts and transports via config.

## Tests

- Unit tests should mock the transport layer; prefer table-driven tests.
- Integration tests should target sandbox mode and be opt-in.
- Aim for deterministic tests; avoid network access in unit tests.

## Documentation

- Keep README accurate and concise.
- Add GoDoc comments to exported types and methods.
- Provide small usage examples as APIs are implemented.

## Quick Commands

These are suggestions; add a Makefile target if you want a single entry point.

```bash
go test ./...
go vet ./...
gofmt -w .
```
