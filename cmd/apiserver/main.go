package main

import (
	"flag"
	"github.com/Bakhram74/rest-api.git/internal/app/apiserver"
	"github.com/Bakhram74/rest-api.git/internal/app/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/.env", "path to config")
}
func main() {
	flag.Parse()
	err := godotenv.Load(configPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	storeConfig := store.Config{DatabaseUrl: os.Getenv("database_url")}
	config := apiserver.Config{
		BindAddr: os.Getenv("bind_addr"),
		LogLevel: os.Getenv("log_level"),
		Store:    &storeConfig,
	}
	s := apiserver.New(&config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
