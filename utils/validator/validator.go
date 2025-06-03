package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 自定义验证标签
const (
	priceRegex = `^\d+(\.\d{1,2})?$` // 价格格式：整数或最多两位小数
)

// Setup 设置自定义验证器
func Setup() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册价格格式验证器
		_ = v.RegisterValidation("price", validatePrice)
	}
}

// validatePrice 验证价格格式
func validatePrice(fl validator.FieldLevel) bool {
	price := fl.Field().Float()
	return price > 0
}
