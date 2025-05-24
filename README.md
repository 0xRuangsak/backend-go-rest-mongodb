# User Management API - Go Backend

A RESTful API built with Go, MongoDB, and JWT authentication for managing users.

## 🚧 Work in Progress

This project is being developed as part of a backend coding challenge.

## Features (Planned)
- [ ] User CRUD operations
- [ ] JWT authentication
- [ ] MongoDB integration
- [ ] Middleware (logging, auth)
- [ ] Background tasks with goroutines
- [ ] Unit tests
- [ ] Docker support
- [ ] gRPC implementation (bonus)

## Tech Stack
- **Language:** Go
- **Database:** MongoDB
- **Authentication:** JWT
- **Architecture:** Clean/Hexagonal

## Development Progress
Check commit history to see development phases.

## Getting Started
```bash
# Clone the repository
git clone https://github.com/yourusername/backend-go-rest-mongodb.git
cd backend-go-rest-mongodb

# Install dependencies
go mod tidy

# Run the application
go run cmd/server/main.go
```

## API Endpoints (Coming Soon)
- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login user
- `GET /users` - Get all users (protected)
- `GET /users/:id` - Get user by ID (protected)
- `PUT /users/:id` - Update user (protected)
- `DELETE /users/:id` - Delete user (protected)

## Environment Variables
Create a `.env` file with:
```
MONGODB_URI=mongodb://localhost:27017/userdb
JWT_SECRET=your-secret-key
PORT=8080
```

## Testing
```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Docker
```bash
# Build and run with Docker Compose
docker-compose up --build
```

## License
This project is for educational purposes.
```