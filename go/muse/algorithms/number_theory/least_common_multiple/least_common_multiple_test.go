package least_common_multiple

import (
	"testing"

	"muse/algorithms/number_theory/least_common_multiple/tests"
)

func TestLeastCommonMultiple(t *testing.T) {
	t.Parallel()

	t.Run("ReduceToBinaryGCD", tests.Derive(ReduceToBinaryGCD))

	t.Run("ReduceToEuclidean", tests.Derive(ReduceToEuclidean))
}
