package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/baldurstod/go-steam-assetclassinfo/config"
	"github.com/baldurstod/go-steam-assetclassinfo/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.Config{}

	var content []byte
	var err error

	if content, err = os.ReadFile("./var/config.json"); err != nil {
		err := fmt.Errorf("error while reading configuration file: %w", err)
		panic(err)
	}

	if err = json.Unmarshal(content, &config); err != nil {
		err := fmt.Errorf("error while reading configuration file: %w", err)
		panic(err)
	}

	server.StartServer(config)
}
