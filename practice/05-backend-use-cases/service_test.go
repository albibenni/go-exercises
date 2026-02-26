package main

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
)

func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func TestOrderService_CreateOrder(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		svc := NewOrderService(NewInMemoryOrderRepo(), NewFakePaymentClient(0), testLogger())

		_, err := svc.CreateOrder(context.Background(), CreateOrderRequest{})
		if !errors.Is(err, ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput, got %v", err)
		}
	})

	t.Run("duplicate idempotency", func(t *testing.T) {
		repo := NewInMemoryOrderRepo()
		svc := NewOrderService(repo, NewFakePaymentClient(0), testLogger())
		ctx := context.Background()

		req := CreateOrderRequest{CustomerID: "c1", AmountCents: 1000, IdempotencyKey: "idem-1"}
		if _, err := svc.CreateOrder(ctx, req); err != nil {
			t.Fatalf("first create should succeed, got %v", err)
		}
		if _, err := svc.CreateOrder(ctx, req); !errors.Is(err, ErrDuplicateReq) {
			t.Fatalf("expected ErrDuplicateReq, got %v", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		svc := NewOrderService(NewInMemoryOrderRepo(), NewFakePaymentClient(0), testLogger())

		order, err := svc.CreateOrder(context.Background(), CreateOrderRequest{
			CustomerID:     "c1",
			AmountCents:    1500,
			IdempotencyKey: "idem-2",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if order.Status != "CREATED" {
			t.Fatalf("expected CREATED, got %s", order.Status)
		}
	})
}
