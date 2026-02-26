package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestServer() *http.ServeMux {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	repo := NewInMemoryOrderRepo()
	pay := NewFakePaymentClient(0)
	svc := NewOrderService(repo, pay, logger)
	h := NewHandler(svc)
	mux := http.NewServeMux()
	RegisterRoutes(mux, h)
	return mux
}

func TestHTTP_CreateOrderEndpoint(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mux := setupTestServer()
		body, _ := json.Marshal(CreateOrderRequest{
			CustomerID:     "cust_1",
			AmountCents:    1200,
			IdempotencyKey: "idem_http_1",
		})

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d, body=%s", rr.Code, rr.Body.String())
		}
	})

	t.Run("duplicate idempotency", func(t *testing.T) {
		mux := setupTestServer()
		payload := CreateOrderRequest{CustomerID: "cust_1", AmountCents: 1200, IdempotencyKey: "idem_http_dup"}
		body, _ := json.Marshal(payload)

		first := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
		first.Header.Set("Content-Type", "application/json")
		rr1 := httptest.NewRecorder()
		mux.ServeHTTP(rr1, first)

		secondBody, _ := json.Marshal(payload)
		second := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(secondBody))
		second.Header.Set("Content-Type", "application/json")
		rr2 := httptest.NewRecorder()
		mux.ServeHTTP(rr2, second)

		if rr2.Code != http.StatusConflict {
			t.Fatalf("expected 409, got %d, body=%s", rr2.Code, rr2.Body.String())
		}
	})
}
