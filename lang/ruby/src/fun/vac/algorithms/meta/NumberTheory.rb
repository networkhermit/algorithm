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
        Coprimality.reduceToBinaryGCD(m, n)
    end

    def self.factorial(n)
        Factorial.iterativeProcedure(n)
    end

    def self.fibonacci(n)
        FibonacciNumber.iterativeProcedure(n)
    end

    def self.gcd(m, n)
        GreatestCommonDivisor.iterativeBinaryGCD(m, n)
    end

    def self.lcm(m, n)
        LeastCommonMultiple.reduceToBinaryGCD(m, n)
    end

    def self.isEven(n)
        Parity.bitwiseIsEven(n)
    end

    def self.isOdd(n)
        Parity.bitwiseIsOdd(n)
    end

    def self.isPrime(n)
        Primality.isPrime(n)
    end

    def self.isComposite(n)
        Primality.isComposite(n)
    end

    def self.sieveOfPrimes(n)
        PrimeSieves.sieveOfEratosthenes(n)
    end
end
