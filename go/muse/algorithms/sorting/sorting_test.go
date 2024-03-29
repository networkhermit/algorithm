package sorting

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort", tests.DeriveRandom(BubbleSort, 32768))
}

func TestInsertionSort(t *testing.T) {
	t.Run("InsertionSort", tests.DeriveRandom(InsertionSort, 32768))
}

func TestMergeSort(t *testing.T) {
	t.Run("MergeSort", tests.DeriveRandom(MergeSort, 32768))
}

func TestQuickSort(t *testing.T) {
	t.Run("QuickSort", tests.DeriveRandom(QuickSort, 32768))
}

func TestSelectionSort(t *testing.T) {
	t.Run("SelectionSort", tests.DeriveRandom(SelectionSort, 32768))
}
