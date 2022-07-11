package main

import (
	"flag"
	"github.com/Bakhram74/rest-api.git/internal/app/apiserver"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
}
func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.Decode(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
