# Go API Starter

A modern, production-ready Go API backend built with clean architecture principles and modular design. This project serves as a comprehensive starter template for building scalable web APIs with authentication, file storage, and business logic modules.

## Overview

**Go API Starter** is a solid foundation for building high-performance, scalable web APIs using **Go**. It follows **Clean Architecture** and **Modular Design** principles to ensure code maintainability, testability, and scalability.

This template comes pre-configured with essential components like Authentication, Database (PostgreSQL), Caching (Redis), and a robust CLI, allowing developers to focus on business logic immediately.

## Key Features

- **Modular Architecture**: Business logic is separated into independent modules (e.g., `auth`, `workers`) with clear boundaries.
- **Clean Architecture**: Strict separation of concerns (Entities, Repositories, Services, Handlers).
- **Authentication**: Built-in user registration, login, and management flow.
- **Dependency Injection**: Uses `samber/do` for lightweight and explicit dependency management.
- **Database & Caching**: Pre-configured PostgreSQL (pgx) and Redis client.
- **CLI Support**: Integrated `spf13/cobra` for running administrative tasks and background workers.
- **Observability**: Structured logging with `rs/zerolog`.
- **Utilities**: Helper functions for email, OTP, validation, passwords, and more.

## Tech Stack

- **Language**: Go 1.24+
- **Framework**: [Echo v4](https://echo.labstack.com/)
- **Database**: PostgreSQL (Driver: `pgx/v5`)
- **Caching**: Redis (`go-redis`)
- **Dependency Injection**: [samber/do](https://github.com/samber/do)
- **Configuration**: [spf13/viper](https://github.com/spf13/viper)
- **CLI**: [spf13/cobra](https://github.com/spf13/cobra)
- **Logging**: [rs/zerolog](https://github.com/rs/zerolog)

## Project Structure

```
/
├── cmd/            # Entry point (Main application & CLI)
├── modules/        # Domain modules (Auth, Workers, etc.)
│   ├── auth/       # Authentication module
│   │   ├── dto/        # Data Transfer Objects
│   │   ├── entity/     # Database models & Domain entities
│   │   ├── handler/    # HTTP/gRPC handlers (Controllers)
│   │   ├── repository/ # Data access layer (DB operations)
│   │   ├── service/    # Business logic layer
│   │   └── router/     # Route definitions
│   └── ...
├── pkg/            # Shared packages & Infrastructure
│   ├── config/     # Configuration loader
│   ├── database/   # Database connection
│   ├── logger/     # Logger setup
│   ├── utils/      # Utility functions
│   └── ...
├── Dockerfile      # Docker build configuration
└── docker-compose.yml
```

## Todo

- [x] Use [do](https://github.com/samber/do) to manage dependency injection
- [ ] Add an example using Kafka to publish and consume messages for processing workers
- [ ] Add [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to expose gRPC services over HTTP/JSON
