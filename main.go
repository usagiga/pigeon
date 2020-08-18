package main

import (
	"fmt"
	"github.com/usagiga/envs-go"
	"github.com/usagiga/pigeon/model"
	"log"
)

func main() {
	// Load config from envs
	config := &model.Config{}
	if err := envs.Load(config); err != nil {
		log.Fatalf("Can't load config: %+v", err)
	}

	// Initialize esa.io client
	_ = ConnectToEsa(config)

	fmt.Println("Hello, World!")
}
