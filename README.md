# 水果价格 API

这是一个使用Go、Gin、PostgreSQL和Ent构建的简单水果价格CRUD API。

## 功能

- 创建水果价格记录
- 获取所有水果价格
- 获取单个水果价格
- 更新水果价格
- 删除水果价格

## 前置要求

- Go 1.16+
- PostgreSQL
- 配置好的PostgreSQL数据库

## 配置

数据库配置在 `config/config.go` 文件中。默认配置如下：

```go
Host: "localhost"
Port: "5432"
User: "root"
Password: "code123"
DBName: "price_db"
```

## API端点

所有接口均使用POST方法：

- POST /fruit-prices/create - 创建新的水果价格
- POST /fruit-prices/list - 获取所有水果价格
- POST /fruit-prices/get - 获取指定ID的水果价格
- POST /fruit-prices/update - 更新指定ID的水果价格
- POST /fruit-prices/delete - 删除指定ID的水果价格

## 请求字段说明

创建和更新请求的必填字段：
- name: 水果名称
- price: 价格（必须为正数）
- unit: 单位（如：kg、个、箱）
- remark: 备注信息

## 运行项目

1. 确保PostgreSQL已运行并创建了数据库
2. 修改配置文件中的数据库连接信息
3. 运行项目：
   ```bash
   make run
   ```

## 示例请求

创建水果价格：
```bash
curl -X POST http://localhost:8080/fruit-prices/create \
  -H "Content-Type: application/json" \
  -d '{
    "name": "苹果",
    "price": 5.99,
    "unit": "kg",
    "remark": "山东烟台红富士"
  }'
```

获取所有水果价格：
```bash
curl -X POST http://localhost:8080/fruit-prices/list
```

获取指定ID的水果价格：
```bash
curl -X POST http://localhost:8080/fruit-prices/get \
  -H "Content-Type: application/json" \
  -d '{"id": "your-uuid-here"}'
```

更新水果价格：
```bash
curl -X POST http://localhost:8080/fruit-prices/update \
  -H "Content-Type: application/json" \
  -d '{
    "id": "your-uuid-here",
    "name": "苹果",
    "price": 6.99,
    "unit": "kg",
    "remark": "山东烟台红富士，新到货"
  }'
```

删除水果价格：
```bash
curl -X POST http://localhost:8080/fruit-prices/delete \
  -H "Content-Type: application/json" \
  -d '{"id": "your-uuid-here"}'
``` 