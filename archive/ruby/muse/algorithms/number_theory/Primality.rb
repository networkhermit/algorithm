module Primality
  def self.isPrime(n)
    return false if n < 2
    return false if (n & 1).zero? && n != 2

    3.step(Integer.sqrt(n), 2) do |i|
      return false if (n % i).zero?
    end

    true
  end

  def self.isComposite(n)
    n > 1 && !isPrime(n)
  end
end
