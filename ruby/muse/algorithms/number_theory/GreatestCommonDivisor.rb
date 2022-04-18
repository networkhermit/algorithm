module GreatestCommonDivisor
  def self.iterativeBinaryGCD(m, n)
    m = -m if m.negative?
    n = -n if n.negative?

    shift = 0

    loop do
      return m << shift if m == n
      return n << shift if m.zero?
      return m << shift if n.zero?

      if (m & 1).zero?
        m >>= 1
        if (n & 1).zero?
          n >>= 1
          shift += 1
        end
      elsif (n & 1).zero?
        n >>= 1
      elsif m > n
        m = (m - n) >> 1
      else
        n = (n - m) >> 1
      end
    end
  end

  def self.recursiveBinaryGCD(m, n)
    m = -m if m.negative?
    n = -n if n.negative?

    return m if m == n
    return n if m.zero?
    return m if n.zero?

    if (m & 1).zero?
      return recursiveBinaryGCD(m >> 1, n >> 1) << 1 if (n & 1).zero?

      return recursiveBinaryGCD(m >> 1, n)
    end
    return recursiveBinaryGCD(m, n >> 1) if (n & 1).zero?
    return recursiveBinaryGCD((m - n) >> 1, n) if m > n

    recursiveBinaryGCD(m, (n - m) >> 1)
  end

  def self.iterativeEuclidean(m, n)
    m = -m if m.negative?
    n = -n if n.negative?

    m, n = n, m % n until n.zero?

    m
  end

  def self.recursiveEuclidean(m, n)
    m = -m if m.negative?
    n = -n if n.negative?

    return m if n.zero?

    recursiveEuclidean(n, m % n)
  end
end
