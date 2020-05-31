module Primality

    def self.isPrime(n)
        if n < 2
            return false
        end
        if (n & 1).zero? && n != 2
            return false
        end

        3.step(Math.sqrt(n).to_i(), 2) do |i|
            if (n % i).zero?
                return false
            end
        end

        true
    end

    def self.isComposite(n)
        n > 1 && !isPrime(n)
    end
end
