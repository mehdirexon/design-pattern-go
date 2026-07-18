# Breeders

A learning project that implements a Go web application around pet breeds and design patterns. This repository is the completed version of the main course project, and it uses the remote service from [tsawler/go-breeders-remote](https://github.com/tsawler/go-breeders-remote) for the cat breed data.

This project is built as part of the course:
Working with Design Patterns in Go (Golang)  
https://www.udemy.com/course/working-with-design-patterns-in-go-golang/

## Features

- HTTP server built with `net/http` and `chi`
- Server-side rendered templates with Go `html/template`
- JSON API endpoints for pets and breeds
- MariaDB persistence with `database/sql`
- Docker Compose setup for local database development
- Factory, abstract factory, and builder pattern examples
- Static assets served from the `static` directory

## Tech Stack

- Go
- `github.com/go-chi/chi/v5`
- `github.com/go-sql-driver/mysql`
- `github.com/tsawler/toolbox`
- MariaDB
- Docker Compose

## Project Structure

- `cmd/web`: application entrypoint, handlers, routes, DB setup, and rendering helpers
- `configuration`: application wiring and service registration
- `adapters`: remote service adapters for cat breed data
- `models`: database repositories and breed models
- `pets`: factory pattern examples and pet builders
- `templates`: Go HTML templates
- `static`: static assets
- `sql`: database schema and seed scripts
- `db-data`: persisted MariaDB data used by the local container

## Prerequisites

- Go installed
- Docker and Docker Compose installed
- A local MariaDB instance, or the included Compose setup

## Environment

The application uses a MySQL/MariaDB DSN. The default values in `cmd/web/main.go` point to the local Compose service:

```text
mariadb:myverysecretpassword@tcp(localhost:12345)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s
```

If you want to override the database connection, pass a custom DSN when starting the app.

## Database

Start MariaDB with Docker Compose:

```zsh
docker compose up -d
```

The container initializes from `sql/breeders_mysql.sql` and persists data in `db-data/mariadb`.

## Run

From the project root:

```zsh
go run ./cmd/web
```

The server starts on `http://localhost:4000`.

## Routes

- `/` - home page
- `/{page}` - renders a template page by name
- `/test-patterns` - test page for the design pattern demos
- `/api/dog-from-factory` - returns a dog built with the factory pattern
- `/api/cat-from-factory` - returns a cat built with the factory pattern
- `/api/dog-from-abstract-factory` - returns a dog from the abstract factory
- `/api/cat-from-abstract-factory` - returns a cat from the abstract factory
- `/api/dog-from-builder` - returns a dog built with the builder pattern
- `/api/cat-from-builder` - returns a cat built with the builder pattern
- `/api/dog-breeds` - returns all dog breeds as JSON
- `/api/cat-breeds` - returns all cat breeds as JSON
- `/api/animal-from-abstract-factory/{species}/{breed}` - returns a pet for the requested species and breed

## Testing

Run all tests:

```zsh
go test ./...
go test ./... -cover
```

## Notes

- The app uses `configuration.New` to wire the database and cat breed adapter into the application instance.
- `adapters.RemoteService` wraps the remote cat breed service from `tsawler/go-breeders-remote`.
- The repository layer is intentionally simple and focused on the patterns shown in the course.

## License

For learning and educational purposes.
