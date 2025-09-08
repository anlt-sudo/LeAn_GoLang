package service

import (
	"errors"
	"fmt"
)

func validateLogin(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username hoặc password không được để trống")
	}
	if len(username) <= 8 {
		return errors.New("username phải dài hơn 8 ký tự")
	}
	if len(password) <= 6 {
		return errors.New("password phải dài hơn 6 ký tự")
	}
	return nil
}

func Login(username, password string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Quay lại trang login vì đăng nhập không thành công.")
			fmt.Println("Chi tiết lỗi:", r)
		}
	}()

	if err := validateLogin(username, password); err != nil {
		panic(err.Error())
	}
	fmt.Println("Đăng nhập thành công!")
}