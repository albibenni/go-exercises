package main

import (
	"context"
	"fmt"
	"sync"
)

type InMemoryOrderRepo struct {
	mu          sync.Mutex
	orders      map[string]Order
	idemToOrder map[string]string
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders:      map[string]Order{},
		idemToOrder: map[string]string{},
	}
}

func (r *InMemoryOrderRepo) IdempotencyExists(_ context.Context, key string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.idemToOrder[key]
	return ok, nil
}

func (r *InMemoryOrderRepo) CreateOrderTx(_ context.Context, order *Order, idemKey string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.idemToOrder[idemKey]; exists {
		return fmt.Errorf("idempotency key already used: %w", ErrDuplicateReq)
	}
	r.orders[order.ID] = *order
	r.idemToOrder[idemKey] = order.ID
	return nil
}

type FakePaymentClient struct {
	failuresLeft int
}

func NewFakePaymentClient(failures int) *FakePaymentClient {
	return &FakePaymentClient{failuresLeft: failures}
}

func (f *FakePaymentClient) Charge(_ context.Context, _ string, _ int64) error {
	if f.failuresLeft > 0 {
		f.failuresLeft--
		return ErrPaymentDeclined
	}
	return nil
}
