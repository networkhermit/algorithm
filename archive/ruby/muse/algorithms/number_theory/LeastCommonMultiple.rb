require "muse/algorithms/number_theory/GreatestCommonDivisor"

module LeastCommonMultiple
  def self.reduceToBinaryGCD(m, n)
    m = -m if m.negative?
    n = -n if n.negative?

    return 0 if m.zero? || n.zero?

    m / GreatestCommonDivisor.iterativeBinaryGCD(m, n) * n
  end

  def self.reduceToEuclidean(m, n)
    m = -m if m.negative?
    n = -n if n.negative?

    return 0 if m.zero? || n.zero?

    m / GreatestCommonDivisor.iterativeEuclidean(m, n) * n
  end
end
