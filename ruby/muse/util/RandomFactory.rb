module RandomFactory

    def self.seed()
        # preserve for consistent interface
    end

    def self.genIntN(min, max)
        min + rand(max - min + 1)
    end

    def self.genInt()
        genIntN(0, 2_147_483_647)
    end

    def self.genEven()
        genInt() >> 1 << 1
    end

    def self.genOdd()
        genInt() | 1
    end
end
