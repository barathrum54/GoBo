# GoBo

GoBo is a modular and scalable backend boilerplate written in Go. It leverages modern tools such as the Fiber framework, GORM ORM, Zap logging, and Swagger for high-performance API development and extensibility.

---

## 🚀 Features

- **Fiber Framework**: A fast and flexible HTTP server.
- **GORM**: Database ORM support for easy modeling and migrations.
- **Zap Logging**: High-performance, configurable logging.
- **Swagger Integration**: Auto-generated API documentation with Swagger UI.
- **Modular Architecture**: Extensible API design for scalability.
- **High Code Quality**: Integrated with `golangci-lint` for linting and static analysis.
- **Testing Support**: Structured testing setup using `testify`.
- **Basic Authentication Middleware**: Protect specific routes with simple Basic Authentication.
- **Rate Limiting Middleware**: Protect routes from abuse by limiting request rates.

---

## 🛠️ Installation and Setup

### 1. **Clone the Repository**

```bash
git clone https://github.com/username/gobo.git
cd gobo
```

### 2. **Install Dependencies**

```bash
go mod tidy
```

### 3. **Create the .env File**

Create a `.env` file with the following environment variables:

```plaintext
DATABASE_URL=postgres://username:password@localhost:5432/dbname
REDIS_URL=localhost:6379
```

### 4. **Generate Swagger Documentation**

Install `swag` if not already installed:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate Swagger documentation:

```bash
swag init
```

### 5. **Run Database Migrations**

Migrations will run automatically when the project starts, creating necessary tables.

### 6. **Start the Server**

```bash
go run cmd/main.go
```

The server will be accessible at `http://localhost:3000`.

### 7. **Access Swagger UI**

Visit `http://localhost:3000/swagger/index.html` to explore the Swagger UI for API documentation.

---

## 📂 Project Structure

```
gobo/
├── cmd/                # Entry point for the HTTP server
├── docs/               # Swagger documentation files
├── internal/
│   ├── app/           # Fiber app initialization and configuration
│   ├── cache/         # Redis connection and helper functions
│   ├── db/            # Database connection and setup
│   ├── logger/        # Zap logger configuration
│   ├── middleware/    # Middleware for request handling
│   ├── models/        # GORM models
│   ├── routes/        # API routes
│   ├── testhelpers/   # Utilities for testing
├── .env               # Environment variables
├── .golangci-lint.yaml # Linter configuration
├── go.mod             # Go module definition
├── go.sum             # Module dependencies
├── main.go            # Application entry point
├── README.md          # Project documentation
```

---

## 📋 Technologies Used

- [Go](https://go.dev/) - Programming Language
- [Fiber](https://gofiber.io/) - HTTP Framework
- [GORM](https://gorm.io/) - ORM Library
- [Zap](https://github.com/uber-go/zap) - Logging Library
- [Redis](https://redis.io/) - Caching
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Swaggo](https://github.com/swaggo/swag) - Swagger Documentation
- [GolangCI-Lint](https://golangci-lint.run/) - Code Analysis and Linter

---

## ✅ Testing

### Run Tests

To execute the test suite:

```bash
go test ./... -v
```

The tests will reset the database, create new tables, and validate CRUD operations.

---

## 🔧 Linter

To run static code analysis and linter checks:

```bash
golangci-lint run
```

---

## 🔧 Redis Cache

The project includes Redis caching support, managed within the `internal/cache` module and available for use in API routes.

### Example Usage:

```go
import "gobo/internal/cache"

// Save data to Redis
cache.Set("key", "value", 60*time.Second)

// Retrieve data from Redis
value, err := cache.Get("key")
if err != nil {
    log.Println("Cache miss")
} else {
    log.Printf("Cache hit: %s", value)
}
```

---

## 🔥 Middleware

### Basic Authentication Middleware

The project includes a Basic Authentication middleware located in the `internal/middleware` directory. This middleware can be used to protect specific routes with simple Basic Authentication.

### Example Usage:

```go
import "gobo/internal/middleware"

func Register(app *fiber.App) {
    app.Get("/public", func(c *fiber.Ctx) error {
        return c.SendString("This is a public route")
    })

    protected := app.Group("/protected", middleware.BasicAuthMiddleware("admin", "password"))

    protected.Get("/secure", func(c *fiber.Ctx) error {
        return c.SendString("You are authorized")
    })
}
```

---

### Rate Limiting Middleware

The project includes a Rate Limiting middleware located in the `internal/middleware` directory. This middleware can be used to protect routes by limiting the number of requests within a specified time frame.

### Example Usage:

```go
import "gobo/internal/middleware"

func Register(app *fiber.App) {
    limited := app.Group("/limited", middleware.RateLimitMiddleware(5, 10*time.Second))

    limited.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("This route is rate limited")
    })
}
```

---

## 🔥 Logging

The project uses **Zap** for high-performance and configurable logging. The logging setup is located in the `internal/logger` directory.

### Example Usage:

```go
import "gobo/internal/logger"

func Example() {
    logger.Log.Info("Example log message", zap.String("key", "value"))
}
```

---

## 🔧 Swagger Integration

The project uses **Swaggo** for generating Swagger API documentation. The documentation is served at `/swagger/index.html`.
Fiber with **Swaggo** requires named functions for routes to work properly.

### Example Annotation:

```go
// @Summary      Example Endpoint
// @Description  An example endpoint.
// @Tags         examples
// @Accept       json
// @Produce      json
// @Success      200 {object} ExampleResponse
// @Failure      400 {object} ErrorResponse
// @Router       /example [get]
```

To add Swagger documentation, annotate your handlers with appropriate tags as shown above. Regenerate the docs with:

```bash
swag init
```

---

## 🤝 Contributing

1. Fork the repository.
2. Create a new branch: `git checkout -b my-new-feature`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push the branch: `git push origin my-new-feature`.
5. Open a Pull Request.

---
