package tests

import (
	"testing"

	"muse/util/sequence_builder"
	"muse/util/sequences"
)

var Derive = func(fn func([]int)) func(t *testing.T) {
	return func(t *testing.T) {
		size := 32768

		arr := make([]int, size)
		sequence_builder.PackRandom(arr)

		checksum := sequences.ParityChecksum(arr)

		fn(arr)

		if sequences.ParityChecksum(arr) != checksum {
			t.Errorf("elements returned by %s(arr) have been altered", t.Name())
		}

		if !sequences.IsSorted(arr) {
			t.Errorf("elements returned by %s(arr) are not in sorted order", t.Name())
		}
	}
}
