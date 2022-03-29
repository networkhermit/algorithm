package Sorting

import (
	"testing"

	"muse/util/SequenceBuilder"
	"muse/util/Sequences"
)

func TestBubbleSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	SequenceBuilder.PackRandom(arr)

	checksum := Sequences.ParityChecksum(arr)

	BubbleSort(arr)

	if Sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !Sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestInsertionSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	SequenceBuilder.PackRandom(arr)

	checksum := Sequences.ParityChecksum(arr)

	InsertionSort(arr)

	if Sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !Sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestMergeSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	SequenceBuilder.PackRandom(arr)

	checksum := Sequences.ParityChecksum(arr)

	MergeSort(arr)

	if Sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !Sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestQuickSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	SequenceBuilder.PackRandom(arr)

	checksum := Sequences.ParityChecksum(arr)

	QuickSort(arr)

	if Sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !Sequences.IsSorted(arr) {
		t.FailNow()
	}
}

func TestSelectionSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	SequenceBuilder.PackRandom(arr)

	checksum := Sequences.ParityChecksum(arr)

	SelectionSort(arr)

	if Sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !Sequences.IsSorted(arr) {
		t.FailNow()
	}
}
