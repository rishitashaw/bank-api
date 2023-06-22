postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:15.3-alpine3.18

createdb:
	docker exec -it postgres createdb --username=root --owner=root bank-api 

dropdb:
	docker exec -it postgres dropdb bank-api

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank-api?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank-api?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank-api?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank-api?sslmode=disable" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	docker run --rm -v C:/Users/rishi/OneDrive/Desktop/programs/go/bank-api:/src -w /src kjconroy/sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/rishitashaw/bank-api/db/sqlc Store

test:
	go test -v -cover ./...

server:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server
