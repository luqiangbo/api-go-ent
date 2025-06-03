package database

import (
	"api-go-ent/config"
	"api-go-ent/ent"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func NewClient(cfg config.DatabaseConfig) *ent.Client {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
} 