module FibonacciNumber
  def self.iterativeProcedure(n)
    sign = 1

    if n.negative?
      sign = -1 if (n & 1).zero?
      n = -n
    end

    return n if n < 2

    prev = 0
    curr = 1

    (n - 1).times do
      prev, curr = curr, prev + curr
    end

    sign * curr
  end

  def self.recursiveProcedure(n)
    if n.negative?
      return -recursiveProcedure(-n) if (n & 1).zero?

      return recursiveProcedure(-n)
    end
    return n if n < 2

    recursiveProcedure(n - 2) + recursiveProcedure(n - 1)
  end
end
