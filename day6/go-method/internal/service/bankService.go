package service

import (
	"fmt"
	"go-method/internal/model"
)

type BankService struct {
	Account *model.Account
}

func (s *BankService) Display() {
	fmt.Printf("Chủ tài khoản: %s, Số dư: %.2f\n", s.Account.Owner, s.Account.Balance)
}


func (s *BankService) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Số tiền nạp phải lớn hơn 0.")
		return
	}
	s.Account.Balance += amount
	fmt.Printf("Nạp thành công %.2f. ", amount)
}


func (s *BankService) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Số tiền rút phải lớn hơn 0.")
		return
	}
	if amount > s.Account.Balance {
		fmt.Printf("Lỗi: Không đủ số dư để rút %.2f.\n", amount)
		return
	}
	s.Account.Balance -= amount
	fmt.Printf("Rút thành công %.2f. ", amount)
}