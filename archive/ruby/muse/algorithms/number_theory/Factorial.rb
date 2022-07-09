module Factorial
  def self.iterativeProcedure(n)
    return 0 if n.negative?

    result = 1
    1.upto(n) do |i|
      result *= i
    end

    result
  end

  def self.recursiveProcedure(n)
    return 0 if n.negative?

    return 1 if n.zero?

    recursiveProcedure(n - 1) * n
  end
end
