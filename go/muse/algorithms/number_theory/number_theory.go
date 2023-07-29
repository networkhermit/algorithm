package number_theory

import (
	"muse/algorithms/number_theory/coprimality"
	"muse/algorithms/number_theory/factorial"
	"muse/algorithms/number_theory/fibonacci_number"
	"muse/algorithms/number_theory/greatest_common_divisor"
	"muse/algorithms/number_theory/least_common_multiple"
	"muse/algorithms/number_theory/parity"
	"muse/algorithms/number_theory/primality"
	"muse/algorithms/number_theory/prime_sieves"
)

func IsCoprime(m, n int64) bool {
	return coprimality.ReduceToBinaryGCD(m, n)
}

func Factorial(n int64) int64 {
	return factorial.IterativeProcedure(n)
}

func Fibonacci(n int64) int64 {
	return fibonacci_number.IterativeProcedure(n)
}

func GCD(m, n int64) int64 {
	return greatest_common_divisor.IterativeBinaryGCD(m, n)
}

func LCM(m, n int64) int64 {
	return least_common_multiple.ReduceToBinaryGCD(m, n)
}

func IsEven(n int64) bool {
	return parity.BitwiseIsEven(n)
}

func IsOdd(n int64) bool {
	return parity.BitwiseIsOdd(n)
}

func IsPrime(n int64) bool {
	return primality.IsPrime(n)
}

func IsComposite(n int64) bool {
	return primality.IsComposite(n)
}

func SieveOfPrimes(n int) []int {
	return prime_sieves.SieveOfEratosthenes(n)
}
