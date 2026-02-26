package main

import (
	"context"
	"errors"
)

type Handler struct {
	svc *OrderService
}

func NewHandler(svc *OrderService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateOrder(ctx context.Context, req CreateOrderRequest) (int, any) {
	order, err := h.svc.CreateOrder(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidInput):
			return 400, err.Error()
		case errors.Is(err, ErrDuplicateReq):
			return 409, err.Error()
		case errors.Is(err, ErrPaymentDeclined):
			return 402, err.Error()
		default:
			return 500, "internal error"
		}
	}

	return 201, order
}
