package custom // Replace with your actual package name

import (
	"crypto/internal/fireblocks/vault"
	"testing"
)

func TestSplit(t *testing.T) {
	// Test cases with different input lengths and batch sizes
	tests := []struct {
		name      string
		accounts  []vault.ListVaultsResponseAccount
		batchSize int
		expected  []Batch
	}{
		{
			name:      "Single batch - less than batch size",
			accounts:  []vault.ListVaultsResponseAccount{{ID: "acc1"}, {ID: "acc2"}},
			batchSize: 3,
			expected: []Batch{
				{
					Accounts: []vault.ListVaultsResponseAccount{{ID: "acc1"}, {ID: "acc2"}},
					Len:      2,
				},
			},
		},
		{
			name:      "Multiple batches - exact division",
			accounts:  []vault.ListVaultsResponseAccount{{ID: "acc1"}, {ID: "acc2"}, {ID: "acc3"}, {ID: "acc4"}},
			batchSize: 2,
			expected: []Batch{
				{
					Accounts: []vault.ListVaultsResponseAccount{{ID: "acc1"}, {ID: "acc2"}},
					Len:      2,
				},
				{
					Accounts: []vault.ListVaultsResponseAccount{{ID: "acc3"}, {ID: "acc4"}},
					Len:      2,
				},
			},
		},
		{
			name:      "Multiple batches - uneven division",
			accounts:  []vault.ListVaultsResponseAccount{{ID: "acc1"}, {ID: "acc2"}, {ID: "acc3"}},
			batchSize: 2,
			expected: []Batch{
				{
					Accounts: []vault.ListVaultsResponseAccount{{ID: "acc1"}, {ID: "acc2"}},
					Len:      2,
				},
				{
					Accounts: []vault.ListVaultsResponseAccount{{ID: "acc3"}},
					Len:      1,
				},
			},
		},
		{
			name:      "Empty slice",
			accounts:  []vault.ListVaultsResponseAccount{},
			batchSize: 2,
			expected:  []Batch{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := split(tc.accounts, tc.batchSize)
			// Compare expected and actual batches
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %d batches, got %d", len(tc.expected), len(actual))
				return
			}
			for i := range actual {
				if len(actual[i].Accounts) != len(tc.expected[i].Accounts) {
					t.Errorf("Expected batch %d to have length %d, got %d", i+1, len(tc.expected[i].Accounts), len(actual[i].Accounts))
					return
				}
				// You can further compare individual account IDs within each batch if needed
			}
		})
	}
}
