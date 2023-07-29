package number_theory

import (
	"testing"

	coprimality_tests "muse/algorithms/number_theory/coprimality/tests"
	factorial_tests "muse/algorithms/number_theory/factorial/tests"
	fibonacci_number_tests "muse/algorithms/number_theory/fibonacci_number/tests"
	greatest_common_divisor_tests "muse/algorithms/number_theory/greatest_common_divisor/tests"
	least_common_multiple_tests "muse/algorithms/number_theory/least_common_multiple/tests"
	parity_tests "muse/algorithms/number_theory/parity/tests"
	"muse/algorithms/number_theory/primality"
	primality_tests "muse/algorithms/number_theory/primality/tests"
	prime_sieves_tests "muse/algorithms/number_theory/prime_sieves/tests"
)

func TestIsCoprime(t *testing.T) {
	t.Run("IsCoprime", coprimality_tests.Derive(IsCoprime))
}

func TestFactorial(t *testing.T) {
	t.Run("Factorial", factorial_tests.Derive(Factorial))
}

func TestFibonacci(t *testing.T) {
	t.Run("Fibonacci", fibonacci_number_tests.Derive(Fibonacci))
}

func TestGCD(t *testing.T) {
	t.Run("GCD", greatest_common_divisor_tests.Derive(GCD))
}

func TestLCM(t *testing.T) {
	t.Run("LCM", least_common_multiple_tests.Derive(LCM))
}

func TestIsEven(t *testing.T) {
	t.Run("IsEven", parity_tests.Derive(IsEven, parity_tests.Even))
}

func TestIsOdd(t *testing.T) {
	t.Run("IsOdd", parity_tests.Derive(IsOdd, parity_tests.Odd))
}

func TestIsPrime(t *testing.T) {
	t.Run("IsPrime", primality_tests.Derive(IsPrime, primality_tests.Prime))
}

func TestIsComposite(t *testing.T) {
	t.Run("IsComposite", primality_tests.Derive(IsComposite, primality_tests.Composite))
}

func TestSieveOfPrimes(t *testing.T) {
	t.Run("SieveOfPrimes", prime_sieves_tests.Derive(SieveOfPrimes, primality.IsPrime))
}
