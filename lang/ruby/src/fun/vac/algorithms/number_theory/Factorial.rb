module Factorial

    def self.iterativeProcedure(n)
        if n < 0
            return 0
        end

        result = 1
        for i in 1 .. n
            result *= i
        end

        return result
    end

    def self.recursiveProcedure(n)
        if n < 0
            return 0
        end

        if n == 0
            return 1
        end
        return recursiveProcedure(n - 1) * n
    end
end
