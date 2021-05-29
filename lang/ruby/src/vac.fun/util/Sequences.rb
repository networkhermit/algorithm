require "vac.fun/algorithms/sorting/QuickSort"
require "vac.fun/util/RandomFactory"

module Sequences

    def self.inspect(arr)
        puts("[")
        arr.each_index do |i|
            print(format("\t#%04X  -->  %d\n", i, arr[i]))
        end
        puts("]")
    end

    def self.isSorted(arr)
        1.upto(arr.length - 1) do |i|
            if arr[i] < arr[i - 1]
                return false
            end
        end

        true
    end

    def self.parityChecksum(arr)
        checksum = 0

        arr.each do |v|
            checksum ^= v
        end

        checksum
    end

    def self.reverse(arr)
        k = 0

        length = arr.length
        0.upto((arr.length >> 1) - 1) do |i|
            k = length - i - 1
            arr[i], arr[k] = arr[k], arr[i]
        end
    end

    def self.shuffle(arr)
        k = 0
        length = arr.length

        RandomFactory.seed()
        0.upto(arr.length - 1) do |i|
            k = RandomFactory.integerN(i, length)
            arr[i], arr[k] = arr[k], arr[i]
        end
    end

    def self.sort(arr)
        QuickSort.sort(arr)
    end
end
