module FibonacciNumber

    def self.iterativeProcedure(n)
        sign = 1

        if n < 0
            if (n & 1) == 0
                sign = -1
            end
            n = -n
        end

        if n < 2
            return n
        end

        prev = 0
        curr = 1

        (2 .. n).each do |i|
            prev, curr = curr, prev + curr
        end

        return sign * curr
    end

    def self.recursiveProcedure(n)
        if n < 0
            if (n & 1) == 0
                return -recursiveProcedure(-n)
            end
            return recursiveProcedure(-n)
        end
        if n < 2
            return n
        end
        return recursiveProcedure(n - 2) + recursiveProcedure(n - 1)
    end
end
