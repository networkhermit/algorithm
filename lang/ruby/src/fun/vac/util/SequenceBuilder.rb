require "fun/vac/util/RandomFactory"
require "fun/vac/util/Sequences"

module SequenceBuilder

    def self.packIncreasing(arr)
        RandomFactory.seed()
        arr[0] = RandomFactory.integerN(1, 3)
        for i in 1 ... arr.length
            arr[i] = arr[i - 1] + RandomFactory.integerN(1, 3)
        end
    end

    def self.packRandom(arr)
        RandomFactory.seed()
        arr.each_index do |i|
            arr[i] = RandomFactory.generateInteger()
        end
    end

    def self.packDecreasing(arr)
        packIncreasing(arr)
        Sequences.reverse(arr)
    end
end
