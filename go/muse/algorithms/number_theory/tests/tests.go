package tests

import "testing"

type UniqueCategorySample[T, C comparable] struct {
	N        T
	Category C
}

func UniqueCategoryDerive[T, C comparable](fn func(T) bool, sample []UniqueCategorySample[T, C], c C) func(t *testing.T) {
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

type MNUniqueCategorySample[T, C comparable] struct {
	M        T
	N        T
	Category C
}

func MNUniqueCategoryDerive[T, C comparable](fn func(T, T) bool, sample []MNUniqueCategorySample[T, C], c C) func(t *testing.T) {
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
