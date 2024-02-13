package wallet

import (
	"fmt"
	"errors"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct{
	fortune Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.fortune += amount
}

// Why here they suggest a pointer?
// It may not see the latest value of 'w' but this means that the program in not thread safe, which is worse
func (w Wallet) Balance() Bitcoin {
	return w.fortune
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.fortune {
		return ErrInsufficientFunds
	}

	w.fortune -= amount
	return nil
}
