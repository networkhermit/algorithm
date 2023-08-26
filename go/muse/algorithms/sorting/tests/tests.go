package tests

import (
	"testing"

	"muse/util/sequence_builder"
	"muse/util/sequences"
)

var Derive = func(fn func([]int), size int, pack func([]int)) func(t *testing.T) {
	return func(t *testing.T) {
		arr := make([]int, size)
		pack(arr)

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

var DeriveDecreasing = func(fn func([]int), size int) func(t *testing.T) {
	return Derive(fn, size, sequence_builder.PackDecreasing)
}

var DeriveEmpty = func(fn func([]int)) func(t *testing.T) {
	return Derive(fn, 0, sequence_builder.PackIdentical)
}

var DeriveIdentical = func(fn func([]int), size int) func(t *testing.T) {
	return Derive(fn, size, sequence_builder.PackIdentical)
}

var DeriveIncreasing = func(fn func([]int), size int) func(t *testing.T) {
	return Derive(fn, size, sequence_builder.PackIncreasing)
}

var DeriveRandom = func(fn func([]int), size int) func(t *testing.T) {
	return Derive(fn, size, sequence_builder.PackRandom)
}
