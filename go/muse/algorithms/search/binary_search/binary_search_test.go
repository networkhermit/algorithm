package binary_search

import (
	"testing"

	"muse/algorithms/search/tests"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()
	t.Run("Empty", tests.DeriveEmpty(Find))
	t.Run("Increasing", tests.DeriveIncreasing(Find, 100000))
}
