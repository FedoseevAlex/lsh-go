package jaccard

import (
	"testing"
)

func TestUnionLen(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		description string
		first       []int
		second      []int
		expected    int
	}{
		{description: "empty union", first: []int{}, second: []int{}, expected: 0},
		{description: "first empty", first: []int{}, second: []int{1, 4, 5}, expected: 3},
		{description: "union overlap", first: []int{1, 2, 3}, second: []int{1, 4, 5}, expected: 5},
		{description: "full union", first: []int{1, 2, 3}, second: []int{4, 5, 6}, expected: 6},
	}
	for _, tc := range testCases {
		t.Run(
			tc.description,
			func(t *testing.T) {
				result := unionLen(tc.first, tc.second)
				if result != tc.expected {
					t.Errorf("failed: %s | expected: %d, got: %d", tc.description, tc.expected, result)
				}
			},
		)
	}
}
