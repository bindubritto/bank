package main

import (
	"database/sql"
	"log"

	"github.com/bindubritto/bank/api"
	db "github.com/bindubritto/bank/db/sqlc"
	"github.com/bindubritto/bank/utils"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
