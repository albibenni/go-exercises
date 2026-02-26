package testingpractice

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	ErrNilInput = errors.New("nil input")
	ErrEmptyID  = errors.New("empty order id")
)

// NormalizeUniqueOrderIDs trims IDs, rejects empties, and removes duplicates preserving first appearance.
func NormalizeUniqueOrderIDs(ids []string) ([]string, error) {
	if ids == nil {
		return nil, ErrNilInput
	}

	seen := make(map[string]struct{}, len(ids))
	out := make([]string, 0, len(ids))

	for _, id := range ids {
		normalized := strings.TrimSpace(id)
		if normalized == "" {
			return nil, ErrEmptyID
		}
		if _, ok := seen[normalized]; ok {
			continue
		}
		seen[normalized] = struct{}{}
		out = append(out, normalized)
	}

	return out, nil
}

func ExecuteWithContext(ctx context.Context, op func(context.Context) error) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context check: %w", err)
	}
	if err := op(ctx); err != nil {
		return fmt.Errorf("operation failed: %w", err)
	}
	return nil
}

type OrderRepo interface {
	Save(ctx context.Context, id string) error
	List() []string
}

type InMemoryOrderRepo struct {
	mu   sync.Mutex
	data []string
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{data: []string{}}
}

func (r *InMemoryOrderRepo) Save(_ context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = append(r.data, id)
	return nil
}

func (r *InMemoryOrderRepo) List() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]string, len(r.data))
	copy(out, r.data)
	return out
}

func SaveUniqueOrders(ctx context.Context, repo OrderRepo, ids []string) error {
	normalized, err := NormalizeUniqueOrderIDs(ids)
	if err != nil {
		return fmt.Errorf("normalize ids: %w", err)
	}

	for _, id := range normalized {
		if err := ExecuteWithContext(ctx, func(opCtx context.Context) error {
			return repo.Save(opCtx, id)
		}); err != nil {
			return fmt.Errorf("save id %s: %w", id, err)
		}
	}

	return nil
}
