package config

import (
	"sync"

	"api-go-ent/config/env"
)

var (
	once     sync.Once
	instance *Config
)

// Config 应用配置结构
type Config struct {
	Env      *env.Config
	Database DatabaseConfig
}

// DatabaseConfig 数据库配置结构
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Load 加载配置
func Load(envFile string) (*Config, error) {
	var err error
	once.Do(func() {
		var envConfig *env.Config
		envConfig, err = env.LoadEnv(envFile)
		if err != nil {
			return
		}

		instance = &Config{
			Env: envConfig,
			Database: DatabaseConfig{
				Host:     envConfig.DBHost,
				Port:     envConfig.DBPort,
				User:     envConfig.DBUser,
				Password: envConfig.DBPassword,
				DBName:   envConfig.DBName,
			},
		}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

// Get 获取配置实例
func Get() *Config {
	if instance == nil {
		// 如果没有显式加载配置，使用默认环境变量
		cfg, _ := Load("")
		return cfg
	}
	return instance
}
