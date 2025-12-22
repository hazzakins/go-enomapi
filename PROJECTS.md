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
