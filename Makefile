# 定义所有可用的任务
.PHONY: generate run debug clean help

# 默认任务显示帮助信息
default: help

# 生成Ent代码
generate:
	@echo "正在生成Ent代码..."
	@go generate ./ent
	@echo "Ent代码生成完成"

# 运行项目
run:
	@echo "正在启动项目..."
	@go run main.go

# 调试项目
debug:
	@echo "正在启动调试..."
	@dlv debug main.go

# 清理生成的文件
clean:
	@echo "正在清理生成的文件..."
	@rm -rf ./ent/generated
	@echo "清理完成"

# 显示帮助信息
help:
	@echo "可用的命令："
	@echo "  make generate  - 生成Ent代码"
	@echo "  make run      - 运行项目"
	@echo "  make debug    - 调试项目"
	@echo "  make clean    - 清理生成的文件" 