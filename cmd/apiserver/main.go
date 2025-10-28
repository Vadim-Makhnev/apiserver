package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Vadim-Makhnev/apiserver/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// @title My API
// @version 1.0
// @description This is a sample API with user authentication
// @host localhost:8080

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer token for authentication

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	err = apiserver.Start(config)
	if err != nil {
		log.Fatal(err)
	}
}
