module RandomFactory

    def self.seed()
        # preserve for consistent interface
    end

    def self.integerN(min, max)
        min + rand(max - min + 1)
    end

    def self.generateInteger()
        integerN(0, 2_147_483_647)
    end

    def self.generateEven()
        generateInteger() >> 1 << 1
    end

    def self.generateOdd()
        generateInteger() | 1
    end
end
