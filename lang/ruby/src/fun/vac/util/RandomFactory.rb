module RandomFactory

    def self.seed()
        # preserve for consistent interface
    end

    def self.integerN(min, max)
        return min + rand(max - min + 1)
    end

    def self.generateInteger()
        return integerN(0, 2_147_483_647)
    end

    def self.generateEven()
        return generateInteger() >> 1 << 1
    end

    def self.generateOdd()
        return generateInteger() | 1
    end
end
