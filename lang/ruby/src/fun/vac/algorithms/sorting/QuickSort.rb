module QuickSort

    def self.partition(arr, lo, hi)
        if lo == hi
            return
        end

        pivot = arr[lo]

        left, right = lo, hi - 1

        until left == right
            i = right
            while i > left
                if arr[i] < pivot
                    arr[left] = arr[i]
                    arr[i] = pivot
                    break
                end
                right -= 1
                i -= 1
            end
            i = left
            while i < right
                if arr[i] > pivot
                    arr[right] = arr[i]
                    arr[i] = pivot
                    break
                end
                left += 1
                i += 1
            end
        end

        partition(arr, lo, left)
        partition(arr, left + 1, hi)
    end

    def self.sort(arr)
        partition(arr, 0, arr.length)
    end
end
