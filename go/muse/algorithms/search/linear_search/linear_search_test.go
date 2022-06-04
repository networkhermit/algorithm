package linear_search

import (
	"testing"

	"muse/util/sequence_builder"
)

func TestLinearSearch(t *testing.T) {
	size := 32768

	arr := make([]int, size)
	sequence_builder.PackIncreasing(arr)

	if Find(arr, -1) != size {
		t.FailNow()
	}

	if Find(arr, 2147483647) != size {
		t.FailNow()
	}

	for i, v := range arr {
		if Find(arr, v) != i {
			t.FailNow()
		}
	}
}
