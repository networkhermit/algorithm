package quick_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestQuickSort(t *testing.T) {
	t.Parallel()
	t.Run("Decreasing", tests.DeriveDecreasing(Sort, 100_000))
	t.Run("Empty", tests.DeriveEmpty(Sort))
	t.Run("Identical", tests.DeriveIdentical(Sort, 100_000))
	t.Run("Increasing", tests.DeriveIncreasing(Sort, 100_000))
	t.Run("Random", tests.DeriveRandom(Sort, 100_000))
}

func TestQuickSortInefficient(t *testing.T) {
	t.Parallel()
	t.Run("Decreasing", tests.DeriveDecreasing(SortInefficient, 100_000))
	t.Run("Empty", tests.DeriveEmpty(Sort))
	t.Run("Identical", tests.DeriveIdentical(SortInefficient, 100_000))
	t.Run("Increasing", tests.DeriveIncreasing(SortInefficient, 100_000))
	t.Run("Random", tests.DeriveRandom(SortInefficient, 100_000))
}
