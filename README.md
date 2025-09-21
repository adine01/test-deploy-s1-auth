# Auth Service

Authentication microservice for the Choreo platform testing project.

## Features

- User registration
- User login with JWT tokens
- Token verification
- Password hashing with bcrypt

## API Endpoints

### Health Check

- `GET /health` - Service health check

### Authentication

- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - User login
- `GET /api/auth/verify` - Verify JWT token

## Request/Response Examples

### Register User

```bash
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}
```

Response:

```json
{
  "token": "jwt_token_here",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

### Login

```bash
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

### Verify Token

```bash
GET /api/auth/verify
Authorization: Bearer jwt_token_here
```

## Running the Service

```bash
# Install dependencies
go mod tidy

# Run the service
go run .

# Or build and run
go build -o auth-service
./auth-service
```

## Environment Variables

- `PORT` - Server port (default: 8080)
- `GIN_MODE` - Gin framework mode (debug/release)
- `JWT_SECRET` - JWT signing secret
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name

## Docker

```bash
# Build image
docker build -t auth-service .

# Run container
docker run -p 8080:8080 auth-service
```
