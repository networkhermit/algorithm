module BinarySearch

    def self.find(arr, key)
        lo = 0
        hi = arr.length

        mid = 0

        while lo < hi
            mid = lo + ((hi - lo) >> 1)
            case
            when key < arr[mid]
                hi = mid
            when key > arr[mid]
                lo = mid + 1
            else
                return mid
            end
        end

        return arr.length
    end
end
