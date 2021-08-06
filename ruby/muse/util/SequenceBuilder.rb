require "muse/util/RandomFactory"
require "muse/util/Sequences"

module SequenceBuilder

    def self.packIncreasing(arr)
        RandomFactory.seed()
        arr[0] = RandomFactory.genIntN(1, 3)
        1.upto(arr.length - 1) do |i|
            arr[i] = arr[i - 1] + RandomFactory.genIntN(1, 3)
        end
    end

    def self.packRandom(arr)
        RandomFactory.seed()
        arr.each_index do |i|
            arr[i] = RandomFactory.genInt()
        end
    end

    def self.packDecreasing(arr)
        packIncreasing(arr)
        Sequences.reverse(arr)
    end
end
