package selection_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestSelectionSort(t *testing.T) {
	t.Run("SelectionSort", tests.DeriveRandom(Sort, 32768))
}
