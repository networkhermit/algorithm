package tests

import "testing"

type UniqueCategorySample[N, C comparable] struct {
	N        N
	Category C
}

func UniqueCategoryDerive[N, C comparable](fn func(N) bool, sample []UniqueCategorySample[N, C], c C) func(t *testing.T) {
	return func(t *testing.T) {
		for _, tc := range sample {
			actual := fn(tc.N)
			expected := tc.Category == c
			if actual != expected {
				t.Errorf("%s(%v) returned %t, expect %t", t.Name(), tc.N, actual, expected)
			}
		}
	}
}

type MNUniqueCategorySample[N, C comparable] struct {
	M        N
	N        N
	Category C
}

func MNUniqueCategoryDerive[N, C comparable](fn func(N, N) bool, sample []MNUniqueCategorySample[N, C], c C) func(t *testing.T) {
	return func(t *testing.T) {
		for _, tc := range sample {
			actual := fn(tc.M, tc.N)
			expected := tc.Category == c
			if actual != expected {
				t.Errorf("%s(%v, %v) returned %t, expect %t", t.Name(), tc.M, tc.N, actual, expected)
			}
		}
	}
}
