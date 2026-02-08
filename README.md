# GoApi

A simple REST API service built with Go for managing user accounts and coin balances.

## Features

- User authentication via token-based authorization
- Get user coin balance
- Mock database for testing
- Chi router for HTTP routing
- Request/response error handling

## Prerequisites

- Go 1.25.6 or higher

## Installation

1. Clone the repository:
```sh
git clone https://github.com/pradeep-iitb/GoApi.git
cd GoApi
```

2. Install dependencies:
```sh
go mod download
```

## Running the Server

Start the API server:
```sh
go run cmd/api/main.go
```

The server will start on `localhost:8000`.

## API Endpoints

### Get Coin Balance

**Endpoint:** `GET /account/coins`

**Required Parameters:**
- `username` (query parameter): The username to fetch coin balance for
- `Authorization` (header): Auth token for the user

**Example Request:**
```sh
curl -H "Authorization: 123ABC" "http://localhost:8000/account/coins?username=alex"
```

**Example Response:**
```json
{
  "Code": 200,
  "Balance": 100
}
```

## Test Users

The mock database includes the following test users:

| Username | Auth Token | Coins |
|----------|-----------|-------|
| alex     | 123ABC    | 100   |
| jason    | 456DEF    | 200   |
| marin    | 789GHI    | 300   |

## Project Structure

```
.
├── api/               # API models and response handlers
├── cmd/api/           # Application entry point
├── internal/
│   ├── handlers/      # HTTP request handlers
│   ├── middleware/    # Authorization and other middleware
│   └── tools/         # Database interface and mock implementation
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
└── README.md          # This file
```

## Dependencies

- [go-chi/chi](https://github.com/go-chi/chi) - HTTP router
- [gorilla/schema](https://github.com/gorilla/schema) - Query parameter decoder
- [sirupsen/logrus](https://github.com/sirupsen/logrus) - Structured logging

## License

[Add your license information here]