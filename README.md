# cloud-resource-manager

A lightweight Go CLI for automating AWS resource provisioning and Kubernetes operations through reusable commands.

Built as a backend engineering project to demonstrate Go, CLI development, concurrency, structured logging, configuration management, Docker, and software engineering best practices.

---

## Features

- AWS resource management commands
- Kubernetes resource operations
- Mock cloud providers for local development
- Structured logging with Zap
- Configuration management with Viper
- Docker support
- Unit testing with Testify
- Modular architecture using interfaces

---

## Project Structure

```text
cloud-resource-manager/
├── cmd/
│   ├── aws.go
│   ├── config.go
│   ├── k8s.go
│   ├── root.go
│   └── version.go
│
├── internal/
│   ├── aws/
│   ├── config/
│   ├── kubernetes/
│   ├── logger/
│   ├── models/
│   ├── services/
│   └── utils/
│
├── configs/
├── scripts/
├── Dockerfile
├── Makefile
└── README.md
```

---

## Commands

### Version

```bash
cloudctl version
```

Output

```text
cloudctl v0.1.0
```

---

### AWS

Create a resource

```bash
cloudctl aws create demo
```

List resources

```bash
cloudctl aws list
```

Delete a resource

```bash
cloudctl aws delete demo
```

---

### Kubernetes

List Pods

```bash
cloudctl k8s pods
```

List Deployments

```bash
cloudctl k8s deployments
```

Check Rollout Status

```bash
cloudctl k8s rollout api-service
```

---

## Technologies

- Go
- Cobra
- Viper
- Zap
- Docker
- AWS SDK (planned)
- Kubernetes client-go (planned)
- Testify

---

## Running Locally

Clone the repository

```bash
git clone https://github.com/Chetana-Thorat/cloud-resource-manager.git

cd cloud-resource-manager
```

Install dependencies

```bash
go mod tidy
```

Run

```bash
go run . version
```

Run tests

```bash
go test ./...
```

---

## Roadmap

- [x] Cobra CLI
- [x] Mock AWS provider
- [x] Mock Kubernetes provider
- [x] Structured logging
- [x] Configuration management
- [x] Unit tests
- [ ] AWS SDK integration
- [ ] Kubernetes client-go integration
- [ ] Docker image publishing
- [ ] GitHub Actions CI

---

## License

MIT
