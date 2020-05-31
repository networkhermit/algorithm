module QuickSort

    def self.partition(arr, lo, hi)
        if lo == hi
            return
        end

        pivot = arr[lo]

        left = lo
        right = hi - 1

        until left == right
            right.downto(left + 1) do |i|
                if arr[i] < pivot
                    arr[left] = arr[i]
                    arr[i] = pivot
                    break
                end
                right -= 1
            end
            left.upto(right - 1) do |i|
                if arr[i] > pivot
                    arr[right] = arr[i]
                    arr[i] = pivot
                    break
                end
                left += 1
            end
        end

        partition(arr, lo, left)
        partition(arr, left + 1, hi)
    end

    def self.sort(arr)
        partition(arr, 0, arr.length)
    end
end
