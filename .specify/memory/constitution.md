<!--
Sync Impact Report:
- Version change: 1.0.0 -> 1.1.0
- Modified principles:
  - I. Provider Decoupling & Manager-Driven Routing: Expanded with complete channel/provider matrix (Email, SMS, WhatsApp, Push Notification, Telegram) and specific manager strategies (IPWarmimgEmailManager, RandomManager, Failover).
  - II. Distributed Observability & End-to-End Tracing: Detailed OpenTelemetry/Uptrace integration (port 4318), KafkaHeaderCarrier context extraction in api/kafka/index.go, child span lineage from authentication/billing, and otelzap logging.
  - III. Zero-Credential Exposure & Mock-First Testing: Explicit prohibition against committing credentials for SendGrid, Twilio, AWS, Mailgun, OneSignal, FCM, Telegram; formalized mock usage in test/mocks/.
  - IV. Resilient Asynchronous Event Processing -> IV. Reactive Kafka Event Processing & Topic Contracts: Formalized operational topics (2FA_EMAIL_CREATED, 2FA_PHONE_CREATED, 2FA_EMAIL_SENT, 2FA_PHONE_SENT, USER_CREATED, ORGANIZATION_CREATED) and concurrent consumer execution via segmentio/kafka-go.
  - V. Contract Integrity & Schema Validation -> V. Synchronous API Contracts & Validation: Formalized HTTP REST endpoints (POST /email, POST /sms, POST /push-notification) on port 5001 with exact payload contracts.
- Added sections:
  - Refined Security & Operational Standards with runtime binaries (cmd/http/http_server.go, cmd/kafka/kafka_server.go).
  - Refined Development Workflow & Quality Gates with AGENTS.md test commands and linting specifications.
- Removed sections: None
- Follow-up TODOs: None
-->

# Auth+ Notification Service Constitution

## Core Principles

### I. Provider Decoupling & Manager-Driven Routing
Notification dispatching MUST NOT bind domain use cases directly to individual third-party provider implementations. All dispatch operations MUST pass through Driven Adapter interfaces (`internal/usecases/driven`) and be coordinated by strategy-based Managers (`internal/managers/`):
- **Channel & Provider Matrix**:
  - **Email**: SendGrid, Mailgun, OneSignal
  - **SMS**: Twilio, AWS SNS
  - **WhatsApp**: Twilio API
  - **Push Notification**: Firebase Cloud Messaging (FCM), OneSignal
  - **Telegram**: Telegram Bot API
- **Routing Strategies**:
  - `IPWarmimgEmailManager`: Enforces gradual IP reputation warm-up for high-volume email delivery, distributing percentage quotas across providers (e.g., priority quota on SendGrid, intermediate quota on OneSignal, fallback on Mailgun).
  - `RandomManager`: Balances traffic and vendor rate limits stochastically across available providers.
  - **Resilience & Failover**: If a designated provider fails, exhausts quotas, or returns retryable errors, the manager MUST select configured contingency routes without crashing the caller.
- **Rationale**: Direct vendor coupling creates systemic fragility, vendor lock-in, and prevents traffic steering or reputation management.

### II. Distributed Observability & End-to-End Tracing
All entry points—both synchronous REST handlers and asynchronous Kafka consumers—MUST maintain and propagate W3C distributed trace contexts using the OpenTelemetry Go SDK (`go.opentelemetry.io/otel`) integrated with Uptrace (`uptrace:4318` via HTTP/gRPC):
- **Kafka Context Extraction**: In Kafka consumers (`api/kafka/index.go`), trace context MUST be extracted from event headers via `KafkaHeaderCarrier` so that delivery spans in Uptrace are direct children of the originating upstream requests (e.g., `authentication`, `billing`).
- **Structured Logging**: Logging MUST use `otelzap` correlated with the active trace and span identifiers.
- **Data Anonymization**: External dispatch logs MUST record anonymized/masked recipient identifiers (emails, phone numbers, tokens) alongside provider HTTP return statuses; raw secrets or plain authentication challenge codes MUST NEVER appear in log streams.
- **Rationale**: Asynchronous notification flows span service boundaries; continuous distributed tracing and sanitized audit trails are non-negotiable for production triage and privacy compliance.

