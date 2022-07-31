package main

import (
    "log"

    "smolneko/config"
    "smolneko/internal/app"
)

func main() {

    cfg, err := config.NewConfig()
    if err != nil {
        log.Fatalf("config error: %s", err)
    }

    app.Run(cfg)
}
