#ifndef MUSE_ALGORITHMS_META_NUMBER_THEORY_HPP
#define MUSE_ALGORITHMS_META_NUMBER_THEORY_HPP

#include <muse/algorithms/number_theory/Coprimality.hpp>
#include <muse/algorithms/number_theory/Factorial.hpp>
#include <muse/algorithms/number_theory/FibonacciNumber.hpp>
#include <muse/algorithms/number_theory/GreatestCommonDivisor.hpp>
#include <muse/algorithms/number_theory/LeastCommonMultiple.hpp>
#include <muse/algorithms/number_theory/Parity.hpp>
#include <muse/algorithms/number_theory/Primality.hpp>
#include <muse/algorithms/number_theory/PrimeSieves.hpp>

namespace NumberTheory {

    bool isCoprime(long m, long n) {
        return Coprimality::reduceToBinaryGCD(m, n);
    }

    long factorial(long n) {
        return Factorial::iterativeProcedure(n);
    }

    long fibonacci(long n) {
        return FibonacciNumber::iterativeProcedure(n);
    }

    long gcd(long m, long n) {
        return GreatestCommonDivisor::iterativeBinaryGCD(m, n);
    }

    long lcm(long m, long n) {
        return LeastCommonMultiple::reduceToBinaryGCD(m, n);
    }

    bool isEven(long n) {
        return Parity::bitwiseIsEven(n);
    }

    bool isOdd(long n) {
        return Parity::bitwiseIsOdd(n);
    }

    bool isPrime(long n) {
        return Primality::isPrime(n);
    }

    bool isComposite(long n) {
        return Primality::isComposite(n);
    }

    std::vector<std::size_t> sieveOfPrimes(std::size_t n) {
        return PrimeSieves::sieveOfEratosthenes(n);
    }
}

#endif
