package main

import (
	"flag"
	"github.com/Bakhram74/rest-api.git/internal/app/apiserver"
	"github.com/joho/godotenv"
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
	config := apiserver.Config{
		BindAddr: os.Getenv("bind_addr"),
		LogLevel: os.Getenv("log_level"),
	}
	s := apiserver.New(&config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
