package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"api-go-ent/ent"

	_ "github.com/lib/pq"
)

// InitDB 初始化数据库连接
func InitDB() (*ent.Client, error) {
	// 从环境变量获取数据库配置
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 构建数据库连接字符串
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	// 创建 ent client
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	// 运行数据库迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}

	return client, nil
}