### III. Zero-Credential Exposure & Mock-First Testing
Under no circumstances shall production or third-party credentials (SendGrid, Mailgun, Twilio, AWS SNS, OneSignal, Firebase, Telegram) be committed to the repository:
- **Secret Decoupling**: All keys and credentials MUST be injected exclusively via environment variables (`.env`).
- **Mock-Driven Testing**: Unit and integration test suites MUST utilize emulated transports or mocks located in `test/mocks/`. Test executions MUST NOT invoke real external provider APIs or require live third-party connectivity.
- **Testing Verification**: Test suites MUST execute cleanly via `go test -v ./...` and verify test coverage via `go test -coverprofile=coverage.out ./...`.
- **Rationale**: Credential leakage is a severe security compromise, and live third-party dependencies during testing introduce flakiness, rate limiting, and CI delays.

### IV. Reactive Kafka Event Processing & Topic Contracts
Asynchronous message consumption MUST be managed concurrently via `segmentio/kafka-go` running in `cmd/kafka/kafka_server.go`, strictly adhering to platform event contracts:
- **Mandatory Operational Topics**:
  - `2FA_EMAIL_CREATED` / `2FA_PHONE_CREATED`: Dispatches verification challenge links and security challenges.
  - `2FA_EMAIL_SENT` / `2FA_PHONE_SENT`: Dispatches OTP authentication codes to target recipients.
  - `USER_CREATED`: Dispatches onboarding instructions and welcome messages.
  - `ORGANIZATION_CREATED`: Dispatches administrative provisioning notices for new tenant onboarding.
- **Idempotency & Resilience**: Consumers MUST handle duplicate message deliveries gracefully and prevent duplicate customer-facing notifications.
- **Rationale**: Event-driven notification pipelines drive core authentication and user lifecycles; topic schema adherence and fault tolerance ensure system consistency.

### V. Synchronous API Contracts & Validation
Synchronous dispatch routes exposed by `cmd/http/http_server.go` on port `5001` MUST strictly validate payloads before passing execution to domain use cases:
- **Standard Endpoint Contracts**:
  - `POST /email`: Requires `{ "email": "...", "subject": "...", "content": "..." }`
  - `POST /sms`: Requires `{ "phone": "...", "content": "..." }`
  - `POST /push-notification`: Requires `{ "device_id": "...", "content": "..." }`
- **Error Handling**: Malformed or invalid requests MUST be rejected immediately with appropriate 4xx HTTP responses and descriptive validation errors before any provider manager is contacted.
- **Rationale**: Front-line validation prevents corrupt data from propagating into downstream delivery pipelines and incurring unnecessary third-party API costs.

## Security & Operational Standards

- **Runtime & Deployment**: The microservice runs on Go 1.20+ and Docker/Docker Compose. Binaries are divided into:
  - HTTP Server: `cmd/http/http_server.go` (listening on port `5001` in both host and container).
  - Kafka Consumer: `cmd/kafka/kafka_server.go` (connecting to `kafka:9092`).
- **Telemetry Infrastructure**: OpenTelemetry collector connection to Uptrace (`uptrace:4318`) via HTTP or gRPC.
- **PII Protection**: Recipient emails, phone numbers, and device tokens MUST be treated as sensitive data and masked in all observability outputs.

## Development Workflow & Quality Gates

- **Architectural Integrity**: Clean Architecture boundaries MUST be respected. Use cases depend only on driven adapter interfaces (`internal/usecases/driven`), never on concrete provider packages (`internal/providers/`).
- **Test Coverage Gate**: Pull requests MUST pass `go test -v ./...` with coverage reporting (`go test -coverprofile=coverage.out ./...`).
- **Linting & Formatting**: Code MUST comply with Go standard tooling (`go fmt`, `go vet`) and pass revive linter checks (`$HOME/go/bin/revive -formatter friendly ./...`).
- **Trace Context Integrity**: Any new Kafka consumer routine or HTTP route MUST maintain W3C trace context extraction and span creation.

## Governance

- This Constitution serves as the primary governance document for `auth-plus-notification`, fully aligned with `AGENTS.md`. It supersedes informal or contradictory practices.
- Amendments MUST follow Semantic Versioning:
  - **MAJOR**: Incompatible governance shifts, architectural redefinitions, or removal of core principles.
  - **MINOR**: Addition of new principles, sections, or materially expanded operational rules.
  - **PATCH**: Non-semantic clarifications, wording refinements, and typo corrections.
- All code reviews, pull requests, and automated AI agent contributions MUST verify strict compliance against these principles.

**Version**: 1.1.0 | **Ratified**: 2026-10-12 | **Last Amended**: 2026-10-12
