package main

import (
	"github.com/m-zagornyak/rest-api.git/internal/handler"
	"github.com/m-zagornyak/rest-api.git/internal/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Print("config initializing")

	handlers := new(handler.Handler)
	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
