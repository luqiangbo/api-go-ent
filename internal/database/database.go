package database

import (
	"context"
	"log"

	"api-go-ent/ent"

	_ "github.com/lib/pq"
)

// NewClient 创建新的数据库客户端
func NewClient(dsn string) (*ent.Client, error) {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// 运行数据库迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("Failed to create schema: %v", err)
		return nil, err
	}

	return client, nil
}
