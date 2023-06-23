package jaccard

import (
	"testing"
)

func TestJaccard(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		description string
		first       []int
		second      []int
		expected    float32
	}{
		{description: "unsimilar", first: []int{1, 2, 3}, second: []int{4, 5, 6}, expected: 0},
		{description: "unsimilar one empty", first: []int{}, second: []int{4, 5, 6}, expected: 0},
		{description: "unsimilar both empty", first: []int{}, second: []int{}, expected: 0},
		{description: "similar part 1", first: []int{1, 2, 3}, second: []int{1, 2, 3}, expected: 1.0},
		{description: "similar part 2", first: []int{1, 2, 3}, second: []int{1, 4, 5}, expected: 0.2},
	}

	for _, tc := range testCases {
		t.Run(
			tc.description,
			func(t *testing.T) {
				result := JaccardSimilarity(tc.first, tc.second)
				if result != tc.expected {
					t.Errorf("failed: %s | expected: %f, got: %f", tc.description, tc.expected, result)
				}
			},
		)
	}
}
