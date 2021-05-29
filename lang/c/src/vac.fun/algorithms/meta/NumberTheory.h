#ifndef VAC_FUN_ALGORITHMS_META_NUMBER_THEORY_H
#define VAC_FUN_ALGORITHMS_META_NUMBER_THEORY_H 1

#include <vac.fun/algorithms/number_theory/Coprimality.h>
#include <vac.fun/algorithms/number_theory/Factorial.h>
#include <vac.fun/algorithms/number_theory/FibonacciNumber.h>
#include <vac.fun/algorithms/number_theory/GreatestCommonDivisor.h>
#include <vac.fun/algorithms/number_theory/LeastCommonMultiple.h>
#include <vac.fun/algorithms/number_theory/Parity.h>
#include <vac.fun/algorithms/number_theory/Primality.h>
#include <vac.fun/algorithms/number_theory/PrimeSieves.h>

bool NumberTheory_isCoprime(long m, long n) {
    return Coprimality_reduceToBinaryGCD(m, n);
}

long NumberTheory_factorial(long n) {
    return Factorial_iterativeProcedure(n);
}

long NumberTheory_fibonacci(long n) {
    return FibonacciNumber_iterativeProcedure(n);
}

long NumberTheory_gcd(long m, long n) {
    return GreatestCommonDivisor_iterativeBinaryGCD(m, n);
}

long NumberTheory_lcm(long m, long n) {
    return LeastCommonMultiple_reduceToBinaryGCD(m, n);
}

bool NumberTheory_isEven(long n) {
    return Parity_bitwiseIsEven(n);
}

bool NumberTheory_isOdd(long n) {
    return Parity_bitwiseIsOdd(n);
}

bool NumberTheory_isPrime(long n) {
    return Primality_isPrime(n);
}

bool NumberTheory_isComposite(long n) {
    return Primality_isComposite(n);
}

size_t *NumberTheory_sieveOfPrimes(size_t n, size_t *length) {
    return PrimeSieves_sieveOfEratosthenes(n, length);
}

#endif