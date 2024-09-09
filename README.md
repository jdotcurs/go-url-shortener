# URL Shortener

This project demonstrates a simple URL shortener service using Go for the backend API and a basic HTML/JavaScript frontend. It's designed to showcase fundamental skills in development with Go. It is built with a Test Driven Development (TDD) approach.

## Features

- Shorten long URLs
- Redirect short URLs to their original long URLs
- In-memory storage (easily extendable to use a database)
- Concurrent-safe operations
- Simple HTML/JavaScript frontend
- CORS support for local development

## Technologies Used

- Go (Golang)
- HTML/CSS/JavaScript (Vanilla)

## Project Structure

- `cmd/urlshortener/main.go`: Contains the Go backend code
- `internal/`: Contains the core logic of the application
- `frontend/index.html`: Frontend interface

## Setup and Running

1. Ensure you have Go installed on your system.

2. Build the project:
   ```
   go build -o urlshortener ./cmd/urlshortener
   ```

3. Run the server:
   ```
   ./urlshortener
   ```

4. Open `frontend/index.html` in a web browser.

## API Endpoints

- POST `/shorten`: Create a new short URL
- GET `/{shortURL}`: Redirect to the original long URL

## What This Project Demonstrates

### Go Backend Development:
- Creating a RESTful API using the standard library
- Handling HTTP requests and responses
- JSON encoding/decoding
- Error handling and appropriate HTTP status codes

### Code Organization:
- Structuring a Go application
- Separating concerns (storage, HTTP handling, etc.)

### Frontend Integration:
- Simple HTML/CSS for user interface
- Vanilla JavaScript for API interactions
- Asynchronous operations with Fetch API

### API Design:
- RESTful principles
- CORS handling for cross-origin requests

### Testing:
- Unit tests for different components
- Test coverage reporting

## Why This Project Was Created

This project serves as a portfolio piece, demonstrating:

- Ability to create a full-stack application
- Understanding of web development concepts
- RESTful API development
- Frontend and backend communication
- Code organization and best practices in Go
- Test Driven Development (TDD) approach
- Code coverage analysis

It's designed to be simple enough to understand quickly, yet comprehensive enough to showcase a range of Go skills.

## Testing

To run the tests for this project, use the following command:

```
go test ./...
```

For more verbose output, you can use the `-v` flag:

```
go test -v ./...
```

To check the coverage of the tests, you can use the following command:

```
go test -cover ./...
```

For a detailed HTML coverage report:

```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Future Improvements

- Implement persistent storage (e.g., PostgreSQL, Redis)
- Implement custom short URLs
- Containerize the application using Docker
- Enhance error handling and input validation
- Use a more robust router for better URL handling