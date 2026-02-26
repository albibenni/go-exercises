package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	repo := NewInMemoryOrderRepo()
	pay := NewFakePaymentClient(0)
	svc := NewOrderService(repo, pay, logger)
	h := NewHandler(svc)

	mux := http.NewServeMux()
	RegisterRoutes(mux, h)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info("server listening", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
