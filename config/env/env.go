package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config 环境变量配置结构
type Config struct {
	// 应用配置
	AppName  string
	AppEnv   string
	AppDebug bool
	AppPort  int

	// 数据库配置
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// 日志配置
	LogLevel  string
	LogFormat string

	// CORS配置
	CorsAllowOrigins string
	CorsAllowMethods string
	CorsAllowHeaders string
	CorsMaxAge       int
}

// LoadEnv 加载环境变量
func LoadEnv(envFile string) (*Config, error) {
	// 如果环境变量文件存在，则加载
	if _, err := os.Stat(envFile); err == nil {
		if err := godotenv.Load(envFile); err != nil {
			return nil, fmt.Errorf("error loading env file: %v", err)
		}
	}

	config := &Config{
		// 应用配置
		AppName:  getEnvOrDefault("APP_NAME", "fruit-price-api"),
		AppEnv:   getEnvOrDefault("APP_ENV", "development"),
		AppDebug: getEnvAsBoolOrDefault("APP_DEBUG", true),
		AppPort:  getEnvAsIntOrDefault("APP_PORT", 8080),

		// 数据库配置
		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "5432"),
		DBUser:     getEnvOrDefault("DB_USER", "root"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "code123"),
		DBName:     getEnvOrDefault("DB_NAME", "price_db"),

		// 日志配置
		LogLevel:  getEnvOrDefault("LOG_LEVEL", "debug"),
		LogFormat: getEnvOrDefault("LOG_FORMAT", "json"),

		// CORS配置
		CorsAllowOrigins: getEnvOrDefault("CORS_ALLOW_ORIGINS", "*"),
		CorsAllowMethods: getEnvOrDefault("CORS_ALLOW_METHODS", "GET,POST,PUT,DELETE,OPTIONS"),
		CorsAllowHeaders: getEnvOrDefault("CORS_ALLOW_HEADERS", "Origin,Content-Type,Accept,Authorization"),
		CorsMaxAge:       getEnvAsIntOrDefault("CORS_MAX_AGE", 86400),
	}

	return config, nil
}

// getEnvOrDefault 获取环境变量，如果不存在则返回默认值
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsIntOrDefault 获取整数类型的环境变量
func getEnvAsIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsBoolOrDefault 获取布尔类型的环境变量
func getEnvAsBoolOrDefault(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

// IsDevelopment 是否为开发环境
func (c *Config) IsDevelopment() bool {
	return c.AppEnv == "development"
}

// IsProduction 是否为生产环境
func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}

// GetDSN 获取数据库连接字符串
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}
