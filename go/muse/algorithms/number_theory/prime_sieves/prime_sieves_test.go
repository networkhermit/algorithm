package prime_sieves

import (
	"testing"

	"muse/algorithms/number_theory/primality"
	"muse/algorithms/number_theory/prime_sieves/tests"
)

func TestPrimeSieves(t *testing.T) {
	t.Run("SieveOfEratosthenes", tests.Derive(SieveOfEratosthenes, primality.IsPrime))
}
