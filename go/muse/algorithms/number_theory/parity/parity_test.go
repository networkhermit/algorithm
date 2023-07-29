package parity

import (
	"testing"

	"muse/algorithms/number_theory/parity/tests"
)

func TestParity(t *testing.T) {
	t.Parallel()

	t.Run("ModuloIsEven", tests.Derive(ModuloIsEven, tests.Even))

	t.Run("ModuloIsOdd", tests.Derive(ModuloIsOdd, tests.Odd))

	t.Run("BitwiseIsEven", tests.Derive(BitwiseIsEven, tests.Even))

	t.Run("BitwiseIsOdd", tests.Derive(BitwiseIsOdd, tests.Odd))
}
