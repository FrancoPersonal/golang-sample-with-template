---
applyTo: '**'
---
# GitHub Copilot Instructions: Go + Hexagonal Architecture + AWS SAM

These are the instructions for Copilot to generate clean, idiomatic Go code using hexagonal architecture, and prepare it for deployment using AWS SAM.

## 🔧 Language and Runtime

- Language: Go (version 1.24.5)
- Runtime: Compiled binary for AWS Lambda
- Build for `linux/arm64` platform
- Follow idiomatic Go practices
- Use `go mod` and structured project layout

## 🧱 Architecture

Use Hexagonal Architecture:
- `core`: domain models and business logic
- `ports`: input/output interfaces
- `adapters`: HTTP (Lambda/APIGW) handler and persistence
- `cmd/lambda/main.go`: entry point for Lambda

## 🚀 AWS SAM Deployment

- Provide a `template.yaml` with:
  - `AWS::Serverless::Function` resources
  - API Gateway integration
  - IAM roles and permissions
- Define `Handler: bootstrap`
- Add a `Makefile` to build with:
```bash
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap cmd/lambda/main.go
```

## 🧪 Testing

- Use Go's `testing` package
- Prefer table-driven tests
- Write unit tests for domain and ports
- Use `testify` or mocks where needed
- Use `t.Helper()` in test helpers

## 📁 Project Layout

```
/project-root
├── cmd/lambda/
│ └── main.go
├── internal/core/
├── internal/ports/
├── internal/adapters/http/
├── pkg/logger/
├── template.yaml
├── Makefile
├── go.mod
└── tests/
```

## 💡 Copilot May Generate

- Lambda-compatible HTTP handlers (API Gateway -> service interface)
- Domain logic and data validation
- SAM template with Lambda definitions
- Logging setup using `log/slog`
- Dependency injection in main

## 🧼 Copilot Should Avoid

- Using `fmt.Println` for production logs
- Placing business logic inside handlers
- Ignoring returned errors
- Using `panic()` outside of initialization
