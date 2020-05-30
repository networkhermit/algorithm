module Factorial

    def self.iterativeProcedure(n)
        if n.negative?
            return 0
        end

        result = 1
        (1 .. n).each do |i|
            result *= i
        end

        result
    end

    def self.recursiveProcedure(n)
        if n.negative?
            return 0
        end

        if n.zero?
            return 1
        end
        recursiveProcedure(n - 1) * n
    end
end
