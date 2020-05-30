require "fun/vac/algorithms/number_theory/GreatestCommonDivisor"

module Coprimality

    def self.reduceToBinaryGCD(m, n)
        GreatestCommonDivisor.iterativeBinaryGCD(m, n) == 1
    end

    def self.reduceToEuclidean(m, n)
        GreatestCommonDivisor.iterativeEuclidean(m, n) == 1
    end
end
