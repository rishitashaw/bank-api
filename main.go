package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/rishitashaw/bank-api/api"
	db "github.com/rishitashaw/bank-api/database/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/bank-api?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	var err error

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store:=db.NewStore(conn)
	server:=api.NewServer(store)

	err=server.Start(serverAddress)

}