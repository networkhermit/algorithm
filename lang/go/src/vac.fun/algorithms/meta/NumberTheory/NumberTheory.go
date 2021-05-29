package NumberTheory

import (
    "vac.fun/algorithms/number_theory/Coprimality"
    factorial "vac.fun/algorithms/number_theory/Factorial"
    "vac.fun/algorithms/number_theory/FibonacciNumber"
    "vac.fun/algorithms/number_theory/GreatestCommonDivisor"
    "vac.fun/algorithms/number_theory/LeastCommonMultiple"
    "vac.fun/algorithms/number_theory/Parity"
    "vac.fun/algorithms/number_theory/Primality"
    "vac.fun/algorithms/number_theory/PrimeSieves"
)

func IsCoprime(m int64, n int64) bool {
    return Coprimality.ReduceToBinaryGCD(m, n)
}

func Factorial(n int64) int64 {
    return factorial.IterativeProcedure(n)
}

func Fibonacci(n int64) int64 {
    return FibonacciNumber.IterativeProcedure(n)
}

func GCD(m int64, n int64) int64 {
    return GreatestCommonDivisor.IterativeBinaryGCD(m, n)
}

func LCM(m int64, n int64) int64 {
    return LeastCommonMultiple.ReduceToBinaryGCD(m, n)
}

func IsEven(n int64) bool {
    return Parity.BitwiseIsEven(n)
}

func IsOdd(n int64) bool {
    return Parity.BitwiseIsOdd(n)
}

func IsPrime(n int64) bool {
    return Primality.IsPrime(n)
}

func IsComposite(n int64) bool {
    return Primality.IsComposite(n)
}

func SieveOfPrimes(n int) []int {
    return PrimeSieves.SieveOfEratosthenes(n)
}
