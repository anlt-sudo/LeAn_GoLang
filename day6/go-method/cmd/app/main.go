package main

import (
	"fmt"
	"go-method/internal/model"
	"go-method/internal/service"
)



func main() {

	account:= &model.Account{
			Owner:   "Nguyen Van A",
			Balance: 1000,
		}
	
	acc := &service.BankService{
		Account: account,
	}

	fmt.Println("--- Trạng thái ban đầu ---")
	acc.Display()

	fmt.Println("\n--- Thực hiện nạp tiền ---")
	acc.Deposit(500)
	acc.Display()

	fmt.Println("\n--- Thực hiện rút tiền ---")
	acc.Withdraw(200)
	acc.Display()

	fmt.Println("\n--- Thử rút quá số dư ---")
	acc.Withdraw(2000)
	acc.Display()
}