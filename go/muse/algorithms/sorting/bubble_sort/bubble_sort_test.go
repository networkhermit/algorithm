package bubble_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestBubbleSort(t *testing.T) {
	t.Parallel()
	t.Run("Decreasing", tests.DeriveDecreasing(Sort, 10000))
	t.Run("Empty", tests.DeriveEmpty(Sort))
	t.Run("Identical", tests.DeriveIdentical(Sort, 100000))
	t.Run("Increasing", tests.DeriveIncreasing(Sort, 100000))
	t.Run("Random", tests.DeriveRandom(Sort, 10000))
}
