# taas

## Description

This is a Go application built with the Gin framework and GORM. It provides a REST API for managing tags, entities, tag mappings, and tenants.

## Features

*   Authentication using JWT
*   Logging using Logrus
*   PostgreSQL database support
*   CRUD operations for tags, entities, tag mappings, and tenants
*   Caching with Memcached

## Configuration

The application is configured using environment variables. You can create a `.env` file in the root of the project and the application will load the variables from it. See `.env.example` for a list of all available options.

| Environment Variable     | Default Value                      | Description                                   |
| ------------------------ | ---------------------------------- | --------------------------------------------- |
| `SERVER_PORT`            | `8080`                             | Port for the HTTP server                      |
| `GIN_MODE`               | `debug`                            | Gin framework mode (`debug` or `release`)     |
| `SERVER_READ_TIMEOUT`    | `10s`                              | Server read timeout                           |
| `SERVER_WRITE_TIMEOUT`   | `10s`                              | Server write timeout                          |
| `SERVER_IDLE_TIMEOUT`    | `60s`                              | Server idle timeout                           |
| `DB_HOST`                | `localhost`                        | Database host                                 |
| `DB_PORT`                | `5432`                             | Database port                                 |
| `DB_USER`                | `taasuser`                         | Database username                             |
| `DB_PASSWORD`            | `taasuser`                         | Database password                             |
| `DB_NAME`                | `taas`                             | Database name                                 |
| `DB_SSL_MODE`            | `disable`                          | Database SSL mode                             |
| `MEMCACHED_HOST`         | `127.0.0.1`                        | Memcached host                                |
| `MEMCACHED_PORT`         | `11211`                            | Memcached port                                |
| `LOG_LEVEL`              | `info`                             | Log level (`debug`, `info`, `warn`, `error`)  |
| `LOG_FORMAT`             | `json`                             | Log format (`json` or `text`)                 |
| `JWT_SECRET`             | `wN3vP8qjR9tL5kZxT2sA7yH0uE6fV4dG` | Secret key for signing JWTs                   |

## Setup

### Prerequisites

*   Docker and Docker Compose

### Installation with Docker

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd taas
    ```

2.  **Create a `.env` file:**

    Copy the example environment file and update it with your configuration if needed.

    ```bash
    cp .env.example .env
    ```

3.  **Build and run the services:**

    ```bash
    docker-compose up --build
    ```

    This will start the application, a PostgreSQL database, and a Memcached instance. The application will be available at `http://localhost:8080`.

4.  **Run database migrations:**

    The migrations are not run automatically. You can run them with the following command:

    ```bash
    docker-compose exec app go run main.go migrate
    ```

## API Endpoints

All endpoints are prefixed with `/api`.

### Tags

*   `GET /tags`: List all tags.
*   `POST /tags`: Create a new tag.
*   `GET /tags/:id`: Get a tag by its ID.
*   `PUT /tags/:id`: Update a tag.
*   `DELETE /tags/:id`: Delete a tag.
*   `GET /tags/search`: Search for tags by name.

### Entities

*   `GET /entities`: List all entities.
*   `POST /entities`: Create a new entity.
*   `DELETE /entities/:type`: Delete an entity by its type.

### Tag Mappings

*   `GET /:entity_type/tag_mappings/:id`: List tag mappings for an entity.
*   `POST /:entity_type/tag_mappings/:id`: Create tag mappings for an entity.
*   `DELETE /:entity_type/tag_mappings/:id`: Delete tag mappings for an entity.

### Tenants

*   `POST /tenants`: Create a new tenant.
*   `GET /tenants/:id`: Get a tenant by its ID.
*   `DELETE /tenants/:id`: Delete a tenant.

## Usage Examples

Replace `<jwt_token>` with a valid JWT token.

### Create a Tenant

```bash
curl -X POST http://localhost:8080/api/tenants -H "Content-Type: application/json" -d '{"name": "My Tenant", "api_key": "my-secret-key"}'
```

### Create a Tag

```bash
curl -X POST http://localhost:8080/api/tags -H "Content-Type: application/json" -H "Authorization: Bearer <jwt_token>" -d '{"name": "example-tag"}'
```

### Get a Tag by ID

```bash
curl -X GET http://localhost:8080/api/tags/1 -H "Authorization: Bearer <jwt_token>"
```

## Contributing

Contributions are welcome! Please open a pull request with your changes.

## License

[MIT](LICENSE)
