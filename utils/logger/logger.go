package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logger *log.Logger

// Setup 初始化日志配置
func Setup() error {
	// 创建logs目录
	if err := os.MkdirAll("logs", 0o755); err != nil {
		return fmt.Errorf("创建日志目录失败: %v", err)
	}

	// 生成日志文件名 (按天)
	logFileName := filepath.Join("logs", fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))

	// 打开日志文件
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return fmt.Errorf("打开日志文件失败: %v", err)
	}

	// 设置日志格式
	Logger = log.New(logFile, "", log.Ldate|log.Ltime)

	return nil
}
