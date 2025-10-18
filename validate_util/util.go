package validate_util

import "regexp"

// ValidatePhone 用于验证手机号是否合法

/*
ValidatePhone 用于验证手机号是否合法
@param phone string 手机号
@return bool 合法返回true, 否则返回false
*/
func ValidatePhone(phone string) bool {
	// 中国手机号正则表达式
	var phoneRegex = `^1[3-9]\d{9}$`
	match, _ := regexp.MatchString(phoneRegex, phone)
	return match
}

// ValidateEmail 用于验证邮箱是否合法

/*
ValidateEmail 用于验证邮箱是否合法
@param email string 手机号
@return bool 合法返回true, 否则返回false
*/
func ValidateEmail(email string) bool {
	// 中国手机号正则表达式
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
