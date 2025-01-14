package sorting

import (
	"testing"

	"muse/algorithms/sorting/tests"
)

func TestBubbleSort(t *testing.T) {
	t.Run("BubbleSort", tests.DeriveRandom(BubbleSort, 10000))
}

func TestInsertionSort(t *testing.T) {
	t.Run("InsertionSort", tests.DeriveRandom(InsertionSort, 10000))
}

func TestMergeSort(t *testing.T) {
	t.Run("MergeSort", tests.DeriveRandom(MergeSort, 100000))
}

func TestQuickSort(t *testing.T) {
	t.Run("QuickSort", tests.DeriveRandom(QuickSort, 100000))
}

func TestSelectionSort(t *testing.T) {
	t.Run("SelectionSort", tests.DeriveRandom(SelectionSort, 10000))
}

func TestShellSort(t *testing.T) {
	t.Run("ShellSort", tests.DeriveRandom(ShellSort, 100000))
}
