package fibonacci_number

import (
	"testing"

	"muse/algorithms/number_theory/fibonacci_number/tests"
)

func TestFibonacciNumber(t *testing.T) {
	t.Parallel()

	t.Run("IterativeProcedure", tests.Derive(IterativeProcedure))

	t.Run("RecursiveProcedure", tests.Derive(RecursiveProcedure))
}
