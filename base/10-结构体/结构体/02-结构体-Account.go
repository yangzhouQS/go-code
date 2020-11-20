package main

import (
	"fmt"
)

type Account struct {
	AccountNo string  // 账号
	Pwd       string  // 密码
	Balance   float64 // 余额
}

/*
存款
*/
func (account *Account) Deposited(money float64, pwd string) {
	if account.Pwd != pwd {
		fmt.Printf("你的密码[ %s ]不正确 \n", pwd)
		return
	}
	if money < 0 {
		fmt.Printf("你输入的金额 %0.3f 不正确\n", money)
		return
	}
	account.Balance += money
	fmt.Printf("你的余额为 %0.3f \n", account.Balance)
}

/*
取款
*/
func (account *Account) WithDraw(money float64, pwd string) {
	if account.Pwd != pwd {
		fmt.Println("输入密码不正确")
		return
	}
	if money < 0 || money > account.Balance {
		fmt.Printf("输入的取款金额 %0.3f不在正确 \n", money)
		return
	}
	account.Balance -= money
	fmt.Printf("取款%0.3f成功 \n", money)
}

/**
查询余额
*/
func (account *Account) QueryMoney(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("输入密码不正确")
		return
	}
	fmt.Printf("你的账号为%s,余额为 %0.3f \n", account.AccountNo, account.Balance)
}
func main() {
	account := Account{"A001", "123", 666.66}
	account.Deposited(33, "123")
	account.WithDraw(3, "123")
	account.QueryMoney("123")

}
