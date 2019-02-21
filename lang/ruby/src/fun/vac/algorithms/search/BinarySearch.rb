module BinarySearch

    def self.find(arr, key)
        lo = 0
        hi = arr.length

        mid = 0

        while lo < hi
            mid = lo + ((hi - lo) >> 1)
            if key < arr[mid]
                hi = mid
            elsif key > arr[mid]
                lo = mid + 1
            else
                return mid
            end
        end

        return arr.length
    end
end
