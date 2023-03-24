package bank

import (
    "fmt"
    "errors"
)

// Customer ...
type Customer struct {
    Name    string
    Address string
    Phone   string
}

// Account ...
type Account struct {
    Customer
    Number  int32
    Balance float64
}

// Accountクラスのメソッドを定義している?

// Deposit ...
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to deposit should be greater than zero")
    }

    a.Balance += amount
    return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to withdraw should be greater than zero")
    }

    if a.Balance < amount {
        return errors.New("the amount to withdraw should be greater than the account's balance")
    }

    a.Balance -= amount
    return nil
}

// Statement ...
// たぶん今後使わなくなる
func (a *Account) Statement() string {
    return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// Remittance ...
func (a *Account) Transfer(amount float64, dest *Account) error {

	if amount <= 0 {
		return errors.New("I have less than zero money in my possession")
	}

	if a.Balance < amount {
		return errors.New("I am trying to transfer more money than I have in my possession")
	}

    // 送金処理
	a.Withdraw(amount)
	dest.Deposit(amount)
	
    return nil
}

// ここまで

// インターフェイスの実装
// Bank Interface
type Bank interface {
	Statement() string
}

// Bankインターフェイスをパラメーターとして受け取るメソッド
func Statement(bank Bank) {
    // Statementメソッドの呼び出しと値の返却
	return bank.Statement()
}

/*
func Hello() string {
    return "Hey! I'm working!"
}
*/