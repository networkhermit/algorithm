package tests

import (
	"testing"

	"muse/util/sequence_builder"
	"muse/util/sequences"
)

var Derive = func(sort func([]int), size int, pack func([]int)) func(t *testing.T) {
	return func(t *testing.T) {
		arr := make([]int, size)
		pack(arr)

		checksum := sequences.ParityChecksum(arr)

		sort(arr)

		if sequences.ParityChecksum(arr) != checksum {
			t.Errorf("elements returned by %s(arr) have been altered", t.Name())
		}

		if !sequences.IsSorted(arr) {
			t.Errorf("elements returned by %s(arr) are not in sorted order", t.Name())
		}
	}
}

var DeriveDecreasing = func(sort func([]int), size int) func(t *testing.T) {
	return Derive(sort, size, sequence_builder.PackDecreasing)
}

var DeriveEmpty = func(sort func([]int)) func(t *testing.T) {
	return Derive(sort, 0, sequence_builder.PackIdentical)
}

var DeriveIdentical = func(sort func([]int), size int) func(t *testing.T) {
	return Derive(sort, size, sequence_builder.PackIdentical)
}

var DeriveIncreasing = func(sort func([]int), size int) func(t *testing.T) {
	return Derive(sort, size, sequence_builder.PackIncreasing)
}

var DeriveRandom = func(sort func([]int), size int) func(t *testing.T) {
	return Derive(sort, size, sequence_builder.PackRandom)
}
