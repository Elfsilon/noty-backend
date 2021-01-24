package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Elfsilon/noty-backend/internal/app/server"
)

var (
	configPath string
)

func initFlags() {
	flag.StringVar(&configPath, "config", "configs/server.toml", "path to config.toml file")
	flag.StringVar(&configPath, "c", "configs/server.toml", "path to config.toml file (shorthand)")

	flag.Parse()
}

func main() {
	initFlags()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
