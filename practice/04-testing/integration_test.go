package testingpractice

import (
	"context"
	"reflect"
	"testing"
)

func TestSaveUniqueOrders_Integration(t *testing.T) {
	repo := NewInMemoryOrderRepo()
	ctx := context.Background()

	err := SaveUniqueOrders(ctx, repo, []string{"o1", " o2 ", "o1", "o3", "o2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := repo.List()
	want := []string{"o1", "o2", "o3"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected repo content: got=%v want=%v", got, want)
	}
}
