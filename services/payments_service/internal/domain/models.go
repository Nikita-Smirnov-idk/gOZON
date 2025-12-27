package domain

import (
	"github.com/google/uuid"
)

type BankAccount struct {
	id      uuid.UUID
	userId  uuid.UUID
	balance int
}

func NewBankAccount(id uuid.UUID, userId uuid.UUID, balance int) *BankAccount {
	return &BankAccount{
		id:      id,
		userId:  userId,
		balance: balance,
	}
}

func (b *BankAccount) GetId() uuid.UUID {
	return b.id
}

func (b *BankAccount) GetUserId() uuid.UUID {
	return b.userId
}

func (b *BankAccount) GetBalance() int {
	return b.balance
}

func (b *BankAccount) TopUpBalance(amount int) {
	b.balance += amount
}

func (b *BankAccount) WithdrawFromBalance(amount int) error {
	if amount > b.balance {
		return ErrInsufficientBalance
	}

	b.balance -= amount
	return nil
}
