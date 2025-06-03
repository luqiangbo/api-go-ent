package validators

import (
	"fmt"
	"regexp"
)

// 常用正则表达式
var (
	phoneRegex    = regexp.MustCompile(`^1[3-9]\d{9}$`)
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)
)

// Price 验证价格
func Price(v float64) error {
	if v <= 0 {
		return fmt.Errorf("价格必须大于0")
	}
	return nil
}

// Phone 验证手机号
func Phone(v string) error {
	if !phoneRegex.MatchString(v) {
		return fmt.Errorf("无效的手机号格式")
	}
	return nil
}

// Email 验证邮箱
func Email(v string) error {
	if !emailRegex.MatchString(v) {
		return fmt.Errorf("无效的邮箱格式")
	}
	return nil
}

// Username 验证用户名
func Username(v string) error {
	if !usernameRegex.MatchString(v) {
		return fmt.Errorf("用户名必须是4-16位字母、数字、下划线或连字符")
	}
	return nil
}

// StringLength 验证字符串长度
func StringLength(min, max int) func(string) error {
	return func(v string) error {
		l := len(v)
		if l < min || l > max {
			return fmt.Errorf("字符串长度必须在%d-%d之间", min, max)
		}
		return nil
	}
}

// IntRange 验证整数范围
func IntRange(min, max int) func(int) error {
	return func(v int) error {
		if v < min || v > max {
			return fmt.Errorf("数值必须在%d-%d之间", min, max)
		}
		return nil
	}
}
