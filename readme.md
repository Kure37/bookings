# Bookings and Reservations

This is a web application for bookings and reservations.

- Built in Go version go1.25.5
- Uses the chi router
- Uses alex edwards SCS session management
- Uses nosurf for CSRF protection

## Getting Started

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Run `go run cmd/web/main.go` to start the server
4. Open `http://localhost:8080` in your browser