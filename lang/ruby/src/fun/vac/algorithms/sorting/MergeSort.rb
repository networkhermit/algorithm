module MergeSort

    def self.merge(arr, lo, mid, hi)
        if lo == mid
            return
        end

        merge(arr, lo, (lo + mid) >> 1, mid)
        merge(arr, mid, (mid + hi) >> 1, hi)

        m = lo
        n = mid

        sorted = Array.new(hi - lo)

        for i in 0 ... sorted.length
            if m != mid && (n == hi || arr[m] < arr[n])
                sorted[i] = arr[m]
                m += 1
            else
                sorted[i] = arr[n]
                n += 1
            end
        end

        cursor = 0
        for i in lo ... hi
            arr[i] = sorted[cursor]
            cursor += 1
        end
    end

    def self.sort(arr)
        merge(arr, 0, arr.length >> 1, arr.length)
    end
end
