package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/rishitashaw/bank-api/api"
	db "github.com/rishitashaw/bank-api/db/sqlc"
	"github.com/rishitashaw/bank-api/util"
)

func main() {
	var err error

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdress)

}
