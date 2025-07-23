# integration-suspect-service

Suspect service written in Go (Golang) following the Clean Architecture approach.

## 🔧 Project Structure

```
.
│   .env
│   go.mod
│   go.sum
│   README.md
│
├───.vscode
│       launch.json
│
├───app
│       main.go
│
├───configs
│       config.go
│
├───docs
│       docs.go
│       swagger.json
│       swagger.yaml
│
├───modules
│   ├───middlewares
│   │       jwt.go
│   │       loggers.go
│   │       recover.go
│   │
│   ├───servers
│   │       handler.go
│   │       server.go
│   │
│   └───suspect
│       ├───controllers
│       │       suspect_controllers.go
│       │
│       ├───entities
│       │       suspect_entities.go
│       │       suspect_mappings.go
│       │
│       ├───repositories
│       │       suspect_repositories.go
│       │
│       └───usecases
│               suspect_usecases.go
│
└───pkg
    ├───clients
    │   └───ktb
    │           ktb_clients.go
    │
    ├───databases
    │       postgresql.go
    │
    ├───loggers
    │       loggers.go
    │       loggers_entities.go
    │       loggers_resty.go
    │
    ├───resty_factory
    │       resty_factory.go
    │
    ├───utils
    │       array.go
    │       connection_url_builder.go
    │       hash.go
    │       time.go
    │
    └───validators
            custom_validators.go
```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/parnupong-geniussoft/integration-suspect-service.git
cd integration-suspect-service
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Project

```bash
go run ./app/main.go
```

## 📌 Features

- ✅ Submit Suspect list to Ktb

## 🛠 Technologies Used

- [Go Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Swagger (Swaggo)](https://github.com/swaggo/swag)

## 📄 API Usage

### POST /v1/integration-api/request_token

## 📚 API Documentation

Swagger UI available at:  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## 🧹 Clean Architecture Overview

- `controllers` → Handles requests and invokes usecases
- `usecases` → Business logic implementation
- `repositories` → Handles data persistence
- `entities` → Data structures

## 📁 Middleware

- `middlewares/jwt.go` → JWT Token verification
- `middlewares/loggers.go` → Logs request data into database
- `middlewares/recover.go` → Handles panic and returns safe response
