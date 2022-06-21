package main

import (
	"log"
	"time"

	"github.com/ArifulProtik/sheba-api/config"
	"github.com/ArifulProtik/sheba-api/ent"
	"github.com/ArifulProtik/sheba-api/pkg/auth"
	"github.com/ArifulProtik/sheba-api/pkg/router"
	"github.com/ArifulProtik/sheba-api/pkg/server"
	"github.com/ArifulProtik/sheba-api/pkg/services"
)

func main() {
	cfg, err := config.New("./", "app.env", "env")
	if err != nil {
		log.Fatal(err)
	}

	server := server.New(&cfg.ServerConfig)
	newdb := ent.NewdbClient(&cfg.Postegres)
	mainservice := services.New(newdb)
	exp := time.Minute * 10
	Auth := auth.NewToken(cfg.ServerConfig.Jwt_Key, exp)
	router.InitRouter(server.Echo.Group("/api/v1"), mainservice, Auth, cfg.Jwt_Key)
	server.Run()
}
