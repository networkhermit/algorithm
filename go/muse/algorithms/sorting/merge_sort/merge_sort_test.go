package merge_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestMergeSort(t *testing.T) {
	t.Run("MergeSort", tests.Derive(Sort))
}
