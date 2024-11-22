# Structure Project
```
E-commerce/
├── cmd/
│   └── main.go                     # Application entry point
│
├── internal/                       # Private application code
│   ├── db/
│   │   └── migrations/
│   │       └── query.sql
│   │
│   ├── delivery/                   # HTTP handlers
│   │   └── handler.go
│   │
│   ├── entity/                     # Domain models
│   │   └── product.go
│   │   └── cart.go
│   │   └── cart_item.go
│   │   └── order.go
│   │   └── order_item.go
│   │
│   ├── middleware/                 # HTTP middleware
│   │   └── middleware.go
│   │
│   ├── repository/                 # Data access layer
│   │   ├── cart_repository.go
│   │   ├── order_repository.go
│   │   └── product_repository.go
│   │
│   └── usecase/                    # Business logic layer
│       ├── cart_usecase.go
│       ├── order_usecase.go
│       └── product_usecase.go
│
├── pkg/                            # Public shared packages
│   ├── constant/
│   ├── error/
│   └── logger/
│
├── .gitignore
├── go.mod                          # Go module definition
├── go.sum                          # Go module checksums
└── Makefile                        # Build and development commands
```


1. `cmd/` - Contains the main application entry point
2. `internal/` - Private application code
    - `db/` - Database migrations and configuration
    - `delivery/` - HTTP handlers and routing
    - `entity/` - Domain models/entities
    - `middleware/` - HTTP middleware components
    - `repository/` - Data access layer implementations
    - `usecase/` - Business logic implementations
3. `pkg/` - Shared packages that could potentially be used by other applications
4. Root level configuration files

The project follows a layered architecture:
- Presentation Layer (delivery)
- Business Logic Layer (usecase)
- Data Access Layer (repository)
- Domain Layer (entity)

This structure makes the code modular, maintainable, and follows Go best practices for project organization.