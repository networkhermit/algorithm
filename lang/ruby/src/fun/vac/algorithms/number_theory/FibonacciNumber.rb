module FibonacciNumber

    def self.iterativeProcedure(n)
        sign = 1

        if n.negative?
            if (n & 1).zero?
                sign = -1
            end
            n = -n
        end

        if n < 2
            return n
        end

        prev = 0
        curr = 1

        (2 .. n).each do
            prev, curr = curr, prev + curr
        end

        sign * curr
    end

    def self.recursiveProcedure(n)
        if n.negative?
            if (n & 1).zero?
                return -recursiveProcedure(-n)
            end
            return recursiveProcedure(-n)
        end
        if n < 2
            return n
        end
        recursiveProcedure(n - 2) + recursiveProcedure(n - 1)
    end
end
