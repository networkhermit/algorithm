package search

import (
	"testing"

	"muse/algorithms/search/tests"
)

func TestBinarySearch(t *testing.T) {
	t.Run("BinarySearch", tests.Derive(BinarySearch))
}

func TestLinearSearch(t *testing.T) {
	t.Run("LinearSearch", tests.Derive(LinearSearch))
}
