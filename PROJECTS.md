## Project Roadmap and Baseline Tasks

This file tracks the state of `go-enomapi` and the remaining work to reach a
usable, production-ready Go client for the ENOM API.

## Current state (snapshot)

- Client core: `Client`, `Command`, XML decoding, and basic error handling are in place; no config struct, transport abstraction, or context-aware methods yet.
- Authentication: reseller ID + API key are supported; sandbox vs production is controlled by the base URL passed to `NewClient`.
- Domains API: availability, registration, orders/queues, renewals, transfers, pricing, suggestions, trademark, and domain-management helpers are implemented.
- CLI: experimental CLI at `cmd/enomcli` for availability checks and domain info.
- Tests: no unit or integration tests yet.

## Roadmap status

- [x] 1. Project init
- [~] 2. Client core (needs Config struct, transport abstraction, context-aware methods)
- [~] 3. Authentication (basic credentials supported; explicit sandbox/prod config still needed)
- [~] 4. Domains API (broad coverage; see gaps below)
- [~] 5. Contacts API (domain contacts implemented; account-level contacts pending)
- [~] 6. Orders API (domain order/queue endpoints implemented; account order endpoints pending)
- [~] 7. DNS API (domain DNS/name servers/host records implemented; DNS hosting API pending)
- [~] 8. TLD metadata (list/details/IDN/ext attrs implemented; watchlist/overview pending)
- [ ] 9. Error handling
- [ ] 10. Rate limiting & retries
- [ ] 11. Logging & observability
- [ ] 12. Testing
- [ ] 13. Examples
- [~] 14. Documentation (README and GoDoc are partial)
- [ ] 15. Release process

## Outstanding endpoints (high-level)

### Domain-related gaps

- Aftermarket (NM_*) commands
- TEL_* commands
- XXX_* commands
- Domain watchlist and overview (TLD_* watchlist/overview)
- Magic folders (AddDomainFolder, DeleteDomainFolder, GetDomainFolderDetail, GetDomainFolderList, RemoveUnsyncedDomains, UpdateDomainFolder, AssignToDomainFolder)

### Non-domain APIs not started

- Account management
- Accounting and reports
- DNS hosting
- Email hosting
- Shopping cart
- SSL certificates
- Value-added services
- Whois publicity service

Notes
- Prioritize a small, usable subset (client core, authentication, domains API, and tests) for the first release.
- Keep the public API surface small and idiomatic; prefer explicit structs and context-aware methods.
