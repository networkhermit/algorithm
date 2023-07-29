package bubble_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort", tests.Derive(Sort))
}
