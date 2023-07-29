package coprimality

import (
	"testing"

	"muse/algorithms/number_theory/coprimality/tests"
)

func TestCoprimality(t *testing.T) {
	t.Parallel()

	t.Run("ReduceToBinaryGCD", tests.Derive(ReduceToBinaryGCD))

	t.Run("ReduceToEuclidean", tests.Derive(ReduceToEuclidean))
}
