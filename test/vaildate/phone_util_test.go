package vaildate

import (
	"testing"

	"github.com/bys-Cai/bys-tools/validate_util"
)

/*
*
测试正则手机号验证
*/
func TestIsValidPhone(t *testing.T) {
	phones := []struct {
		phone string
		want  bool
	}{
		{"13812345678", true},
		{"12345678901", false},
		{"19999999999", true},
		{"11111111111", false},
		{"13403504277", false},
		{"19912456998", true},
	}
	for _, phone := range phones {
		result := validate_util.ValidatePhone(phone.phone)
		t.Logf("IsValidPhone(%q) = %v; want %v", phone.phone, result, phone.want)
	}

}

func TestIsValidEmail(t *testing.T) {
	emails := []struct {
		email string
		want  bool
	}{
		{"13812345678@163.com", true},
		{"12345678901@126.com", true},
		{"1546254@aa.com", true},
		{"165456@22.", false},
	}
	for _, email := range emails {
		result := validate_util.ValidateEmail(email.email)
		t.Logf("IsValidEmail(%q) = %v; want %v", email.email, result, email.want)
	}

}
