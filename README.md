# taas

## Description

This is a Go application built with the Gin framework and GORM. It provides a REST API for managing tags, entities, tag mappings, and tenants.

## Features

*   Authentication using JWT
*   Logging using Logrus
*   PostgreSQL database support
*   CRUD operations for tags, entities, tag mappings, and tenants

## Setup

### Prerequisites

*   Go 1.24 or higher
*   PostgreSQL database

### Installation

1.  Clone the repository:

    ```bash
    git clone <repository_url>
    ```

2.  Navigate to the project directory:

    ```bash
    cd taas
    ```

3.  Install dependencies:

    ```bash
    go mod download
    ```

4.  Configure the application:

    Create a `config.yaml` file with the following contents:

    ```yaml
    server:
      port: "8080"
      read_timeout: 10s
      write_timeout: 10s
      idle_timeout: 30s
    database:
      host: "localhost"
      port: "5432"
      user: "postgres"
      password: "password"
      dbname: "taas"
    jwt_secret: "secret"
    ```

    Replace the placeholder values with your actual configuration.

5.  Run the database migrations:

    ```bash
    go run main.go migrate
    ```

6.  Start the server:

    ```bash
    go run main.go
    ```

## Usage

The API endpoints are documented in the `router/router.go` file. You can use a tool like curl or Postman to interact with the API.

### Examples

#### Create a tag

```bash
curl -X POST http://localhost:8080/api/tags -H "Content-Type: application/json" -H "Authorization: Bearer <jwt_token>" -d '{"name": "example"}'
```

#### Get a tag by ID

```bash
curl -X GET http://localhost:8080/api/tags/1 -H "Authorization: Bearer <jwt_token>"
```

## Contributing

Contributions are welcome! Please open a pull request with your changes.

## License

[MIT](LICENSE)
