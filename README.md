# bank-api

The Bank Service API is a backend web service for a bank, providing APIs for the frontend to perform various banking operations. This project is developed using Golang and incorporates several technologies and tools such as Docker, Kubernetes, Redis, AWS, PostgreSQL, unit testing, continuous integration, continuous deployment, gRPC, and Gin.

## Features

1. The Simple Bank Service offers the following features through its APIs:
2. Create and manage bank accounts, including owner's name, balance, and currency.
3. Record all balance changes to each account by creating account entry records.
4. Perform money transfers between two accounts, ensuring atomicity and consistency.

## Database model

![bank-api](https://github.com/rishitashaw/bank-api/assets/75828535/28f131a8-7645-4ef9-ba15-82a3fda10080)

# Local Development Setup

To set up the development environment locally, follow these steps:

## Prerequisites

1. Docker Desktop
2. TablePlus
3. Golang
4. Golang Migrate: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
5. DB Docs: `npm install -g dbdocs`
6. DBML CLI: `npm install -g @dbml/cli`
7. Sqlc: `docker pull kjconroy/sqlc`
8. Gomock: `go install github.com/golang/mock/mockgen@v1.6.0`

## Infrastructure Setup

1. Start the PostgreSQL container: `make postgres`
2. Create the simple_bank database: `make createdb`
3. Run database migrations to set up the schema:

```shell
Up all versions: `make migrateup`
Down all versions: `make migratedown`
```

### Code Generation

1. Generate SQL CRUD with sqlc: `make sqlc`
