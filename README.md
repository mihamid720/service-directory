# Service Directory

A scalable REST API service for managing service information and discovery, built with Go, Gin, and PostgreSQL.

## Features

- List services with pagination and sorting
- Get detailed information about a specific service
- Search services by name or description
- PostgreSQL database for persistence
- Docker support for easy database setup
- Integration tests for API endpoints

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose (for PostgreSQL)

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd services-api
```

2. Start the PostgreSQL database:
```bash
docker-compose up -d db
```

3. Run the service:
```bash
DB_HOST=localhost \
DB_USER=postgres \
DB_PASSWORD=postgres \
DB_NAME=services_db \
DB_PORT=5432 \
go run cmd/api/main.go
```

The service will automatically:
- Connect to the database
- Create necessary tables
- Seed initial data

## API Endpoints

### GET /services
Lists all services with pagination and search support.

Query parameters:
- `page`: Page number (default: 1)
- `page_size`: Number of items per page (default: 10, max: 100)
- `search`: Search term for filtering by name or description
- `sort_by`: Sort by field (options: "name", "-name")

Example response:
```json
{
    "total": 5,
    "page": 1,
    "page_size": 10,
    "services": [
        {
            "id": 1,
            "name": "Authentication Service",
            "description": "Handles user authentication and authorization using JWT tokens",
            "versions": 2,
            "created_at": "2024-03-13T01:07:39.401Z",
            "updated_at": "2024-03-13T01:07:39.401Z"
        }
        // ... more services
    ]
}
```

### GET /services/:id
Get detailed information about a specific service.

Example response:
```json
{
    "id": 1,
    "name": "Authentication Service",
    "description": "Handles user authentication and authorization using JWT tokens",
    "versions": 2,
    "created_at": "2024-03-13T01:07:39.401Z",
    "updated_at": "2024-03-13T01:07:39.401Z"
}
```

## Running Tests

1. Start the test database:
```bash
docker-compose up -d test_db
```

2. Run the tests:
```bash
go test ./... -v
```

The test suite includes integration tests that verify:
- Listing services with pagination
- Getting a specific service by ID
- Handling non-existent services
- Database seeding and cleanup

## Authentication & Authorization

The service can be extended with JWT-based authentication and role-based access control (RBAC). Implementation would require creating several new files and updating existing ones: a new `internal/models/user.go` file would define the user model and authentication structures; `internal/middleware/auth.go` would handle JWT validation and role-based access control; `internal/handlers/auth_handler.go` would implement login and registration endpoints; and `internal/database/database.go` would need user-related database operations. The main router in `cmd/api/main.go` would need to be updated to include authentication middleware and new auth routes. Implementation would include secure password handling with bcrypt, token refresh mechanisms, and proper security headers for API protection, all configured through environment variables defined in the project root.

## CRUD Operations

The service architecture supports full Create, Read, Update, and Delete (CRUD) operations for managing services. Implementation would require updates to several key files: the `internal/handlers/service_handler.go` file would need new handler functions for create, update, and delete operations; the `internal/models/service.go` file would need input validation structs and additional model methods; and the routing configuration in `cmd/api/main.go` would need new endpoint registrations. These operations would be implemented using RESTful principles, with proper request validation, error handling, and database transaction management. The existing pagination and search capabilities in `internal/database/database.go` would seamlessly integrate with these new endpoints, maintaining the service's scalability and performance characteristics.

## Development History

The project was developed through the following logical commits:

1. **Initial Project Setup**
   - Initialize Go module
   - Create basic project structure

2. **Database Layer Implementation**
   - Add database connection configuration
   - Create Service model
   - Implement database operations
   - Add environment variable handling

3. **Basic API Implementation**
   - Set up Gin router
   - Implement GET /services endpoint
   - Add pagination support
   - Basic error handling

4. **Service Details Endpoint**
   - Add GET /services/:id endpoint
   - Implement detailed service view

5. **Search and Sort Functionality**
   - Add search capability
   - Implement sorting

6. **Testing Infrastructure**
   - Set up test database configuration
   - Create test helpers
   - Add integration tests
   - Implement database seeding

7. **Documentation and Cleanup**
   - Complete API documentation
   - Add usage examples
   - Code cleanup and formatting
   - Update README
