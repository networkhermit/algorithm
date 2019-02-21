require "fun/vac/algorithms/number_theory/GreatestCommonDivisor"

module LeastCommonMultiple

    def self.reduceToBinaryGCD(m, n)
        if m < 0
            m = -m
        end
        if n < 0
            n = -n
        end

        if m == 0 || n == 0
            return 0
        else
            return m / GreatestCommonDivisor.iterativeBinaryGCD(m, n) * n
        end
    end

    def self.reduceToEuclidean(m, n)
        if m < 0
            m = -m
        end
        if n < 0
            n = -n
        end

        if m == 0 || n == 0
            return 0
        else
            return m / GreatestCommonDivisor.iterativeEuclidean(m, n) * n
        end
    end
end
