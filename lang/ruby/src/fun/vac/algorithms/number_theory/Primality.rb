module Primality

    def self.isPrime(n)
        if n < 2
            return false
        end
        if (n & 1).zero? && n != 2
            return false
        end

        i, bound = 3, Math.sqrt(n).to_i() + 1
        while i < bound
            if (n % i).zero?
                return false
            end
            i += 2
        end

        true
    end

    def self.isComposite(n)
        n > 1 && !isPrime(n)
    end
end
