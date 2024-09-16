# github.com/Sabyradinov/go-dating-matchmaking

Golang project with hexagonal architecture

## Hexagonal architecture

- In this project implemented hexagonal architecture structure with gin web server [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

### Pros of hexagonal architecture:

- Maintainability, low level of dependency between project layers, which gives an advantage when adding and changing functionality.
- Possibility for several developers to work in parallel in different parts of the code.
- Testability, ability to write unit tests for each component of the project without dependencies on other components.
- Adaptability to different task requirements.

### Project structure:

- In hexagonal architecture there are two main concepts - ports and adapters, conditionally for understanding: in ports (contracts) interfaces are declared, in adapters interfaces are implemented (more details here [https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))).
- Also, before starting development, study the basic concepts of Domain Driven Design (DDD).
- `/common` package contains common functions, constants, errors, etc.
- `/config` package contains configuration files, environment variables, etc.
- `/internal/domain` package contains domain services, entities, value objects, etc.
- `/internal/domain/port` package contains interfaces for domain services.
- `/internal/domain/service` package contains domain services.
- `/internal/handler` package contains handlers for web server.
- `/internal/http` package contains http server, REST API.
- `/postgres_bd_scripts` package contains SQL scripts for creating tables, indexes, etc.`

## Setup Instructions

### Prerequisites

- Go 1.16 or later
- PostgreSQL
- Git

### Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/Sabyradinov/go-dating-matchmaking.git
    cd go-dating-matchmaking
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Set up the database:**

    - Create a PostgreSQL database.
    - Update the database configuration in the `config` package.
    - Run the database migrations:

    ```sh
    go run cmd/migrate/main.go
    ```

4. **Run the application:**

    ```sh
    go run cmd/server/main.go
    ```

5. **Update swagger documentation:**

    ```sh
    swag init -g ./cmd/server/main.go -o cmd/docs
    ```

### Running Tests

To run the tests, use the following command:

```sh
go test ./...