# GoAPI Essential

A comprehensive Go learning project that demonstrates fundamental Go concepts, API development, and database integration patterns. This repository contains four distinct modules showcasing different aspects of Go programming from basic syntax to advanced web API development.

## 📚 Project Structure

```
INDIV-GoAPI-essential/
├── basic Go/          # Go fundamentals and syntax examples
├── GoAPI/            # REST API with Fiber framework (in-memory)
├── GoDB/             # Database operations with PostgreSQL
├── GORM/             # ORM-based API with GORM and PostgreSQL
└── README.md         # This file
```

## 🚀 Modules Overview

### 1. Basic Go (`basic Go/`)
**Learning Go fundamentals and core concepts**

- **Purpose**: Introduction to Go programming language basics
- **Key Topics**:
  - Variables and data types
  - Control structures (loops, conditionals)
  - Data structures (arrays, slices, maps, structs)
  - Functions and interfaces
  - Error handling
  - Package management

**Key Files**:
- `main.go` - Main examples demonstrating Go concepts
- `Jo/test.go` - Package visibility examples (public/private functions)
- `Jo/internal/testInternal/` - Internal package demonstration

### 2. GoAPI (`GoAPI/`)
**REST API development with Fiber framework**

- **Purpose**: Building RESTful APIs using Go Fiber framework
- **Features**:
  - JWT authentication
  - CRUD operations for books
  - Environment variable configuration
  - Middleware implementation
  - In-memory data storage

**Key Technologies**:
- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- [JWT](https://github.com/golang-jwt/jwt) - Authentication
- [godotenv](https://github.com/joho/godotenv) - Environment variables

**API Endpoints**:
```
POST   /login           # User authentication
GET    /books           # Get all books (protected)
GET    /books/:id       # Get book by ID (protected)
POST   /books           # Create new book (protected)
PUT    /books/:id       # Update book (protected)
DELETE /books/:id       # Delete book (protected)
GET    /env             # Get environment variables (protected)
```

### 3. GoDB (`GoDB/`)
**Raw SQL database operations with PostgreSQL**

- **Purpose**: Direct database interaction using standard SQL
- **Features**:
  - PostgreSQL connection management
  - Raw SQL queries and operations
  - CRUD operations for products
  - Fiber API endpoints
  - Docker PostgreSQL setup

**Key Technologies**:
- [lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- PostgreSQL database
- Docker & Docker Compose

**API Endpoints**:
```
GET    /products        # Get all products
GET    /products/:id    # Get product by ID
POST   /products        # Create new product
PUT    /products/:id    # Update product
DELETE /products/:id    # Delete product
```

### 4. GORM (`GORM/`)
**ORM-based API with advanced features**

- **Purpose**: Database operations using GORM ORM with advanced features
- **Features**:
  - GORM ORM integration
  - User authentication with JWT
  - Password hashing with bcrypt
  - Soft deletes
  - Database migrations
  - Comprehensive API endpoints

**Key Technologies**:
- [GORM](https://gorm.io/) - Go ORM library
- [bcrypt](https://golang.org/x/crypto/bcrypt) - Password hashing
- [JWT](https://github.com/golang-jwt/jwt) - Authentication
- PostgreSQL with GORM driver
- [Fiber v2](https://github.com/gofiber/fiber) - Web framework

**API Endpoints**:
```
# User Management
POST   /users/register  # User registration
POST   /users/login     # User login

# Book Management (protected routes)
GET    /books           # Get all books
GET    /books/:id       # Get book by ID
POST   /books           # Create new book
PUT    /books/:id       # Update book
DELETE /books/:id       # Delete book (soft delete)
```

## 🛠️ Prerequisites

- **Go 1.24.3** or later
- **Docker & Docker Compose** (for database setup)
- **PostgreSQL** (for GoDB and GORM modules)

## 🚀 Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/RookieJoel/GoAPI-essential.git
cd GoAPI-essential
```

### 2. Basic Go Module
```bash
cd "basic Go"
go mod tidy
go run main.go
```

### 3. GoAPI Module
```bash
cd GoAPI
go mod tidy

# Create .env file
echo "JWT_SECRET=your-secret-key" > .env

go run main.go
# Server runs on http://localhost:8080
```

### 4. GoDB Module
```bash
cd GoDB

# Start PostgreSQL with Docker
docker-compose up -d

# Wait for database to be ready, then run
go mod tidy
go run .
# Server runs on http://localhost:8080
```

### 5. GORM Module
```bash
cd GORM

# Start PostgreSQL with Docker (if not already running)
docker-compose up -d

# Run the application
go mod tidy
go run .
# Server runs on http://localhost:8080
```

## 🐳 Database Setup

Both GoDB and GORM modules use PostgreSQL. The Docker Compose configuration provides:

- **PostgreSQL Database**
  - Host: `localhost`
  - Port: `5433`
  - Database: `mydatabase`
  - User: `myuser`
  - Password: `mypassword`

- **pgAdmin** (Database Management UI)
  - URL: http://localhost:5050
  - Email: `admin@admin.com`
  - Password: `admin`

Start the database services:
```bash
cd GoDB  # or GORM
docker-compose up -d
```

## 📝 API Testing

### GoAPI Module Testing
```bash
# Login to get JWT token
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password"}'

# Use the returned token in subsequent requests
curl -X GET http://localhost:8080/books \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### GoDB/GORM Module Testing
```bash
# Create a product/book
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Sample Product", "price": 100}'

# Get all products/books
curl -X GET http://localhost:8080/products
```

## 🎯 Learning Objectives

This project demonstrates:

1. **Go Language Fundamentals**
   - Syntax and basic concepts
   - Package management and visibility
   - Error handling patterns

2. **Web API Development**
   - RESTful API design
   - HTTP request/response handling
   - Middleware implementation

3. **Authentication & Security**
   - JWT token-based authentication
   - Password hashing with bcrypt
   - Route protection

4. **Database Integration**
   - Raw SQL operations
   - ORM usage patterns
   - Database migrations
   - Connection management

5. **Modern Go Practices**
   - Project structure organization
   - Dependency management with go modules
   - Environment configuration
   - Docker containerization

## 🤝 Contributing

This is a learning project. Feel free to:
- Fork and experiment
- Submit issues for improvements
- Share your learning experience

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👨‍💻 Author

**RookieJoel**
- GitHub: [@RookieJoel](https://github.com/RookieJoel)

---

*Happy Learning! 🚀*