package main

import (
	"database/sql"
	"log"
	"simple_bank/db/api"
	db "simple_bank/db/sqlc"
	"simple_bank/db/util"

	_ "github.com/lib/pq"
)

func main() {

	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(conf.DB_DRIVER, conf.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(conf.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
