# bank-api

![bank-api](https://github.com/rishitashaw/bank-api/assets/75828535/13c0c0ab-dad5-4097-a8f8-76622935a93a)

## Setup

1. Setup postgres using docker
   `docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:15.3-alpine3.18`
2. start postgres shell
   `docker exec -it postgres psql -U root`
3. install golang migrate
   `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
4. migarte database
