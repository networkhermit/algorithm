package greatest_common_divisor

import (
	"testing"

	"muse/algorithms/number_theory/greatest_common_divisor/tests"
)

func TestGreatestCommonDivisor(t *testing.T) {
	t.Parallel()

	t.Run("IterativeBinaryGCD", tests.Derive(IterativeBinaryGCD))

	t.Run("RecursiveBinaryGCD", tests.Derive(RecursiveBinaryGCD))

	t.Run("IterativeEuclidean", tests.Derive(IterativeEuclidean))

	t.Run("RecursiveEuclidean", tests.Derive(RecursiveEuclidean))
}
