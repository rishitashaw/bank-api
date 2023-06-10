postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:15.3-alpine3.18

createdb:
	docker exec -it postgres createdb --username=root --owner=root bank-api 

dropdb:
	docker exec -it postgres dropdb bank-api

migrateup:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/bank-api?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/bank-api?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v C:/Users/rishi/Desktop/programs/go/bank-api:/src -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
