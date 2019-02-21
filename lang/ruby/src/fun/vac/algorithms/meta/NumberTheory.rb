require "fun/vac/algorithms/number_theory/Coprimality"
require "fun/vac/algorithms/number_theory/Factorial"
require "fun/vac/algorithms/number_theory/FibonacciNumber"
require "fun/vac/algorithms/number_theory/GreatestCommonDivisor"
require "fun/vac/algorithms/number_theory/LeastCommonMultiple"
require "fun/vac/algorithms/number_theory/Parity"
require "fun/vac/algorithms/number_theory/Primality"
require "fun/vac/algorithms/number_theory/PrimeSieves"

module NumberTheory

    def self.isCoprime(m, n)
        return Coprimality.reduceToBinaryGCD(m, n)
    end

    def self.factorial(n)
        return Factorial.iterativeProcedure(n)
    end

    def self.fibonacci(n)
        return FibonacciNumber.iterativeProcedure(n)
    end

    def self.gcd(m, n)
        return GreatestCommonDivisor.iterativeBinaryGCD(m, n)
    end

    def self.lcm(m, n)
        return LeastCommonMultiple.reduceToBinaryGCD(m, n)
    end

    def self.isEven(n)
        return Parity.bitwiseIsEven(n)
    end

    def self.isOdd(n)
        return Parity.bitwiseIsOdd(n)
    end

    def self.isPrime(n)
        return Primality.isPrime(n)
    end

    def self.isComposite(n)
        return Primality.isComposite(n)
    end

    def self.sieveOfPrimes(n)
        return PrimeSieves.sieveOfEratosthenes(n)
    end
end
