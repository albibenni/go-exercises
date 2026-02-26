package main

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidInput    = errors.New("invalid input")
	ErrDuplicateReq    = errors.New("duplicate request")
	ErrPaymentDeclined = errors.New("payment declined")
)

type CreateOrderRequest struct {
	CustomerID     string
	AmountCents    int64
	IdempotencyKey string
}

type Order struct {
	ID          string
	CustomerID  string
	AmountCents int64
	Status      string
	CreatedAt   time.Time
}

type OrderRepo interface {
	IdempotencyExists(ctx context.Context, key string) (bool, error)
	CreateOrderTx(ctx context.Context, order *Order, idemKey string) error
}

type PaymentClient interface {
	Charge(ctx context.Context, customerID string, amountCents int64) error
}
