require "fun/vac/algorithms/sorting/InsertionSort"
require "fun/vac/util/RandomFactory"

module Sequences

    def self.inspect(arr)
        puts("[")
        arr.each_index do |i|
            print("\t#%04X  -->  %d\n" % [i, arr[i]])
        end
        puts("]")
    end

    def self.isSorted(arr)
        for i in 1 ... arr.length
            if arr[i] < arr[i - 1]
                return false
            end
        end

        return true
    end

    def self.parityChecksum(arr)
        checksum = 0

        for v in arr
            checksum ^= v
        end

        return checksum
    end

    def self.reverse(arr)
        k = 0

        length = arr.length
        for i in 0 ... arr.length >> 1
            k = length - i - 1
            arr[i], arr[k] = arr[k], arr[i]
        end
    end

    def self.shuffle(arr)
        k = 0
        length = arr.length

        RandomFactory.seed()
        for i in 0 ... arr.length
            k = RandomFactory.integerN(i, length)
            arr[i], arr[k] = arr[k], arr[i]
        end
    end

    def self.sort(arr)
        InsertionSort.sort(arr)
    end
end
