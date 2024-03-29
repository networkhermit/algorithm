module Parity
  def self.moduloIsEven(n)
    (n % 2).zero?
  end

  def self.moduloIsOdd(n)
    n.odd?
  end

  def self.bitwiseIsEven(n)
    (n & 1).zero?
  end

  def self.bitwiseIsOdd(n)
    (n & 1) != 0
  end
end
