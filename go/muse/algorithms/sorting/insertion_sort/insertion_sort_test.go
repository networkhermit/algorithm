package insertion_sort

import (
	"testing"

	"muse/util/sequence_builder"
	"muse/util/sequences"
)

func TestInsertionSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	sequence_builder.PackRandom(arr)

	checksum := sequences.ParityChecksum(arr)

	Sort(arr)

	if sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !sequences.IsSorted(arr) {
		t.FailNow()
	}
}
