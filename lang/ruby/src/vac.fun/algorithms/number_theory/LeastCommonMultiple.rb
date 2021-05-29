require "vac.fun/algorithms/number_theory/GreatestCommonDivisor"

module LeastCommonMultiple

    def self.reduceToBinaryGCD(m, n)
        if m.negative?
            m = -m
        end
        if n.negative?
            n = -n
        end

        if m.zero? || n.zero?
            return 0
        end
        m / GreatestCommonDivisor.iterativeBinaryGCD(m, n) * n
    end

    def self.reduceToEuclidean(m, n)
        if m.negative?
            m = -m
        end
        if n.negative?
            n = -n
        end

        if m.zero? || n.zero?
            return 0
        end
        m / GreatestCommonDivisor.iterativeEuclidean(m, n) * n
    end
end
