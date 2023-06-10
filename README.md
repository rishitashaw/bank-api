# bank-api

![bank-api](https://github.com/rishitashaw/bank-api/assets/75828535/13c0c0ab-dad5-4097-a8f8-76622935a93a)

## Setup

1. install golang migrate
   `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
2. Start container
   `make postgres`
3. Create database
   `make createdb`
4. Migrate database
   `make migrateup`
