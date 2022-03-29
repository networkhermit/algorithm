package QuickSort

import (
	"testing"

	"muse/util/SequenceBuilder"
	"muse/util/Sequences"
)

func TestQuickSort(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	SequenceBuilder.PackRandom(arr)

	checksum := Sequences.ParityChecksum(arr)

	Sort(arr)

	if Sequences.ParityChecksum(arr) != checksum {
		t.FailNow()
	}

	if !Sequences.IsSorted(arr) {
		t.FailNow()
	}
}
