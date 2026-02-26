package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type OrderService struct {
	repo   OrderRepo
	pay    PaymentClient
	logger *slog.Logger
}

func NewOrderService(repo OrderRepo, pay PaymentClient, logger *slog.Logger) *OrderService {
	return &OrderService{repo: repo, pay: pay, logger: logger}
}

func (s *OrderService) CreateOrder(ctx context.Context, req CreateOrderRequest) (*Order, error) {
	if req.CustomerID == "" || req.AmountCents <= 0 || req.IdempotencyKey == "" {
		return nil, fmt.Errorf("validate create order: %w", ErrInvalidInput)
	}

	exists, err := s.repo.IdempotencyExists(ctx, req.IdempotencyKey)
	if err != nil {
		return nil, fmt.Errorf("check idempotency key %s: %w", req.IdempotencyKey, err)
	}
	if exists {
		return nil, fmt.Errorf("idempotency key %s: %w", req.IdempotencyKey, ErrDuplicateReq)
	}

	if err := retry(ctx, 3, func() error {
		return s.pay.Charge(ctx, req.CustomerID, req.AmountCents)
	}); err != nil {
		return nil, fmt.Errorf("charge customer %s: %w", req.CustomerID, err)
	}

	order := &Order{
		ID:          fmt.Sprintf("ord_%d", time.Now().UnixNano()),
		CustomerID:  req.CustomerID,
		AmountCents: req.AmountCents,
		Status:      "CREATED",
		CreatedAt:   time.Now(),
	}

	if err := s.repo.CreateOrderTx(ctx, order, req.IdempotencyKey); err != nil {
		return nil, fmt.Errorf("persist order %s: %w", order.ID, err)
	}

	s.logger.Info("order created",
		"order_id", order.ID,
		"customer_id", order.CustomerID,
		"amount_cents", order.AmountCents,
	)

	return order, nil
}

func retry(ctx context.Context, attempts int, fn func() error) error {
	var err error
	for i := range attempts {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		err = fn()
		if err == nil {
			return nil
		}
		if i < attempts-1 {
			time.Sleep(50 * time.Millisecond)
		}
	}
	return err
}
