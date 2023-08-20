package tests

import "testing"

var sample = []struct {
	n        int64
	expected int64
}{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{5, 120},
	{6, 720},
	{7, 5040},
	{8, 40320},
	{9, 362880},
	{10, 3_628_800},
	{11, 39_916_800},
	{12, 479_001_600},
}

var Derive = func(fn func(int64) int64) func(t *testing.T) {
	return func(t *testing.T) {
		for _, tt := range sample {
			actual := fn(tt.n)
			if actual != tt.expected {
				t.Errorf("%s(%d) returned %d, expected %d", t.Name(), tt.n, actual, tt.expected)
			}
		}
	}
}
