package main

import (
	"log"

	"github.com/dentist/api"
	"github.com/dentist/config"
	"github.com/dentist/pkg/db"
	"github.com/dentist/pkg/logger"
	"github.com/dentist/storage"
)

func main() {
	cfg := config.Load()

	psql, _, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Println("Error connecting to database", logger.Error(err))
	}
	
	stor := storage.NewStoragePg(psql)
	
	apiServ := api.New(api.RoutOptions{
		Cfg: &cfg,
		Storage: stor,
	})

	err = apiServ.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}