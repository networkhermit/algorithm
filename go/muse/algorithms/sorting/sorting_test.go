package sorting

import (
	"testing"

	"muse/util/sequence_builder"
	"muse/util/sequences"
)

func TestBubbleSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	sequence_builder.PackRandom(arr)

	checksum := sequences.ParityChecksum(arr)

	BubbleSort(arr)

	if sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestInsertionSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	sequence_builder.PackRandom(arr)

	checksum := sequences.ParityChecksum(arr)

	InsertionSort(arr)

	if sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestMergeSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	sequence_builder.PackRandom(arr)

	checksum := sequences.ParityChecksum(arr)

	MergeSort(arr)

	if sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestQuickSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	sequence_builder.PackRandom(arr)

	checksum := sequences.ParityChecksum(arr)

	QuickSort(arr)

	if sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestSelectionSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	sequence_builder.PackRandom(arr)

	checksum := sequences.ParityChecksum(arr)

	SelectionSort(arr)

	if sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !sequences.IsSorted(arr) {
		t.FailNow()
	}
}
