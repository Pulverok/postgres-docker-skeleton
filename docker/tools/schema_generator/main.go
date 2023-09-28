package main

import (
	"log"

	"github.com/Pulverok/postgres-docker-skeleton/schema_generator/internal/config"
	"github.com/Pulverok/postgres-docker-skeleton/schema_generator/internal/services/processor"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to read app config: %v", err)
		return
	}

	processingService := processor.New(cfg)
	err = processingService.Process()
	if err != nil {
		log.Fatalf("failed to process data: %v", err)
		return
	}
}
