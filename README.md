# E-Commerce Backend

This is a simple e-commerce backend built using Gin, PostgreSQL, and SQLC.

## Getting Started

### Prerequisites

- PostgreSQL
- Go


### Installation

1. Clone the repository
2. Run `go mod tidy` to download the required dependencies
3. Run `go run cmd/main.go` to start the server
4. Add `.env` file with the following content:

```bash
# Change the db connection string to your postgres db
GOOSE_DBSTRING="host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"
GOOSE_DRIVER="postgres"
GOOSE_MIGRATION_DIR="./internal/adapters/postgresql/migrations"
```


## Tech Stack

- Go
- Chi
- PostgreSQL
- Postgres-SQLC
- Goose
- Docker Compose

## Usage

### API Endpoints

#### List Products

`GET /products/list`

Example Response:

```json
[
  {
    "id": 1,
    "name": "Product 1",
    "priceInCenters": 100
  },
  {
    "id": 2,
    "name": "Product 2",
    "priceInCenters": 200
  }
]
```

#### Get Product By ID

`GET /products?id={id}`

Example Response:

```json
{
  "id": 1,
  "name": "Product 1",
  "priceInCenters": 100
}
```

#### Place Order

`POST /orders`

Example Request Body:

```json
{
  "customerId": 1,
  "items": [
    {
      "productId": 1,
      "quantity": 2
    },
    {
      "productId": 2,
      "quantity": 1
    }
  ]
}
```

Example Response:

```json
{
  "id": 1,
  "customerId": 1,
  "createdAt": "2023-01-01T00:00:00Z"
}
```
