# Project Title

## Introduction
This is a Go project that serves as a book management system. It uses a clean architecture pattern, separating the project into `domain`, `usecase`, `repository`, and `delivery` layers. This structure allows for clear separation of concerns and makes the codebase easier to test and maintain.

## Getting Started
To start the server, navigate to the project directory and run the following command:

```sh
docker-compose up
then call api with port 8081
```

## API Endpoints
The project exposes the following RESTful API endpoints:

- `GET /books`: Fetch all books
- `GET /books/{id}`: Fetch a book by its ID
- `POST /books`: Create a new book
- `PUT /books/{id}`: Update a book by its ID
- `DELETE /books/{id}`: Delete a book by its ID

## Concurrency Handling
Concurrency in this project is handled using Go's built-in goroutines and channels. This allows the server to handle multiple requests simultaneously, improving the overall performance and responsiveness of the API. The implementation of goroutines can be found in the [`timeout.go`](delivery/middleware/timeout.go) file in the `middleware` package and in the [`book_repo.go`](repository/book_repo.go) file in the `repository` package.

## Error Handling
Errors are handled using a custom `apperror` package. This package defines a custom `Error` type that includes an error `Type` and a `Message`. The `Type` is a string that represents the kind of error (e.g., "AUTHORIZATION", "BADREQUEST", "CONFLICT", etc.), and the `Message` is a string that provides more detail about the error. The `apperror` package also provides several "factory" functions for creating new instances of these custom errors.

## Proxy and IP Filtering
The project uses Traefik as a reverse proxy. Traefik is configured to only route requests to the application that originate from a specific IP range. This IP range can be configured in the `docker-compose.yml` file.