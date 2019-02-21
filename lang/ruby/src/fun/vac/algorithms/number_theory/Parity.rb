module Parity

    def self.moduloIsEven(n)
        return n % 2 == 0
    end

    def self.moduloIsOdd(n)
        return n % 2 != 0
    end

    def self.bitwiseIsEven(n)
        return (n & 1) == 0
    end

    def self.bitwiseIsOdd(n)
        return (n & 1) != 0
    end
end
