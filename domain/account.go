package domain

import (
	"errors"
	"time"
)

var ErrInsufficientBalance = errors.New("You don't have enough money! You're so poor :(")

type Account struct {
	id        string
	name      string
	balance   float64
	createdAt time.Time
}

func (a *Account) SetName(name string) {
	a.name = name
}

func (a Account) GetAccountName() string {
	return a.name
}

func NewAccount(id string, name string, balance float64, createdAt time.Time) Account {
	return Account{
		id:        id,
		name:      name,
		balance:   balance,
		createdAt: createdAt,
	}
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return ErrInsufficientBalance
	}
	return nil
}
