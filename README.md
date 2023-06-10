# bank-api

![bank-api](https://github.com/rishitashaw/bank-api/assets/75828535/aaeb617b-ff4f-4f1b-8031-6260c70e00b9)


## Setup

1. install golang migrate
   `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
2. Start container
   `make postgres`
3. Create database
   `make createdb`
4. Migrate database
   `make migrateup`
