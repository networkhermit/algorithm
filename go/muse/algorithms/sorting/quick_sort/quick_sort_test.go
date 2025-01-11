package quick_sort

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestQuickSort(t *testing.T) {
	t.Parallel()
	t.Run("Decreasing", tests.DeriveDecreasing(Sort, 100000))
	t.Run("Empty", tests.DeriveEmpty(Sort))
	t.Run("Identical", tests.DeriveIdentical(Sort, 100000))
	t.Run("Increasing", tests.DeriveIncreasing(Sort, 100000))
	t.Run("Random", tests.DeriveRandom(Sort, 100000))
}

func TestQuickSortThreeWayPartition(t *testing.T) {
	t.Parallel()
	t.Run("Decreasing", tests.DeriveDecreasing(SortThreeWayPartition, 10000))
	t.Run("Empty", tests.DeriveEmpty(SortThreeWayPartition))
	t.Run("Identical", tests.DeriveIdentical(SortThreeWayPartition, 100000))
	t.Run("Increasing", tests.DeriveIncreasing(SortThreeWayPartition, 10000))
	t.Run("Random", tests.DeriveRandom(SortThreeWayPartition, 100000))
	t.Run("DecreasingWithShuffle", tests.DeriveDecreasing(SortThreeWayPartitionWithShuffle, 100000))
	t.Run("IncreasingWithShuffle", tests.DeriveIncreasing(SortThreeWayPartitionWithShuffle, 100000))
}

func TestQuickSortInefficient(t *testing.T) {
	t.Parallel()
	t.Run("Decreasing", tests.DeriveDecreasing(SortInefficient, 10000))
	t.Run("Empty", tests.DeriveEmpty(SortInefficient))
	t.Run("Identical", tests.DeriveIdentical(SortInefficient, 10000))
	t.Run("Increasing", tests.DeriveIncreasing(SortInefficient, 10000))
	t.Run("Random", tests.DeriveRandom(SortInefficient, 100000))
}
