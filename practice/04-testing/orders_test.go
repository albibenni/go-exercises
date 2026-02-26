package testingpractice

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestNormalizeUniqueOrderIDs(t *testing.T) {
	t.Run("table driven", func(t *testing.T) {
		tests := []struct {
			name    string
			in      []string
			want    []string
			wantErr error
		}{
			{
				name:    "nil input",
				in:      nil,
				wantErr: ErrNilInput,
			},
			{
				name: "empty slice",
				in:   []string{},
				want: []string{},
			},
			{
				name:    "contains empty id",
				in:      []string{"A", "   ", "B"},
				wantErr: ErrEmptyID,
			},
			{
				name: "deduplicate preserve order",
				in:   []string{" A ", "B", "A", "B", "C"},
				want: []string{"A", "B", "C"},
			},
		}

		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				got, err := NormalizeUniqueOrderIDs(tc.in)
				if tc.wantErr != nil {
					if !errors.Is(err, tc.wantErr) {
						t.Fatalf("expected error %v, got %v", tc.wantErr, err)
					}
					return
				}
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("unexpected output: got=%v want=%v", got, tc.want)
				}
			})
		}
	})
}

func TestExecuteWithContext_Canceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := ExecuteWithContext(ctx, func(context.Context) error {
		t.Fatal("operation should not run when context is canceled")
		return nil
	})

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}
