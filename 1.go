package main

import (
 "errors"
 "fmt"
)

type BankAccount struct {
 AccountNumber string
 HolderName    string
 Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
 b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
 if b.Balance < amount {
  return errors.New("недостаточно средств")
 }
 b.Balance -= amount
 return nil
}

func (b *BankAccount) GetBalance() float64 {
 return b.Balance
}

func main() {
 acc := BankAccount{"123", "Матвей", 1000}
 acc.Deposit(500)
 fmt.Println("Баланс:", acc.GetBalance())
 err := acc.Withdraw(2000)
 if err != nil {
  fmt.Println("Ошибка:", err)
 }
}
