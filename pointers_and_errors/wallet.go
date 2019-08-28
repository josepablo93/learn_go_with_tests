package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is type of crypto
type Bitcoin int

// Stringer makes a stringer
type Stringer interface {
	String() string
}

// Wallet Type
type Wallet struct {
	balance Bitcoin
}

// ErrInsufficientFunds error type when the user wants to withdraw but doesnt have the amount necessary that he wants to do the operation
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit adds the amount wanted
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the value of the wallet currnet balance.
// We use a pointer to specify the current wallet wanted instead
// of the 'copy' it creates when it is called without the pointer
func (w *Wallet) Balance() Bitcoin {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	return w.balance

}

// Withdraw a certain amount of bitcoins
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
