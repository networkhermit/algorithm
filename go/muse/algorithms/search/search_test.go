package search

import (
	"testing"

	"muse/algorithms/search/tests"
)

func TestBinarySearch(t *testing.T) {
	t.Run("BinarySearch", tests.DeriveIncreasing(BinarySearch, 100000))
}

func TestLinearSearch(t *testing.T) {
	t.Run("LinearSearch", tests.DeriveIncreasing(LinearSearch, 10000))
}
