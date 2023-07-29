package factorial

import (
	"testing"

	"muse/algorithms/number_theory/factorial/tests"
)

func TestFactorial(t *testing.T) {
	t.Parallel()

	t.Run("IterativeProcedure", tests.Derive(IterativeProcedure))

	t.Run("RecursiveProcedure", tests.Derive(RecursiveProcedure))
}
