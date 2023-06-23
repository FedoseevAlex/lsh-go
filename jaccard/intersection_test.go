package jaccard

import (
	"testing"
)

func TestIntersectionLen(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		description string
		first       []int
		second      []int
		expected    int
	}{
		{description: "empty intersection", first: []int{1, 2, 3}, second: []int{4, 5, 6}, expected: 0},
		{description: "empty intersection part 2", first: []int{}, second: []int{4, 5, 6}, expected: 0},
		{description: "intersection of empty", first: []int{}, second: []int{}, expected: 0},
		{description: "not empty intersection", first: []int{1, 2, 3}, second: []int{1, 4, 5}, expected: 1},
	}

	for _, tc := range testCases {
		t.Run(
			tc.description,
			func(t *testing.T) {
				result := intersectionLen(tc.first, tc.second)
				if result != tc.expected {
					t.Errorf("failed: %s | expected: %d, got: %d", tc.description, tc.expected, result)
				}
			},
		)
	}
}
