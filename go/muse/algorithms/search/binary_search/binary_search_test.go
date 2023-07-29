package binary_search

import (
	"testing"

	"muse/algorithms/search/tests"
)

func TestBinarySearch(t *testing.T) {
	t.Run("BinarySearch", tests.Derive(Find))
}
