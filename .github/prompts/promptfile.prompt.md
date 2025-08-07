# GitHub Copilot Prompt: Go (Golang) with Hexagonal Architecture + AWS SAM

Generate Go code that follows idiomatic Go best practices and hexagonal architecture (ports and adapters), and is suitable for deployment as AWS Lambda functions using AWS SAM.

## 🧱 Architecture Style

- Follow hexagonal architecture:
  - `core`: business logic (pure Go)
  - `ports`: interfaces for input/output
  - `adapters`: HTTP handlers, repositories, etc.
  - `cmd`: app entry point (used by Lambda or local run)
- Use dependency injection
- Domain logic should be independent of AWS SDK

## ⚙️ AWS SAM Deployment

- Use AWS SAM for packaging and deploying
- Use `sam.yaml` for defining functions, roles, and events
- Go Lambda handler must be a standalone binary
- Use `Makefile` or `buildspec.yml` to build the Go binary
- Set handler to `bootstrap` and use `GOOS=linux GOARCH=arm64` for compilation

## ✅ Task Example

**Task**: Build a Go Lambda that receives an order from API Gateway.

- Domain: `Order` struct with `ID`, `CustomerID`, `Amount`
- Port: `OrderService` with `CreateOrder`
- Adapter: API Gateway input via Lambda (`adapter/http`)
- Entry point: `main.go` in `cmd/lambda`
- Add SAM template: `template.yaml` with function, role, and permissions

Include unit tests for domain logic, and logging using `log/slog`.

## 📦 Suggested Layout

```
/myapp
├── cmd/
│ └── lambda/
│ └── main.go
├── internal/
│ ├── core/
│ ├── ports/
│ └── adapters/
│ └── http/
├── pkg/
│ └── logger/
├── template.yaml # AWS SAM template
├── Makefile # Build script for Lambda binary
├── go.mod
└── README.md
```


