# 🧠 GitHub Copilot Mode: Go Developer (v1.24.5) + Hexagonal + AWS SAM

Copilot should assist in generating Go code using version `1.24.5`, following idiomatic Go patterns, Hexagonal Architecture, and AWS SAM for serverless deployment.

---

## 🔧 Language Settings

- Language: Go
- Version: 1.24.5
- Use Go Modules
- Follow [Effective Go](https://go.dev/doc/effective_go)

---

## 🧱 Architecture: Hexagonal + Serverless

- Domain (`internal/core`) must remain free of AWS-specific code
- Ports (`internal/ports`) define input/output contracts
- Adapters (`internal/adapters/http`) implement ports using Lambda
- Entry point (`cmd/lambda/main.go`) wires dependencies for Lambda
- Avoid coupling domain with infrastructure

---

## 🚀 AWS SAM Deployment Integration

- Use AWS SAM CLI to deploy functions
- Define functions in `template.yaml`:
  ```yaml
  Resources:
    OrderFunction:
      Type: AWS::Serverless::Function
      Properties:
        Handler: bootstrap
        Runtime: provided.al2
        CodeUri: .
        MemorySize: 128
        Timeout: 10
        Events:
          Api:
            Type: Api
            Properties:
              Path: /orders
              Method: post

## Use Makefile to compile for AWS:

```bash
build:
  GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ./cmd/lambda/main.go
```

## ✅ Copilot Should Prioritize
- Writing idiomatic, testable Go
- Using interfaces to decouple logic
- Wrapping errors with context

## Injecting dependencies

- Creating clear, maintainable handlers for AWS Lambda

## 🧪 Testing
- Table-driven tests with the testing package
- testify or mocks for service isolation
- Test core logic without AWS dependencies

## 🧼 Copilot Should Avoid
- Tight coupling between core logic and AWS SDK
- Business logic inside Lambda handler
- Ignoring errors
- Using global state

# 📦 Project Layout
```plaintext
/project-root
├── cmd/lambda/            # Lambda entrypoint (main.go)
├── internal/core/         # Business logic
├── internal/ports/        # Interfaces
├── internal/adapters/http # Lambda handler
├── pkg/logger/            # Logging helpers
├── template.yaml          # AWS SAM template
├── Makefile               # Build instructions
├── go.mod
└── tests/                 # Integration/unit tests
```

Copilot should generate clean, modular, testable Go code designed for real-world AWS Lambda deployments using SAM and hexagonal architecture principles.