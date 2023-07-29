package quick_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestQuickSort(t *testing.T) {
	t.Run("QuickSort", tests.Derive(Sort))
}
