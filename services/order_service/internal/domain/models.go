package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type OrderStatus int

const (
	StatusNew OrderStatus = iota
	StatusFinished
	StatusCancelled
)

func (s OrderStatus) String() string {
	switch s {
	case StatusNew:
		return "NEW"
	case StatusFinished:
		return "FINISHED"
	case StatusCancelled:
		return "CANCELLED"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", s)
	}
}

type Order struct {
	id          uuid.UUID
	userId      uuid.UUID
	amount      int
	description string
	status      OrderStatus
}

func NewOrder(id uuid.UUID, userId uuid.UUID, amount int, description string, status OrderStatus) *Order {
	return &Order{
		id:          id,
		userId:      userId,
		amount:      amount,
		description: description,
		status:      status,
	}
}

func (o *Order) GetId() uuid.UUID {
	return o.id
}
func (o *Order) GetUserId() uuid.UUID {
	return o.userId
}

func (o *Order) GetAmount() int {
	return o.amount
}

func (o *Order) GetDescription() string {
	return o.description
}

func (o *Order) GetStatus() OrderStatus {
	return o.status
}

func (o *Order) SetStatus(status OrderStatus) {
	o.status = status
}
