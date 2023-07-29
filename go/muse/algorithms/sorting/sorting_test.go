package sorting

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort", tests.Derive(BubbleSort))
}

func TestInsertionSort(t *testing.T) {
	t.Run("InsertionSort", tests.Derive(InsertionSort))
}

func TestMergeSort(t *testing.T) {
	t.Run("MergeSort", tests.Derive(MergeSort))
}

func TestQuickSort(t *testing.T) {
	t.Run("QuickSort", tests.Derive(QuickSort))
}

func TestSelectionSort(t *testing.T) {
	t.Run("SelectionSort", tests.Derive(SelectionSort))
}
