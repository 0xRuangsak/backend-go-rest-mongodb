# User Management API - Go Backend

A production-ready RESTful API built with Go, MongoDB, and JWT authentication for managing users. Features Docker containerization, and automated testing.

## 🚀 Features
- ✅ **User CRUD operations** - Create, read, update, delete users
- ✅ **JWT authentication** - Secure token-based authentication  
- ✅ **MongoDB integration** - Persistent data storage with full CRUD
- ✅ **Authentication middleware** - Protected endpoints with Bearer tokens
- ✅ **Password hashing** - Secure bcrypt password encryption
- ✅ **Architecture** - Domain-driven design with separated concerns
- ✅ **Docker containerization** - Production-ready container setup
- ✅ **Automated testing** - Integration tests with Docker
- ✅ **Environment configuration** - Configurable via environment variables

## 🏗️ Architecture

###  Architecture Layers
```
┌─────────────────┐
│   HTTP Layer    │ ← REST API endpoints
│   (Handlers)    │
├─────────────────┤
│ Service Layer   │ ← Business logic & JWT auth
│ (Business Logic)│
├─────────────────┤
│Repository Layer │ ← Data access & MongoDB ops
│ (Data Access)   │
├─────────────────┤
│  Domain Layer   │ ← Core business entities
│ (Entities)      │
└─────────────────┘
```

### Tech Stack
- **Language:** Go 1.24
- **Database:** MongoDB 7.0
- **Authentication:** JWT with HMAC-SHA256
- **Architecture:** Hexagonal Architecture
- **Containerization:** Docker & Docker Compose
- **Testing:** Automated integration tests

## 🐳 Quick Start with Docker

### Prerequisites
- Docker and Docker Compose installed

### Run the Application
```bash
# Clone the repository
git clone https://github.com/0xRuangsak/backend-go-rest-mongodb.git
cd backend-go-rest-mongodb

# Start all services (API + MongoDB)
docker compose up --build

# The API will be available at http://localhost:8080
```

### Run Automated Tests
```bash
# Run integration tests
docker compose --profile test up --build

# Or run individual test script
chmod +x scripts/test-api.sh
./scripts/test-api.sh
```

## 🔧 Local Development

### Prerequisites
- Go 1.24+
- MongoDB running locally

### Setup
```bash
# Install dependencies
go mod tidy

# Set environment variables
export MONGODB_URI=mongodb://localhost:27017/userdb
export JWT_SECRET=your-super-secret-key
export PORT=8080

# Run the application
go run cmd/server/main.go
```

### Build
```bash
# Build for production
go build -o user-api cmd/server/main.go

# Run binary
./user-api
```

## 📋 API Endpoints

### Authentication Endpoints
- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login user and receive JWT token

### Protected User Endpoints (Require Bearer Token)
- `GET /users` - Get all users
- `GET /users/{id}` - Get user by ID  
- `PUT /users/{id}` - Update user information
- `DELETE /users/{id}` - Delete user

### Example Usage

#### Register User
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'
```

#### Login User
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

#### Access Protected Endpoint
```bash
curl -X GET http://localhost:8080/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## ⚙️ Configuration

### Environment Variables
| Variable | Description | Default |
|----------|-------------|---------|
| `MONGODB_URI` | MongoDB connection string | `mongodb://localhost:27017/userdb` |
| `JWT_SECRET` | Secret key for JWT signing | `default-secret-key` |
| `PORT` | Server port | `8080` |

### Docker Environment
Environment variables are configured in `docker-compose.yaml` for containerized deployment.

## 🧪 Testing

### Automated Integration Tests
```bash
# Run with Docker (recommended)
docker compose --profile test up --build

# Run locally (requires API to be running)
./scripts/test-api.sh
```

### Manual Testing
```bash
# Unit tests
go test ./...

# With coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

## 📁 Project Structure

```
user-api/
├── cmd/server/           # Application entry point
├── internal/
│   ├── domain/          # Business entities and interfaces
│   ├── repository/      # Data access layer
│   ├── service/         # Business logic layer
│   ├── handler/         # HTTP handlers
│   └── middleware/      # HTTP middleware
├── pkg/auth/            # JWT authentication utilities
├── scripts/             # Automation scripts
├── docker-compose.yml   # Docker orchestration
├── Dockerfile          # Container definition
└── README.md
```

## 🔐 Security Features

- **Password Hashing**: Bcrypt with salt
- **JWT Authentication**: HMAC-SHA256 signed tokens
- **Protected Routes**: Middleware-based authentication
- **Input Validation**: Request validation and sanitization
- **Environment Secrets**: Externalized configuration

## 🚀 Deployment

### Docker Production Deployment
```bash
# Build production images
docker compose build

# Run in production mode
docker compose up -d

# View logs
docker compose logs -f
```

### Manual Deployment
```bash
# Build optimized binary
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-api cmd/server/main.go

# Set production environment variables
export MONGODB_URI=mongodb://your-production-db:27017/userdb
export JWT_SECRET=your-production-secret
export PORT=8080

# Run
./user-api
```

## 📚 Learning Outcomes

This project demonstrates:
- **Clean Architecture** implementation in Go
- **Domain-Driven Design** principles
- **JWT authentication** and middleware patterns
- **MongoDB integration** with Go drivers
- **Docker containerization** best practices
- **Automated testing** strategies
- **RESTful API** design and implementation

## 📄 License

This project is for educational purposes and learning Go backend development.

---

**Built with ❤️ using Go, MongoDB, and Docker.**