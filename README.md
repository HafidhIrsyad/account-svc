# Account Service (account-svc)

A Go-based microservice providing RESTful APIs for account registration, deposits, withdrawals, and balance inquiries. Built with Echo, GORM, PostgreSQL, and containerized with Docker Compose.

## Features

- RESTful API with Go (Echo framework)
- PostgreSQL database accessed via GORM
- Configuration management with Viper
- Structured, leveled logging via Zerolog
- UUID-based identifiers for accounts
- Containerized with Docker & Docker Compose
- Task automation via Makefile

## Table of Contents

- [Features](#features)
- [API Endpoints](#api-endpoints)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Local Development](#local-development)
  - [Docker Setup](#docker-setup)
- [Available Make Commands](#available-make-commands)
- [License](#license)

## API Endpoints

### POST `/account/daftar`

Register a new account.
**Request Body**:

```json
{
  "nama": "John Doe",
  "nik": "1234567890123456",
  "no_hp": "081234567890"
}
```

### POST `/account/tabung`

Deposit into an account.
**Request Body**:

```json
{
  "no_rekening": "1000001",
  "nominal": 500000
}
```

### POST `/account/tarik`

Withdraw from an account.
**Request Body**:

```json
{
  "no_rekening": "1000001",
  "nominal": 200000
}
```

### GET `/account/saldo/{no_rekening}`

Get current balance for an account.

## Getting Started

### Prerequisites

- Go 1.23.x
- Docker & Docker Compose
- Git

### Local Development

1. **Clone the repo**

   ```bash
   git clone https://github.com/HafidhIrsyad/account-svc.git
   cd account-svc
   ```

2. **Copy & customize environment file**

   ```bash
   cp .env.example .env
   ```

   Update the following values in `.env`:

   ```dotenv
   APP_PORT=8080
   DRIVER_NAME=postgres
   POSTGRES_HOST=db
   POSTGRES_USER=
   POSTGRES_PASSWORD=
   POSTGRES_DB_NAME=
   POSTGRES_PORT=
   POSTGRES_DB_MAX_OPEN_CONNECTION=25
   POSTGRES_DB_MAX_IDLE_CONNECTION=10
   DB_CONNECTION_MAX_LIFE_TIME=300
   ```

3. **Fetch dependencies**

   ```bash
   go mod download
   ```

4. **Run the service**
   ```bash
   make run
   ```
   The API will be available at `http://localhost:8080`.

### Docker Setup

Bring up the service and PostgreSQL in containers:

```bash
make docker-up
```

Tear down:

```bash
make docker-down
```

View logs:

```bash
make docker-logs
```

## Available Make Commands

- `make run`   Run service locally (`go run main.go`)
- `make build`  Build the Go binary
- `make docker-up` Start Docker Compose stacks
- `make docker-down` Stop & remove containers
- `make docker-logs` Tail service & DB logs
