package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/buniekbua/simplebank/api"
	db "github.com/buniekbua/simplebank/db/sqlc"
	"github.com/buniekbua/simplebank/util"
)

func main() {
	config, err := util.LoadConfig("/simplebank")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
