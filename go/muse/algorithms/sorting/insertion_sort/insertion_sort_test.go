package insertion_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestInsertionSort(t *testing.T) {
	t.Run("InsertionSort", tests.DeriveRandom(Sort, 32768))
}
