Backend for [wflow-front](https://github.com/rogerok/wflow-backend)

### Prerequisites

Before setting up the development environment, ensure you have the following tools installed:

| Tool           | Version | Purpose                       |
|----------------|---------|-------------------------------|
| Go             | 1.19+   | Application runtime           |
| Docker         | Latest  | Database containerization     |
| Docker Compose | Latest  | Multi-container orchestration |
| migrate CLI    | Latest  | Database migrations           |
| swag           | Latest  | API documentation generation  |
| air            | Latest  | Live reloading (optional)     |

### Environment Setup

1. **Clone and Configure Environment**
    ```
    git clone https://github.com/rogerok/wflow-backend
    cd wflow-backend
    ```
    Create a .env file in the project root with the following configuration:
    ```
    # Database Configuration
    POSTGRES_USER=your_user
    POSTGRES_PASSWORD=your_password
    POSTGRES_DB_NAME=your_db_name
    DB_PORT=5432
    HOST=localhost
    SSL_MODE=disable
    
    # Application Configuration
    PORT=5000
    
    # JWT Configuration (add your secrets)
    JWT_SECRET=your_jwt_secret
    REFRESH_SECRET=your_refresh_secret
    ```

2. **Database Setup**
    ```
    make db-up
    ```
    Run database migrations:
    ```
   make migrate-up
   ```

### Build and Run Commands

The `Makefile` provides several commands for different development scenarios:

| Command      | Purpose           | Description                          |
|--------------|-------------------|--------------------------------------|
| `make build` | Compile           | Builds binary to `bin/api`           |
| `make run`   | Execute           | Builds and runs the application      |
| `make air`   | Live Reload       | Starts with automatic recompilation  |
| `make dev`   | Full Development  | Builds app and starts database       |
| `make test`  | Testing           | Runs all Go tests                    |

### Live Development Mode
For active development with automatic reloading:

```
# Terminal 1: Start database
make db-up

# Terminal 2: Start application with live reload
make air
```

### API Documentation
The application automatically generates and serves API documentation using Swagger:

- Documentation Generation: make swagger-docs
- Interactive Documentation: Available at http://localhost:5000/swagger/
- Source Annotations: Swagger comments in Go source files

### Migration Management
```
# Create new migration
make dbseq name=your_migration_name

# Apply migrations
make migrate-up

# Rollback last migration
make migrate-down
```

[DeepWiki](https://deepwiki.com/rogerok/wflow-backend)

