package main

import (
	"log"

	"github.com/smolneko-team/smolneko/config"
	"github.com/smolneko-team/smolneko/internal/app"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
